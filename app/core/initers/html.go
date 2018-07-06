package initers

import (
	xtemplate "github.com/sknv/next/app/lib/html/template"
)

const (
	tplRoot = "templates/"
	tplExt  = ".tpl"
)

var (
	html *xtemplate.HTML
)

func init() {
	html = xtemplate.NewHTML(tplRoot, tplExt, config.IsRelease())
}

func GetHTML() *xtemplate.HTML {
	return html
}
