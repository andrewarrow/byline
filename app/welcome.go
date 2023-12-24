package app

import (
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
	if len(c.User) > 0 {
		handleSpaceIndex(c)
		return
	}
	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}
