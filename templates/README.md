<style>
html, body {
background: {{color "black"}};
color: {{color "limelight"}};
}

h1,h2,h3,h4,h5,h6 {
color: {{color "limelight"}};
}

table,tr,td {
border: 2px solid {{color "limelight"}};
padding: 20px;
}

table {
background:{{color "black"}};
color:{{color "limelight"}};
border-collapse: collapse;
}

</style>

Just another color scheme. There's no science no logic. Just random colors I liked. If you don't like them, use something else. Everythings prone to change to my personal preference so todays blue might be green tomorrow.

## Colors

<table>
{{range .Colors}}
<tr><td>{{.Name}}</td><td style="background:{{.Hex}};color:{{ if .Dark }}{{color "limelight"}}{{else}}{{color "black"}}{{end}}">{{.Hex}}</td></tr>
{{end}}
</table>
