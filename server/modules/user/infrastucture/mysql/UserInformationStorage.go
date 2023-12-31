package infrastucture

import (
	"context"
	"example/social/common"
	"example/social/modules/user/business/entity"

	"gorm.io/gorm"
)

type userInformationStorage struct {
	db *gorm.DB
}

func NewUserInformationStorage(db *gorm.DB) userInformationStorage{
	return userInformationStorage{db: db}
}

func (store userInformationStorage)InsertUserInformation(ctx context.Context, data *entity.UserInformation) error{
	
	if err := store.db.Create(&data).Error; err != nil{
		return err
	}
	return nil

}

func (store userInformationStorage)UpdataUserInformation(ctx context.Context, data *entity.UserInformationForUpdate) error{
	
	if err := store.db.Save(data).Error; err != nil{
		return err
	}

	return nil
	
}

func (store userInformationStorage)FindUserInformations(ctx context.Context, keyword string, paging common.Paging) ([]*entity.UserInformation, error){
	
	userInformations := make([]*entity.UserInformation, paging.Limit)

	if err := store.db.Where("name like ? or user_name like ?", "%" + keyword + "%", "%" + keyword + "%" ).Offset((paging.Page-1)*paging.Limit).Limit(paging.Limit).Find(userInformations).Error; err != nil{
		return nil, err
	}

	return userInformations, nil

}

