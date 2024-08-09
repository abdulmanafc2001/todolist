package repository

import (
	"github.com/abdulmanafc2001/todolist/pkg/models"
	repo "github.com/abdulmanafc2001/todolist/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) repo.User {
	db.AutoMigrate(&models.User{})
	return &User{db}
}

func (u *User) Create(data models.User) error {
	db := u.DB.Create(&data)
	return db.Error
}

func (u *User) ListUser(username string) (models.User, error) {
	var user models.User
	db := u.DB.Where("user_name=?", username).First(&user)
	return user, db.Error
}
