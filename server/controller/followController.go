package controller

//helo
import (
	"example/social/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Follow(r *gin.Context) {

	follow := models.Follow{}
	follow.FollowerId = r.Param("id")
	follow.UserId = r.Query("user_id")
	follow.CreateAt = time.Now()
	if follow.FollowerId == "" || follow.UserId == "" {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing data to create follow",
		})
		return
	}

	tx := models.DB.Begin()

	//update follower of user
	updateUserInfo := tx.Model(models.UserInformation{}).Where("id = ?", follow.UserId).Update("followers", gorm.Expr("followers + ?", 1))
	if updateUserInfo.Error != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": updateUserInfo.Error.Error(),
		})
		return
	}
	if updateUserInfo.RowsAffected == 0 {
		tx.Rollback()
		r.JSON(http.StatusNotFound, gin.H{
			"error": "No User Found To Follow",
		})
		return
	}

	//update following of user
	if err := tx.Model(models.UserInformation{}).Where("id = ?", follow.FollowerId).Update("followings", gorm.Expr("followings + ?", 1)).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to save follow",
			"error":   err.Error(),
		})
		return
	}

	//Save follow to DB
	if err := tx.Create(follow).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to save follow",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UnFollow(r *gin.Context) {
	followerId := r.Param("id")
	userId := r.Query("user_id")

	tx := models.DB.Begin()

	//delete follow
	deleteFollowResult := tx.Where("follower_id = ? and user_id = ?", followerId, userId).Delete(models.Follow{})
	if deleteFollowResult.Error != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete follow",
			"error":   deleteFollowResult.Error.Error(),
		})
		return
	}
	if deleteFollowResult.RowsAffected == 0 {
		tx.Rollback()
		r.JSON(http.StatusNotFound, gin.H{
			"message": "No Follow Found",
		})
		return
	}

	//update follower for user
	if err := tx.Model(models.UserInformation{}).Where("id = ?", userId).Update("followers", gorm.Expr("followers + ?", -1)).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete follow",
			"error":   err.Error(),
		})
		return
	}

	//update following for follower
	if err := tx.Model(models.UserInformation{}).Where("id = ?", followerId).Update("followings", gorm.Expr("followings + ?", -1)).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete follow",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
