package main

import (
	"fmt"
	"mfc-creations/post-app/config"
	"mfc-creations/post-app/database"
	"mfc-creations/post-app/database/models"
)

func main() {
	config.LoadEnv()
	db := database.Connect()
	fmt.Println(db)
	database.DB.AutoMigrate(&models.Post{})
}
