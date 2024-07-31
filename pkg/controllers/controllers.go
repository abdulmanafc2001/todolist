package controllers

import (
	"log"

	"github.com/abdulmanafc2001/todolist/pkg/repository/interfaces"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Repo interfaces.Repository
}

func NewControllers(repo interfaces.Repository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Home(ctx *gin.Context) {
	todos, err := c.Repo.List()
	if err != nil {
		log.Println(err)
		return
	}
	ctx.HTML(200, "index.html", gin.H{
		"todo": todos,
	})
}
