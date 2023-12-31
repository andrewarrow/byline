package app

import (
	"byline/common"
	"html/template"
	"strings"

	"github.com/andrewarrow/feedback/markup"
	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {
	send := map[string]any{}
	lines := strings.Split(common.Sample, "\n")
	h := markup.ToHTMLFromLines(nil, lines)
	send["preview"] = template.HTML(h)
	c.SendContentInLayout("vim.html", send, 200)
}
