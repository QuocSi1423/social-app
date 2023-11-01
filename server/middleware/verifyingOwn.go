package middleware
//helo
import (
	"example/social/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyPostOwn(r *gin.Context) {
	post := models.Post{}
	if err := models.DB.Where("id = ?", r.Param("post_id")).Select("user_id").First(&post).Error; err != nil {
		r.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "",
			"error": err.Error(),
		})
		return
	}
	
	allowed := r.MustGet("id") == post.UserId
	if !allowed {
		r.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You'r not allowed",
		})
		return
	}

	r.Next()
}
