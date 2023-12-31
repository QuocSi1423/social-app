package business

import (
	"context"
	"example/social/modules/user/business/entity"
)

type IUserAccountRepository interface { 

	InsertUserAccount(ctx context.Context, data *entity.UserAccount) error
	GetUserAccountByEmail(ctx context.Context, data *entity.UserAccount) error
	UpdateUserAccount(ctx context.Context, data *entity.UserAccount) error
}

