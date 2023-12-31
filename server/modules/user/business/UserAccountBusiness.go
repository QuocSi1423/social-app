package business

import (
	"context"
	"example/social/modules/user/business/entity"

	"golang.org/x/crypto/bcrypt"
)

type userAccountBusiness struct {
	store IUserAccountRepository
}

func NewUserAccountBusiness(store IUserAccountRepository) userAccountBusiness{
	return userAccountBusiness{store: store}
}

//đăng kí tài khoản
func (business userAccountBusiness)Register(ctx context.Context, data *entity.UserAccount) error{
	
	if data.Email == ""{
		return entity.ErrorBlankEmail
	}

	if data.Password == ""{
		return entity.ErrorBlankPassword
	}

	if err := business.store.InsertUserAccount(ctx, data); err != nil{
		return err
	}

	return nil
}

//đăng nhập
func (business userAccountBusiness)Login(ctx context.Context, data *entity.UserAccount) error{

	if data.Email == ""{
		return entity.ErrorBlankEmail
	}

	if data.Password == ""{
		return entity.ErrorBlankPassword
	}

	password := data.Password

	if err := business.store.GetUserAccountByEmail(ctx, data); err != nil{
		return err
	}

	//kiểm tra mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password)); err != nil{
		return entity.ErrorIncorrectEmailOrPassword
	}

	return nil

}

func (business userAccountBusiness) ChangePassword(ctx context.Context, data *entity.UserAccount) error{

	if data.Password == ""{
		return entity.ErrorBlankPassword
	}

	if err := business.store.UpdateUserAccount(ctx, data); err != nil{
		return err
	}


	return nil
}

