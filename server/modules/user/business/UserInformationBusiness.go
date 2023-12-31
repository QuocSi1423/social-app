package business

import (
	"context"
	"errors"
	"example/social/common"
	"example/social/modules/user/business/entity"
	"regexp"
	"time"
)

type userInformationBusiness struct {
	store IUserInformationRepository
}

func NewUserInformationBusiness(store IUserInformationRepository) userInformationBusiness{
	return userInformationBusiness{store: store}
}


func IsInvalidFormattingUserName(s string) bool{
	regexPattern := `^@[\w\d]+$`
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(s)

}

func (business userInformationBusiness) InitUserInformation(ctx context.Context, data *entity.UserInformation) error{
	
	if data.Id == ""{
		return entity.ErrorBlankID
	}

	if data.UserName == ""{
		return entity.ErrorBlankUserName
	}

	if time.Now().Compare(data.Birthday) <= 0{
		return entity.ErrorInvalidBirthday
	}

	if err := business.store.InsertUserInformation(ctx, data); err != nil{
		return err
	}

	return nil
}

func (business userInformationBusiness) ChangeUserInformation(ctx context.Context, data *entity.UserInformationForUpdate) error{

	if data.UserName == ""{
		return entity.ErrorBlankUserName
	}

	if IsInvalidFormattingUserName(data.UserName){
		return entity.ErrorInvalidFormattingUserName
	}

	if data.Name == ""{
		return entity.ErrorBlankName
	}

	if time.Now().Compare(data.Birthday) <= 0{
		return entity.ErrorInvalidBirthday
	}

	if err := business.store.UpdataUserInformation(ctx, data); err != nil{
		return err
	}

	return nil
}

func (business userInformationBusiness) FindUserInformations(ctx context.Context, keyword string, paging common.Paging) ([]*entity.UserInformation, error){

	if keyword == ""{
		return nil, errors.New("Cannot find")
	}
	userInformation, err := business.store.FindUserInformations(ctx, keyword, paging)

	if err == nil{
		return nil, err
	}

	return userInformation, nil

}