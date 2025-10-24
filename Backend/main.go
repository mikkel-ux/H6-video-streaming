package main

import (
	config "VideoStreamingBackend/Config"
	models "VideoStreamingBackend/Models"
	"time"

	"github.com/joho/godotenv"
)

/* type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
} */

func main() {
	/* r := gin.Default() */
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDB()

	err = config.DB.Migrator().DropTable(&models.Video{}, &models.Channel{}, &models.User{})
	if err != nil {
		panic("Failed to drop tables!")
	}

	/* err = config.DB.AutoMigrate(&models.User{}, &models.Channel{}, &models.Video{})
	if err != nil {
		panic("Failed to migrate tables!")
	} */

	time.Sleep(100 * time.Millisecond)

	err = config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("User migration error: " + err.Error())
	}

	err = config.DB.AutoMigrate(&models.Channel{})
	if err != nil {
		panic("Channel migration error: " + err.Error())
	}

	err = config.DB.AutoMigrate(&models.Video{})
	if err != nil {
		panic("Video migration error: " + err.Error())
	}

	/* r.Run(":" + os.Getenv("PORT")) */
}
