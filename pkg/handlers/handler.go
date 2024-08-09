package handlers

import (
	"net/http"
	"os"

	"github.com/abdulmanafc2001/todolist/pkg/controllers"
	"github.com/abdulmanafc2001/todolist/pkg/repository/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Handlers struct {
	Mux *gin.Engine
}



func NewHandler(repo interfaces.Repository, user interfaces.User) *Handlers {
	mux := gin.Default()
	mux.Use(func(c *gin.Context) {
		session, err := controllers.Store.Get(c.Request, "jwt_token")
		if err == nil {
			c.Set("session", session)
		}
		c.Next()
	})
	controllers := controllers.NewControllers(repo, user)
	registerRoutes(mux, controllers)
	return &Handlers{mux}
}

func registerRoutes(mux *gin.Engine, ctrls *controllers.Controller) {
	mux.LoadHTMLGlob("templates/*.html")
	mux.GET("/", ctrls.Home)
	mux.POST("/add-task", ctrls.CreateTodo)
	mux.GET("/delete/:number", ctrls.DeleteTodo)
	mux.GET("/login", ctrls.LoginRender)
	mux.GET("/signup", ctrls.SignupRender)

	mux.POST("/signup", ctrls.Signup)
	mux.POST("/login", ctrls.Login)

	mux.GET("/logout",ctrls.Logout)
}

func (h *Handlers) Run() error {
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: h.Mux,
	}
	return srv.ListenAndServe()
}
