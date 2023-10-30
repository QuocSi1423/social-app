package controller

import (
	"example/social/common"
	"example/social/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// comment
func CreateComment(r *gin.Context) {

	comment := models.Comment{}

	if err := r.ShouldBind(&comment); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing data or data is invalid",
			"error":   err.Error(),
		})
		return
	}

	comment.CountReplyComment = 0
	comment.CreateAt = time.Now()
	comment.UpdateAt = time.Now()
	comment.Id = uuid.NewString()
	comment.UserId = r.Param("id")
	comment.PostId = r.Query("post_id")
	if comment.PostId == "" {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing data or data is invalid",
			"error":   "Missing post id",
		})
		return
	}
	tx := models.DB.Begin()

	//Update count comment of post
	updatePost := tx.Model(models.Post{}).Where("id = ?", comment.PostId).Update("count_comment", gorm.Expr("count_comment + ?", 1))
	if updatePost.Error != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to update comment of post",
			"error":   updatePost.Error.Error(),
		})
		return
	}
	if updatePost.RowsAffected == 0 {
		tx.Rollback()
		r.JSON(http.StatusNotFound, gin.H{
			"message": "Post Not Found",
		})
		return
	}

	//Save comment to DB
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to save comment",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
		"comment": comment,
	})

}

func DeleteComment(r *gin.Context) {
	comment := models.Comment{Id: r.Param("comment_id")}
	tx := models.DB.Begin()

	//Get comment
	if err := models.DB.First(&comment).Error; err != nil {
		if err.Error() == "record not found" {
			r.JSON(http.StatusNotFound, gin.H{
				"message": "Comment Not Found",
			})
			return
		}
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get comment",
			"error":   err.Error(),
		})
		return

	}

	//Delete comment
	if err := tx.Delete(comment).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete comment",
			"error":   err.Error(),
		})
		return
	}

	//Update comment of post
	if err := tx.Model(models.Post{}).Where("id = ?", comment.PostId).Update("count_comment", gorm.Expr("count_comment + ?", -1)).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to update comment of post",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdateComment(r *gin.Context) {
	comment := models.CommentUpdate{}

	if err := r.ShouldBind(&comment); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Data comment is invalid",
			"error":   err.Error(),
		})
		return
	}

	comment.UpdateAt = time.Now()
	comment.Id = r.Param("comment_id")

	if err := models.DB.Updates(comment).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to update comment",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    comment,
	})

}

func GetCommentsOfPost(r *gin.Context) {
	paging := common.Paging{}

	if err := common.GetPaging(&paging, r); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Input is invalid",
			"error":   err.Error(),
		})
		return
	}

	comments := []models.Comment{}

	if err := models.DB.Where("post_id = ?", r.Param("id")).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&comments).Error; err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get comment",
			"error":   err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"data":  comments,
		"limit": paging.Limit,
		"page":  paging.Page,
	})
}