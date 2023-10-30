package controller

import (
	"example/social/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateInteract(r *gin.Context) {
	interact := models.Interact{}

	interact.PostId = r.Query("post_id")
	interact.UserId = r.Param("id")
	interact.CreateAt = time.Now()

	tx := models.DB.Begin()

	//update interacr of post
	updatePost := tx.Model(models.Post{}).Where("id = ?", interact.PostId).Update("count_interact", gorm.Expr("count_interact + ?", 1))
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

	//Save interact to DB
	if err := tx.Create(interact).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to save interact",
			"error":   err.Error(),
		})
		return
	}
	tx.Commit()
	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func DeleteInteract(r *gin.Context) {
	interact := models.Interact{}

	interact.PostId = r.Query("post_id")
	interact.UserId = r.Param("id")
	tx := models.DB.Begin()

	//delete interact
	deleteInteractResult := tx.Where("post_id = ? and user_id = ?", interact.PostId, interact.UserId).Delete(interact)
	if deleteInteractResult.Error != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete interact",
			"error":   deleteInteractResult.Error.Error(),
		})
		return
	}
	if deleteInteractResult.RowsAffected == 0 {
		tx.Rollback()
		r.JSON(http.StatusNotFound, gin.H{
			"message": "Interact not found",
		})
		return
	}

	//update interact of post
	if err := tx.Model(models.Post{}).Where("id = ?", interact.PostId).Update("count_interact", gorm.Expr("count_interact + ?", -1)).Error; err != nil {
		tx.Rollback()
		r.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to update interact of post",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()

	r.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

