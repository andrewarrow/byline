package app

import (
	"io/ioutil"

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
	if second == "save" && third == "" && c.Method == "POST" {
		handleSpaceSave(c)
		return
	}
	if second == "load" && third != "" && c.Method == "GET" {
		handleSpaceLoad(c, third)
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

func handleSpaceSave(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	lines := c.Params["lines"].(string)
	ioutil.WriteFile("index.mu", []byte(lines), 0644)
	c.SendContentAsJson("", 200)
}
func handleSpaceLoad(c *router.Context, id string) {
	send := map[string]any{}
	b, _ := ioutil.ReadFile("index.mu")
	send["lines"] = string(b)
	c.SendContentAsJson(send, 200)
}
