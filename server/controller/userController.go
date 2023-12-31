package controller

import (
	"example/social/models"
	"fmt"
	"net/http"
	"os"

	"gorm.io/gorm"

	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(r *gin.Context) {
	user := models.RegisterUser{}
	if err := r.ShouldBind(&user); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	emailExist := false
	nameExist := false
	if err := models.DB.Where("email = ?", user.Email).First(&models.User{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			r.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
	} else {
		emailExist = true
	}

	if err := models.DB.Where("id = ?", user.Id).First(&models.User{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			r.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
	} else {
		nameExist = true
	}

	if emailExist || nameExist {
		r.JSON(http.StatusConflict, gin.H{
			"emailExist": emailExist,
			"nameExist":  nameExist,
		})
		return
	}

	cost := bcrypt.DefaultCost
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	user.Password = string(hashPassword)

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Save(user).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    models.RegisterUser{Id: user.Id},
	})
}

func Login(r *gin.Context) {
	clientData := models.LoginUser{}

	if err := r.ShouldBind(&clientData); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	dbData := models.User{}

	if err := models.DB.Where("id = ?", clientData.Id).First(&dbData).Error; err != nil {
		
		if err == gorm.ErrRecordNotFound {
			r.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(clientData.Password)); err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	t := time.Now().Add(time.Hour)
	fmt.Println(t)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      dbData.Id,
		"expired": t.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create token",
			"error":   err.Error(),
		})
		return
	}

	r.SetSameSite(http.SameSiteNoneMode)
	r.SetCookie("access_token", tokenStr, 3600, "/", "localhost", true, false)
	r.JSON(http.StatusAccepted, gin.H{
		"message": "success",
		"data":    dbData,
		"token":   tokenStr,
	})

}

