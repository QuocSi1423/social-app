package models

import "time"

type UserInformation struct {
	Id             string     `json:"id" gorm:"id"`
	UserName       string     `json:"user_name" gorm:"user_name"`
	AvatarImageUrl string     `json:"avatar_image_url" gorm:"avatar_image_url"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	Followers      int        `json:"followers" gorm:"followers"`
	Followings     int        `json:"following" gorm:"following"`
	Friends        int        `json:"friends" gorm:"friends"`
	UpdateAt       time.Time `json:"update_at" gorm:"update_at"`
}

type InitUserInformation struct {
	Id             string     `json:"id" gorm:"id"`
	UserName       string     `json:"user_name" gorm:"user_name"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	AvatarImageUrl string     `json:"avatar_image_url" gorm:"avatar_image_url"`
}

type UpdateUserInformation struct {
	UserName       string     `json:"user_name" gorm:"user_name"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	AvatarImageUrl string     `json:"avatar_image_url" gorm:"avatar_image_url"`
}

type BriefUserInformation struct {
	Id             string `json:"id" gorm:"id"`
	UserName       string `json:"user_name" gorm:"user_name"`
	AvatarImageUrl string `json:"avatar_image_url" gorm:"avatar_image_url"`
}

type FollowerUserInformation struct {
	Followers int `json:"followers" gorm:"followers"`
}

type FollowingUserInformation struct {
	Followings int `json:"following" gorm:"following"`
}

func (userInfo UserInformation) TableName() string {
	return "user_informations"
}

func (userInfo InitUserInformation) TableName() string {
	return "user_informations"
}

func (userInfo UpdateUserInformation) TableName() string {
	return "user_informations"
}

func (userInfo BriefUserInformation) TableName() string {
	return "user_informations"
}

func (userInfo FollowerUserInformation) TableName() string {
	return "user_informations"
}

func (userInfo FollowingUserInformation) TableName() string {
	return "user_informations"
}
