package controller

import (
	"example/social/common"
	"example/social/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitUserInformation(r *gin.Context) {
	userInfoInit := models.InitUserInformation{}
	userInfoInit.Id = r.Param("id")
	if err := r.ShouldBind(&userInfoInit); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	userInfo := models.UserInformation{
		Id: userInfoInit.Id,
		UserName: userInfoInit.UserName,
		Birthday: userInfoInit.Birthday,
		AvatarImageUrl: userInfoInit.AvatarImageUrl,
		Followers: 0,
		Followings: 0,
		Friends: 0,
		UpdateAt: time.Now(),
	}
	if err := models.DB.Save(userInfo).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user_info": userInfo,
	})

}

func UpdateUserInformation(r *gin.Context) {
	userInfo := models.UpdateUserInformation{}
	userId := r.Param("id")
	if err := r.ShouldBind(&userInfo); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Where("id = ?", userId).Updates(userInfo).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})

}

func GetUserInformation(r *gin.Context) {
	userId := r.Param("id")
	field := r.Query("field")
	if field == "" {
		field = "user_name id birthday avatar_image_url create_at update_at followers followings"
	}
	fields := strings.Split(field, " ")

	user := models.UserInformation{}

	if err := models.DB.Select(fields).Where("id = ?", userId).Find(&user).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func GetBriefUserInformation(r *gin.Context) {
	userId := r.Param("id")
	field := "id user_name avatar_image_url"
	
	fields := strings.Split(field, " ")

	user := models.BriefUserInformation{}

	if err := models.DB.Select(fields).Where("id = ?", userId).Find(&user).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}



func GetFollowersInformation(r *gin.Context) {
	users := []models.BriefUserInformation{}
	userId := r.Param("id")
	paging := common.Paging{}

	

	if err := common.GetPaging(&paging, r); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Table("user_informations").InnerJoins("join follows on user_informations.id = follows.follower_id and follows.user_id = ?", userId).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&users).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data": users,
		"page": paging.Page,
		"limit": paging.Limit,
	})
}

func GetFollowingUsersInformation(r *gin.Context) {
	users := []models.BriefUserInformation{}
	userId := r.Param("id")
	paging := common.Paging{}

	if err := common.GetPaging(&paging, r); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Table("user_informations").InnerJoins("join follows on user_informations.id = follows.user_id and follows.follower_id = ?", userId).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&users).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
