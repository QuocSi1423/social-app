package business

import (
	"context"
	"example/social/modules/user/business/entity"
)

type IUserAccountBusiness interface {
	
	Login(ctx context.Context, data *entity.UserAccount) error
	Register(ctx context.Context, data * entity.UserAccount) error
	ChangePassword(ctx context.Context, data *entity.UserAccount) error

}