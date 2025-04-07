package main

import (
	"github.com/a-h/templ"
	"github.com/dev-sandip/go-auth/templates/pages"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	return template.Render(c.Request.Context(), c.Writer)
}

func RenderLoginPage(c *gin.Context) {
	component := pages.Login()
	render(c, 200, component)
}
func RenderRegisterPage(c *gin.Context) {
	component := pages.Register()
	render(c, 200, component)
}

func RenderHomePage(c *gin.Context) {
	component := pages.Index()
	render(c, 200, component)
}
func RenderIndexPage(c *gin.Context) {
	component := pages.Home()
	render(c, 200, component)
}
