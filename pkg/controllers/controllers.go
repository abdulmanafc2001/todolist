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
	User interfaces.User
}

func NewControllers(repo interfaces.Repository, usr interfaces.User) *Controller {
	return &Controller{repo, usr}
}

func (c *Controller) Home(ctx *gin.Context) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ok, userName := UserLoged(ctx)
	fmt.Println("Checking: ", ok, userName)
	if !ok {
		ctx.Redirect(303, "/login")
		return
	}
	todos, err := c.Repo.ListWithUsername(userName)
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

	_, userName := UserLoged(ctx)

	todo := models.Todo{
		TaskNumber:  number,
		Description: description,
		Completed:   "No",
		DayCount:    count,
		UserName:    userName,
	}

	if err := c.Repo.Create(todo); err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect(303, "/")
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	taskNo := ctx.Param("number")
	if err := c.Repo.Delete(taskNo); err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	ctx.Redirect(303, "/")
}
