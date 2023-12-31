package infrastucture

import (
	"context"
	"example/social/modules/user/business/entity"

	"gorm.io/gorm"
)

type userAccountStorage struct {
	db *gorm.DB
}

func NewUserAccountStorage(db *gorm.DB)userAccountStorage{
	return userAccountStorage{db}
}

func (store userAccountStorage)InsertUserAccount(ctx context.Context, data *entity.UserAccount) error{

	if err := store.db.Create(data).Error; err != nil{
		return err
	}
	
	return nil
}

func (store userAccountStorage)GetUserAccountByEmail(ctx context.Context, data *entity.UserAccount) error{
	if err := store.db.Where("email = ?", data.Email).First(&data).Error; err != nil{
		return err
	}
	return nil
}

func (store userAccountStorage) UpdateUserAccount(ctx context.Context, data *entity.UserAccount) error{
	if err := store.db.Table("user_accounts").Where("id = ?", data.Id).Update("password",data.Password).Error; err != nil{
		return err
	}
	return nil
}