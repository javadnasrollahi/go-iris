package controllers

import (
	"manlogin/keynodes"

	"github.com/kataras/iris"
)

type FileController struct {
}

func (c *FileController) GetBy(uid string, ctx iris.Context, FServers map[string]string) {
	path := uid
	trueUrl := keynodes.DFSProtc + keynodes.DFS + "/file/" + path
	ctx.Header("Location", trueUrl)
	ctx.StatusCode(iris.StatusFound)
	//ctx.Redirect(trueUrl, iris.StatusFound)
}
func (c *FileController) OptionsBy(uid string, ctx iris.Context, FServers map[string]string) {}
