package repository

import (
	"log"

	"github.com/abdulmanafc2001/todolist/pkg/models"
	repo "github.com/abdulmanafc2001/todolist/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type Todo struct {
	DB *gorm.DB
}

func NewTodo(db *gorm.DB) repo.Repository {
	db.AutoMigrate(&models.Todo{})
	return &Todo{db}
}

func (r *Todo) Create(todo models.Todo) error {
	db := r.DB.Create(&todo)
	if db.Error != nil {
		log.Println(db.Error)
		return db.Error
	}
	return nil
}

func (r *Todo) List() ([]models.Todo, error) {
	var todos []models.Todo
	db := r.DB.Find(&todos)
	return todos, db.Error
}

func (r *Todo) ListWithUsername(username string) ([]models.Todo, error) {
	var todos []models.Todo
	db := r.DB.Where("user_name = ?", username).Find(&todos)
	return todos, db.Error
}

func (r *Todo) Delete(number string) error {
	db := r.DB.Where("task_number = ?", number).Delete(&models.Todo{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
