package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Color defines the options for a single color
type Color struct {
	Name string `yaml:"name"`
	Hex  string `yaml:"hex"`
	Dark bool   `yaml:"dark"`
}

// Config defines the configuration for the theme currently including the colors
type Config struct {
	Colors []Color `yaml:"colors"`
}

func main() {
	data, _ := os.ReadFile("./colors.yaml")

	conf := &Config{}

	err := yaml.Unmarshal(data, conf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf)

	err = filepath.Walk("./templates", func(path string, d fs.FileInfo, err error) error {
		p := strings.TrimPrefix(path, "templates/")
		p = "./" + p

		fmt.Println(path, p)

		if d.IsDir() {
			os.MkdirAll(p, 0o755)
			return nil
		}

		f, err := os.Create(p)
		if err != nil {
			return err
		}
		defer f.Close()

		tp, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		t, err := template.New("temp").
			Funcs(template.FuncMap{
				"color": func(color string) string {
					for _, c := range conf.Colors {
						if c.Name == color {
							return c.Hex
						}
					}

					return ""
				},
				"hex": func(color string) string {
					for _, c := range conf.Colors {
						if c.Name == color {
							return strings.TrimPrefix(c.Hex, "#")
						}
					}

					return ""
				},
			}).
			Parse(string(tp))
		if err != nil {
			return err
		}

		err = t.ExecuteTemplate(f, "temp", conf)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
