package main

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	routes "VideoStreamingBackend/Routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*
	 type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
*/
var testing = false

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDB()

	if testing {
		err = config.DB.Migrator().DropTable(&models.RefreshToken{}, &models.Comment{}, &models.Video{}, &models.Channel{}, &models.User{})
		if err != nil {
			panic("Failed to drop tables!")
		}

		time.Sleep(100 * time.Millisecond)

		err = config.DB.AutoMigrate(&models.User{})
		if err != nil {
			panic("User migration error: " + err.Error())
		}

		err = config.DB.AutoMigrate(&models.RefreshToken{})
		if err != nil {
			panic("RefreshToken migration error: " + err.Error())
		}

		err = config.DB.AutoMigrate(&models.Channel{})
		if err != nil {
			panic("Channel migration error: " + err.Error())
		}

		err = config.DB.AutoMigrate(&models.Comment{})
		if err != nil {
			panic("Comment migration error: " + err.Error())
		}

		err = config.DB.AutoMigrate(&models.Video{})
		if err != nil {
			panic("Video migration error: " + err.Error())
		}
	}

	routes.SetupRoutes(r)

	r.Run(":" + os.Getenv("PORT"))
}
