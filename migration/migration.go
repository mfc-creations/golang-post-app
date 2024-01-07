package main

import (
	"mfc-creations/post-app/database"
	"mfc-creations/post-app/database/models"
)

func main(){
	db:=database.Connect()
	db.AutoMigrate(&models.Post{})
}