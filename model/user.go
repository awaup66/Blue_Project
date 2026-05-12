package model

import "gorm.io/gorm"

type User1 struct {
	gorm.Model
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Age      *int   `json:"age" binding:"required,min=1,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

func (User1) TableName() string {
	return "user1s"
}
