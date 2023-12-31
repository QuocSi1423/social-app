package handler

import (
	"errors"
	"example/social/common"
	"example/social/modules/user/business"
	"example/social/modules/user/business/entity"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userAccountHandler struct {
	business business.IUserAccountBusiness
}

func NewUsUserAccountHandler(business business.IUserAccountBusiness) userAccountHandler {
	return userAccountHandler{business: business}
}

func hashPassword(userAccount *entity.UserAccount){
			cost := bcrypt.DefaultCost
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userAccount.Password), cost)
			userAccount.Password = string(hashPassword)
}

func (handler userAccountHandler) Register() func(c *gin.Context) {

	return func(c *gin.Context) {

		userAccount := entity.UserAccount{}

		//Lấy dữ liệu từ req
		if err := c.ShouldBind(&userAccount); err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
			return
		}

		userAccount.Id = uuid.NewString()

		//thiết lập thời gian
		userAccount.CreateAt = time.Now()
		userAccount.UpdateAt = time.Now()

		//hash password
		if userAccount.Password != "" {
			hashPassword(&userAccount)
		}

		//kiểm tra dữ liệu và lưu vào database
		if err := handler.business.Register(c.Request.Context(), &userAccount); err != nil {

			var sqlErr *mysql.MySQLError
			if errors.As(err, &sqlErr) && sqlErr.Number == 1062 {
				c.JSON(http.StatusConflict, common.NewErrorRespone(err.Error()))
				return
			}

			if err == entity.ErrorBlankEmail || err == entity.ErrorBlankPassword {
				c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
				return
			}

			c.JSON(http.StatusInternalServerError, common.NewErrorRespone(err.Error()))
			return

		}

		//đăng kí thành công
		c.JSON(http.StatusConflict, gin.H{
			"message": "success",
			"account":    userAccount,
		})

	}

}

func (handler userAccountHandler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		userAccount := entity.UserAccount{}

		if err := c.ShouldBind(&userAccount); err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
			return
		}

		if err := handler.business.Login(c.Request.Context(), &userAccount); err != nil {

			if err == gorm.ErrRecordNotFound || err == entity.ErrorIncorrectEmailOrPassword {
				c.JSON(http.StatusUnauthorized, common.NewErrorRespone(entity.ErrorIncorrectEmailOrPassword.Error()))
				return
			}

			if err == entity.ErrorBlankEmail || err == entity.ErrorBlankPassword {
				c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
				return
			}

			c.JSON(http.StatusInternalServerError, common.NewErrorRespone(err.Error()))
			return
		}

		// generate jwt
		t := time.Now().Add(time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":      userAccount.Id,
			"expired": t.Unix(),
		})

		tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

		//tạo token thất bại
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create token",
				"error":   err.Error(),
			})
			return
		}

		//đăng nhập thành công
		c.JSON(http.StatusOK, gin.H{
			"message": "sucess",
			"token": tokenStr,
			"account": userAccount,
		})

	}
}

func (handler userAccountHandler) ChangePassword() func(c *gin.Context){
	return func(c *gin.Context) {
		userAccount := entity.UserAccount{}

		if err := c.ShouldBind(&userAccount); err != nil{
			c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
			return 
		}

		userAccount.Id = c.Param("id")

		//hash password
		if userAccount.Password != "" {
			hashPassword(&userAccount)
		}

		if err := handler.business.ChangePassword(c.Request.Context(), &userAccount); err != nil{
			if err == entity.ErrorBlankPassword{
				c.JSON(http.StatusBadRequest, common.NewErrorRespone(err.Error()))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewErrorRespone(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "sucess",
		})

	}
}

func (handler userAccountHandler) SetUpRoutes(group *gin.RouterGroup){
	group.POST("/user_accounts", handler.Register())
	group.POST("/login", handler.Login())

	//use middleware
	
	group.PATCH("/user_accounts/:id", handler.ChangePassword())
}
