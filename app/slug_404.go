package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Slug404(r *router.Router, c *router.Context) {

	if c.NotFound {
		router.Redirect(c, "/")
		return
	}

	if c.NotFound && c.Layout != "json" {
		c.LayoutMap["title"] = "404 not found"
		c.SendContentInLayout("404.html", nil, 404)
	} else if c.NotFound && c.Layout == "json" {
		c.SendContentAsJsonMessage("not found", 404)
	}
}
