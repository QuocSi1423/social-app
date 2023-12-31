package controller

import (
	"example/social/common"
	"example/social/models"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePost(r *gin.Context) {
	postCreate := models.PostCreate{}

	postCreate.Id = uuid.NewString()
	postCreate.UserId = r.MustGet("id").(string)
	if err := r.ShouldBind(&postCreate); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing data or data is invalid",
			"error":   err.Error(),
		})
		return
	}

	post := models.Post{
		Id:            postCreate.Id,
		UserId:        postCreate.UserId,
		Title:         postCreate.Title,
		Description:   postCreate.Description,
		ImageUrl:      postCreate.ImageUrl,
		CountInteract: 0, // Hoặc giá trị mặc định khác nếu cần
		CountComment:  0, // Hoặc giá trị mặc định khác nếu cần
		CreateAt:      time.Now(),
	}

	if err := models.DB.Create(&post).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to save data to DB",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}

func DeletePost(r *gin.Context) {
	if err := models.DB.Where("id = ?", r.Param("post_id")).Delete(models.Post{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.JSON(http.StatusNotFound, gin.H{
				"message": "No Post Found",
				"error":   err.Error(),
			})
			return
		}
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete post",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdatePost(r *gin.Context) {
	post := models.Post{}

	if err := r.ShouldBind(&post); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Data is invalid",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Where("id = ?", r.Param("post_id")).Updates(post).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to update post",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func GetAPost(r *gin.Context) {
	post := models.Post{}

	if err := models.DB.Where("id = ?", r.Param("id")).First(&post).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get post",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data": post,
	})

}

func GetPosts(r *gin.Context) {
	posts := []models.Post{}

	userId := r.Query("user_id")
	paging := common.Paging{}

	if err := common.GetPaging(&paging, r); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if err := models.DB.Where("user_id = ?", userId).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&posts).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get posts",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"page":  paging.Page,
		"limit": paging.Limit,
	})

}

func GetUserInteractOfPost(r *gin.Context) {
	users := []models.BriefUserInformation{}
	post_id := r.Param("id")
	if err := models.DB.Select([]string{"id", "user_name", "avatar_image_url"}).Table("interacts").InnerJoins("join user_informations on interacts.user_id = user_informations.id and interacts.post_id = ?", post_id).Find(&users).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to load data",
			"error":   err.Error(),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetAnInteractOfPost(r *gin.Context) {
	userId := r.MustGet("id")
	postId := r.Param("id")

	Interact := models.Interact{}
	if err := models.DB.Where("user_id = ? and post_id = ?", userId, postId).First(&Interact).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.JSON(http.StatusNotFound, gin.H{
				"message": "No Interact Found",
				"error":   err.Error(),
			})
			return
		}
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to load data",
			"error":   err.Error(),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"data": Interact,
	})

}

func GetPostForUser(r *gin.Context) {
	user_id := r.MustGet("id")

	paging := common.Paging{}

	if err := common.GetPaging(&paging, r); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	posts := []models.Post{}
	if err := models.DB.Select("posts.*").Table("follows").InnerJoins("join posts on posts.user_id = follows.user_id and follows.follower_id = ?", user_id).Order("posts.create_at desc").Limit(10).Offset(paging.Page - 1).Find(&posts).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to load data",
			"error":   err.Error(),
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
