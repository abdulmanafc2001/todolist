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

func NewHandler(repo interfaces.Repository) *Handlers {
	mux := gin.Default()
	controllers := controllers.NewControllers(repo)
	registerRoutes(mux, controllers)
	return &Handlers{mux}
}

func registerRoutes(mux *gin.Engine, ctrls *controllers.Controller) {
	mux.LoadHTMLGlob("templates/*.html")
	mux.GET("/", ctrls.Home)
	mux.POST("/add-task", ctrls.CreateTodo)
	mux.GET("/loaderio-a2049cfeeeabeb2c6543b8655eb91a2e",ctrls.Check)
}

func (h *Handlers) Run() error {
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: h.Mux,
	}
	return srv.ListenAndServe()
}
