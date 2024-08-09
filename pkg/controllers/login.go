package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/abdulmanafc2001/todolist/helper"
	"github.com/abdulmanafc2001/todolist/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

func (c *Controller) LoginRender(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}

func (c *Controller) SignupRender(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", nil)
}

func (c *Controller) Signup(ctx *gin.Context) {
	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	user, _ := c.User.ListUser(username)

	if user.UserName == username {
		ctx.HTML(400, "signup.html", gin.H{
			"message": "this username already exist",
		})
		return
	}

	if username == "" {
		ctx.HTML(400, "signup.html", gin.H{
			"message": "empty username or password try again",
		})
		return
	}
	if password == "" {
		ctx.HTML(400, "signup.html", gin.H{
			"message": "empty username or password try again",
		})
		return
	}
	pass, err := helper.HashPassword(password)
	if err != nil {
		ctx.HTML(400, "signup.html", gin.H{
			"message": "failed to hash password",
		})
		return
	}

	user = models.User{
		UserName: username,
		Password: pass,
	}
	fmt.Println(user)
	err = c.User.Create(user)
	if err != nil {
		ctx.HTML(400, "signup.html", gin.H{
			"message": "create user failed",
		})
		return
	}
	ctx.Redirect(303, "/login")
}

func (c *Controller) Login(ctx *gin.Context) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")

	user, err := c.User.ListUser(username)
	if err != nil {
		ctx.HTML(400, "login.html", gin.H{
			"message": "incorrect username or password",
		})
		return
	}

	verifyPassword := helper.VerifyPassword(password, user.Password)
	if !verifyPassword {
		ctx.HTML(400, "login.html", gin.H{
			"message": "incorrect username or password",
		})
		return
	}
	token, err := helper.CreateToken(username, "User")
	if err != nil {
		ctx.HTML(500, "login.html", gin.H{
			"message": "something failed",
		})
		return
	}
	session, err := Store.Get(ctx.Request, "jwt_token")
	if err != nil {
		ctx.HTML(500, "login.html", gin.H{
			"message": "failed to create session",
		})
		return
	}
	session.Values["token"] = token
	session.Values["user"] = username
	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		ctx.HTML(500, "login.html", gin.H{
			"message": "failed to save session",
		})
		return
	}

	ctx.Redirect(303, "/")
}

func UserLoged(c *gin.Context) (bool, string) {
	session, _ := c.Get("session")
	if session == nil {
		return false, ""
	}

	sess, ok := session.(*sessions.Session)
	if !ok {
		return false, ""
	}

	token, tokenOk := sess.Values["token"].(string)
	user, userOk := sess.Values["user"].(string)


	if !tokenOk || !userOk {
		return false, ""
	}
	 _ = token

	return true, user
}

func (c *Controller) Logout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("jwt_token")
	if err != nil {
		ctx.Redirect(303, "/login")
	}
	ctx.SetCookie("jwt_token", "", -1, "/", "localhost", false, true)
	_ = cookie
	ctx.Redirect(http.StatusSeeOther, "/login")
}
