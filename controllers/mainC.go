package controllers

import (
	"manlogin/keynodes"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type MainController struct {
}

func (c *MainController) Get(ctx iris.Context, ses *sessions.Session) {
	var UserUid string
	var Token string

	ctx.ViewData("BaseURL", keynodes.BaseURL)
	ctx.ViewData("HttpProtc", keynodes.HttpProtc)
	ctx.ViewData("UserUid", UserUid)
	ctx.ViewData("Token", Token)
	ctx.ViewData("Service", strings.TrimRight(keynodes.HttpProtc+keynodes.BaseURL, "/"))
	ctx.ViewData("Data", "")
	ctx.View("tmp/head.html")
	ctx.View("tmp/footer.html")
	return
}
func (c *MainController) GetTest(ctx iris.Context) {
	type Person struct {
		Name string
		Age  int
	}
	ctx.ViewData("people", []Person{Person{
		Name: "hamid",
		Age:  15,
	},
		Person{
			Name: "javad",
		}, Person{
			Name: "Mahdi",
			Age:  20,
		},
	})
	ctx.View("tmp/head.html")
	ctx.View("body.handlebars")
	ctx.View("tmp/footer.html")

}
