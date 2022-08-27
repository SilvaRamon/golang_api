package controllers

import (
	"hello/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.IService
	router  *gin.Engine
}

func Initialize(service service.IService, router *gin.Engine) *controller {
	return &controller{service, router}
}

func (c *controller) ExposeRoutes() {
	c.router.GET("/users/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			c.service.GetUsers(),
		)
	})

	c.router.GET("/user/:mail", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			c.service.GetUser(ctx.Param("mail")),
		)
	})
}
