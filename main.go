package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rizqirenaldy/api-go-gin-gorm/controllers"
	"github.com/rizqirenaldy/api-go-gin-gorm/models"
)

type error interface {
	Error() string
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Load .env
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	// setup DB
	models.DatabaseConfig()

	// route
	r.GET("/", controllers.Home)
	r.POST("/friend/request", controllers.RequestFriend)
	r.POST("/friend/:status", controllers.RequestFriendStatus)
	r.POST("/friend/request-list", controllers.RequestFriendData)
	r.POST("/friend-list", controllers.FriendList)
	r.POST("/friend-list/common", controllers.FriendListCommon)
	r.POST("/friend/block", controllers.BlockFriend)

	// running app
	appPort := ":" + os.Getenv("PORT")
	log.Println("Server Running in port" + appPort)
	r.Run(appPort)
}
