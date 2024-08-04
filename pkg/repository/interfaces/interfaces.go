package interfaces

import "github.com/abdulmanafc2001/todolist/pkg/models"

type Repository interface {
	Create(models.Todo) error
	List() ([]models.Todo, error)
	Delete(string) error
}
