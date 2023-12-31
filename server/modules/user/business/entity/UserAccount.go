package entity

import (
	"errors"
	"example/social/common"
)

var(
	ErrorBlankID = errors.New("ID cannot be blank")
	ErrorBlankEmail = errors.New("Email cannot be blank")
	ErrorBlankPassword = errors.New("Password cannot be blank")
	ErrorIncorrectEmailOrPassword = errors.New("Email or password is incorrect")
)

type UserAccount struct {
	Id       string `gorm:"id"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	common.ObjectTime
}

func (userAccount *UserAccount) TableName() string {
	return "user_accounts"
}

