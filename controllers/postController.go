package controllers

import (
	"mfc-creations/post-app/database"
	"mfc-creations/post-app/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	var body struct {
		Title       string
		Description string
	}
	err := context.BindJSON(&body)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Enter valid inputs"})
		return
	}
	post := models.Post{Title: body.Title, Description: body.Description}
	result := database.DB.Create(&post)

	if result.Error != nil {
		context.Status(400)
		return
	}
	context.IndentedJSON(http.StatusCreated, post)
}
