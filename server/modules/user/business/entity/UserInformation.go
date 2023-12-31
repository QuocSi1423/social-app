package entity

import (
	"errors"
	"time"
)

var (
	ErrorBlankUserName   = errors.New("UserName cannot be blank")
	ErrorInvalidFormattingUserName = errors.New("UserName's Fotmatting is invalid")
	ErrorBlankName = errors.New("Name cannot be blank")
	ErrorInvalidBirthday = errors.New("Birthday cannot after current")
)

type UserInformation struct {
	Id             string    `json:"id" gorm:"id"`
	UserName       string    `json:"user_name" gorm:"user_name"`
	Name           string    `json:"name" gorm:"name"`
	AvatarImageUrl string    `json:"avatar_image_url" gorm:"avatar_image_url"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	Followers      int       `json:"followers" gorm:"followers"`
	Followings     int       `json:"following" gorm:"following"`
	Friends        int       `json:"friends" gorm:"friends"`
}

type InitUserInformation struct {
	Id             string    `json:"id" gorm:"id"`
	Name           string    `json:"name" gorm:"name"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	AvatarImageUrl string    `json:"avatar_image_url" gorm:"avatar_image_url"`
}

type UserInformationForUpdate struct {
	UserName       string    `json:"user_name" gorm:"user_name"`
	Name           string    `json:"name" gorm:"name"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	AvatarImageUrl string    `json:"avatar_image_url" gorm:"avatar_image_url"`
}

type BriefUserInformation struct {
	Id             string `json:"id" gorm:"id"`
	UserName       string `json:"user_name" gorm:"user_name"`
	Name           string `json:"name" gorm:"name"`
	AvatarImageUrl string `json:"avatar_image_url" gorm:"avatar_image_url"`
}

func (userInformation UserInformation) TableName() string {
	return "user_informations"
}

func (userInformation InitUserInformation) TableName() string {
	return "user_informations"
}
