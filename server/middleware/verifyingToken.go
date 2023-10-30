package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(r *gin.Context) {

	if r.GetHeader("token") == "" {
		r.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You'r not authenticated",
		})
		return
	}

	tokenStr := r.GetHeader("token")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		r.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "token is not valid",
		})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"]
	r.Set("id", id)
	if time.Unix(int64(claims["expired"].(float64)),0).Before(time.Now()){
		r.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token has expired",
		})
		return
	}
	r.Next()


}

func VerifyIdentity(r *gin.Context) {

	allowed := r.MustGet("id") == r.Param("id")
	if !allowed {
		r.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You'r not allowed",
		})
		return
	}

	defer r.Next()
}
