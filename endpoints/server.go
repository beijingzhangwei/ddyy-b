package endpoints

import (
	"fmt"
	"github.com/beijingzhangwei/ddyy-b/endpoints/controllers"
	"github.com/beijingzhangwei/ddyy-b/endpoints/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load() // 环境变量
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}) //database migration
	//seed.Load(server.DB)
	server.Run(":3000")
}
