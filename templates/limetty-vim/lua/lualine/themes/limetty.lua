local colors = {
	lime = '{{color "lime"}}',
	cyan = '{{color "picton"}}',
	black = '{{color "black"}}',
	white = '{{color "limelight"}}',
	red = '{{color "carnation"}}',
	eucalyptus = '{color "eucalyptus"}',
	grey = '{{color "purp"}}',
	pearl = '{{color "pearl"}}',
}

local limetty = {
	normal = {
		a = { fg = '{{color "deepdark"}}', bg = colors.pearl },
		b = { fg = colors.white, bg = colors.grey },
		c = { fg = colors.black, bg = colors.black },
	},

	insert = { a = { fg = colors.black, bg = colors.eucalyptus } },
	visual = { a = { fg = colors.black, bg = colors.cyan } },
	replace = { a = { fg = colors.black, bg = colors.red } },

	command = {
		a = { fg = colors.black, bg = colors.lime },
	},

	inactive = {
		a = { fg = colors.white, bg = colors.black },
		b = { fg = colors.white, bg = colors.black },
		c = { fg = colors.black, bg = colors.black },
	},
}

return limetty
