package main

import (
	"example/social/controller"
	"example/social/middleware"
	"example/social/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "github.com/gin-contrib/cors"
)

func main() {
	godotenv.Load()
	models.Conn_DB()
	
	r := gin.Default()
	config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"} // Cho phép truy cập từ tất cả các nguồn
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    r.Use(cors.New(config))
	v1 := r.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)

			information := user.Group("/:id/informations")
			{
				information.GET("/", controller.GetUserInformation)
				information.GET("/followers", controller.GetFollowersInformation)
				information.GET("/followings", controller.GetFollowingUsersInformation)
				information.GET("/brief", controller.GetBriefUserInformation)

				information.Use(middleware.VerifyJWT,middleware.VerifyIdentity)
				
				information.POST("/", controller.InitUserInformation)
				information.PATCH("/", controller.UpdateUserInformation)
			}

			follow := user.Group("/:id/following")
			{


				information.Use(middleware.VerifyJWT,middleware.VerifyIdentity)

				follow.POST("/", controller.Follow)
				follow.DELETE("/", controller.UnFollow)
			}
			post := user.Group("/:id/posts")
			{
				post.Use(middleware.VerifyJWT,middleware.VerifyIdentity)

				post.POST("/", controller.CreatePost)

				post.Use(middleware.VerifyPostOwn)

				post.DELETE("/:post_id", controller.DeletePost)
				post.PATCH("/:post_id", controller.UpdatePost)
			}
			comment := user.Group("/:id/comments")
			{
				comment.Use(middleware.VerifyJWT,middleware.VerifyIdentity)
				comment.POST("/", controller.CreateComment)
				comment.DELETE("/:comment_id", controller.DeleteComment)
				comment.PATCH("/:comment_id", controller.UpdateComment)
			}
			interact := user.Group(":id/likes")
			{
				interact.Use(middleware.VerifyJWT,middleware.VerifyIdentity)

				interact.POST("/", controller.CreateInteract)
				interact.DELETE("/", controller.DeleteInteract) 
			}

		}

		post := v1.Group("/posts")
		{
			post.GET("/:id", controller.GetAPost)
			post.GET("/", controller.GetPosts)
			post.GET("/:id/comments", controller.GetCommentsOfPost)
			post.GET("/:id/interacts", controller.GetUserInteractOfPost)
		}
	}

	r.Run()

}
