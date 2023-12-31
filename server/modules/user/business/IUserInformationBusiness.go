package business

import (
	"context"
	"example/social/common"
	"example/social/modules/user/business/entity"
)

type IUserInformationBusiness interface {
	InitUserInformation(ctx context.Context, data *entity.UserInformation) error
	ChangeUserInformation(ctx context.Context, data *entity.UserInformationForUpdate) error
	FindUserInformations(ctx context.Context, keyword string, paging common.Paging) ([]*entity.UserInformation, error)
}