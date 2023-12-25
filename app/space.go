package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleSpace(c *router.Context, second, third string) {
	//if router.NotLoggedIn(c) {
	//	return
	//}
	if second == "" && third == "" && c.Method == "GET" {
		handleSpaceIndex(c)
		return
	}
	if second == "vim" && third == "" && c.Method == "GET" {
		handleSpaceVim(c)
		return
	}
	c.NotFound = true
}

func handleSpaceIndex(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("space.html", send, 200)
}

func handleSpaceVim(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("vim.html", send, 200)
}
