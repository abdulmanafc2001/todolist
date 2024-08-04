package controllers

import (
	"fmt"
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

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	taskNo := ctx.Param("number")
	fmt.Println(taskNo)
	if err := c.Repo.Delete(taskNo); err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	ctx.Redirect(303, "/")
}
