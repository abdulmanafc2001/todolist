package interfaces

import "github.com/abdulmanafc2001/todolist/pkg/models"

type Repository interface {
	Create(models.Todo) error
	List() ([]models.Todo, error)
	ListWithUsername(string) ([]models.Todo,error)
	Delete(string) error
}

type User interface {
	Create(models.User) error
	ListUser(string) (models.User, error)
}
