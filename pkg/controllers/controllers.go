package controllers

import (
	"log"
	"strconv"

	"github.com/abdulmanafc2001/todolist/pkg/models"
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

func (c *Controller) CreateTodo(ctx *gin.Context) {
	number, _ := strconv.Atoi(ctx.Request.FormValue("task-number"))
	description := ctx.Request.FormValue("description")
	count, _ := strconv.Atoi(ctx.Request.FormValue("day-count"))

	todo := models.Todo{
		TaskNumber:  number,
		Description: description,
		Completed:   "No",
		DayCount:    count,
	}

	if err := c.Repo.Create(todo); err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect(303, "/")
}

func (c *Controller) Check(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"hello":"hello",
	})
}
