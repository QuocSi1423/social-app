package main

import (
	"example/social/models"
	"example/social/modules/user/business"
	"example/social/modules/user/handler"
	infrastucture "example/social/modules/user/infrastucture/mysql"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	models.Conn_DB()
	r := gin.Default()
	

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Fatal(err)
	}



	//setup dependencies
	userAccountStorage := infrastucture.NewUserAccountStorage(db)
	userAccountBusiness := business.NewUserAccountBusiness(userAccountStorage)
	userAccountHandler := handler.NewUsUserAccountHandler(userAccountBusiness)

	// userInformationStorage := infrastucture.NewUserInformationStorage(db)
	// userInformationBusiness := business.NewUserInformationBusiness(userInformationStorage)
	// userInformationHandler := handler.NewUserInformationHandler(userInformationBusiness)

	
	cors_ := cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5173"} ,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin, Authorization, Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(cors_))
	v1 := r.Group("/v1")

	userAccountHandler.SetUpRoutes(v1)

	// {
	// 	v1.POST("/register", userAccountHandler.Register())
	// 	v1.POST("/login",  userAccountHandler.Login())

	// 	follow := v1.Group("/following")
	// 	{
	// 		follow.Use(middleware.VerifyJWT)

	// 		follow.POST("/", controller.Follow)
	// 		follow.DELETE("/", controller.UnFollow)
	// 	}

	// 	user := v1.Group("/users/:id")
	// 	{
	// 		user.GET("/informations", controller.GetUserInformation)
	// 		user.GET("/informations/brief", controller.GetBriefUserInformation)
	// 		user.GET("/followers", controller.GetFollowersInformation)
	// 		user.GET("/followings", controller.GetFollowingUsersInformation)
	// 	}

	// 	information := v1.Group("/informations")
	// 	{
	// 		//information.Use(middleware.VerifyJWT)

	// 		information.POST("/", userInformationHandler.InitUserInformation())
	// 		information.PATCH("/", controller.UpdateUserInformation)
	// 	}

	// 	post := v1.Group("/posts")
	// 	{
	// 		post.GET("/:id", controller.GetAPost)
	// 		post.GET("/:id/comments", controller.GetCommentsOfPost)
	// 		post.GET("/:id/interacts", controller.GetUserInteractOfPost)

	// 		post.Use(middleware.VerifyJWT)
	// 		post.GET("/", controller.GetPostForUser)
	// 		post.POST("/", controller.CreatePost)

	// 		post.Use(cors.New(cors_))

	// 		comment := post.Group("/:id/comments")
	// 		{
	// 			comment.POST("/", controller.CreateComment)
	// 			comment.DELETE("/:comment_id", controller.DeleteComment)
	// 			comment.PATCH("/:comment_id", controller.UpdateComment)
	// 		}

	// 		interact := post.Group("/:id/likes")
	// 		{
	// 			interact.GET("", controller.GetAnInteractOfPost)
	// 			interact.POST("", controller.CreateInteract)
	// 			interact.DELETE("", controller.DeleteInteract)
	// 		}
	// 		post.Use(middleware.VerifyPostOwn)
	// 		post.DELETE("/:id", controller.DeletePost)
	// 		post.PATCH("/:id", controller.UpdatePost)

	// 	}
	// }

	r.Run()

}
