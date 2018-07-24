package globals

import (
	"github.com/sknv/next/app/core/cfg"
	xtemplate "github.com/sknv/next/app/lib/html/template"
)

const (
	tplRoot = "templates/"
	tplExt  = ".tpl"
)

var (
	html *xtemplate.HTML
)

func InitHTML(config *cfg.Config) {
	html = xtemplate.NewHTML(tplRoot, tplExt, config.IsRelease())
}

func GetHTML() *xtemplate.HTML {
	return html
}
