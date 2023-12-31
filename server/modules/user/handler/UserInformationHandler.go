package handler

import (
	"errors"
	"example/social/common"
	"example/social/modules/user/business"
	"example/social/modules/user/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type userInformationHandler struct {
	business business.IUserInformationBusiness
}
func NewUserInformationHandler(business business.IUserInformationBusiness) userInformationHandler{
	return userInformationHandler{business: business}
}

func (handler userInformationHandler)InitUserInformation() func(c *gin.Context){
	return func(c *gin.Context) {

		userInformation := entity.UserInformation{}

		if err := c.ShouldBind(&userInformation); err != nil{
			c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
			return
		}

		if err := handler.business.InitUserInformation(c.Request.Context(), &userInformation); err != nil{

			var sqlErr *mysql.MySQLError

			//id người dùng không tồn tại 
			if errors.As(err, &sqlErr) && sqlErr.Number == 1452{
				c.JSON(http.StatusNotFound, common.NewErrorRespone(err.Error()))
				return
			}

			//thông tin người dùng đã được khởi tạo 
			if errors.As(err, &sqlErr) && sqlErr.Number == 1062{
				c.JSON(http.StatusConflict, common.NewErrorRespone(err.Error()))
				return
			}

			c.JSON(http.StatusInternalServerError, common.NewErrorRespone(err.Error()))
				return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "sucess",
			"information": userInformation,
		})

	}
}

func (handler userInformationHandler) ChangeUserInformation() func(c *gin.Context){
	return func (c* gin.Context)  {
		userInformation := entity.UserInformationForUpdate{}

		if err := c.ShouldBind(&userInformation); err != nil{
			c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
			return
		}

		if err := handler.business.ChangeUserInformation(c.Request.Context(), &userInformation); err != nil{
			if err == entity.ErrorBlankName || err == entity.ErrorBlankUserName || err == entity.ErrorInvalidBirthday || err == entity.ErrorInvalidFormattingUserName{
				c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewErrorRespone(err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "sucess",
			"user_information": userInformation,
		})
	}
}



