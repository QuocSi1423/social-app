package models

import "time"

type User struct {
	Id       string    `json:"id" gorm:"id"`
	Email    string    `json:"email" gorm:"email"`
	Password string    `json:"password" gorm:"password"`
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
}

type UserCheckLoginName struct{
	Id string `json:"id"`
}

type RegisterUser struct {
	Id       string `json:"id" gorm:"id"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

type LoginUser struct {
	Id       string `json:"id" gorm:"id"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *RegisterUser) TableName() string {
	return "users"
}

func (user *LoginUser) TableName() string {
	return "users"
}
