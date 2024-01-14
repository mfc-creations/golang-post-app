package main

import (
	"errors"
	"fmt"
	"mfc-creations/post-app/database"
	"net/http"

	"mfc-creations/post-app/config"
	"mfc-creations/post-app/controllers"

	"github.com/gin-gonic/gin"
)

type post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var posts = []post{
	{ID: "1", Title: "Test title", Description: "Test description"},
}

func getPosts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, posts)
}

func createPost(context *gin.Context) {
	var newPost post
	err := context.BindJSON(&newPost)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Enter valid inputs"})
		return
	}
	posts = append(posts, newPost)
	context.IndentedJSON(http.StatusCreated, newPost)

}

func postById(id string) (*post, error) {
	for index, val := range posts {
		if val.ID == id {
			return &posts[index], nil
		}
	}
	return nil, errors.New("Post not found")
}

func updatePost(context *gin.Context) {
	id := context.Param("id")
	currentPost, err := postById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	var newPost post
	error := context.BindJSON(&newPost)
	if error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Enter valid inputs"})
		return
	}

	currentPost.Title = newPost.Title
	currentPost.Description = newPost.Description

	context.IndentedJSON(http.StatusOK, currentPost)
}

func deletePost(context *gin.Context) {
	id := context.Param("id")
	for index, val := range posts {
		if val.ID == id {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func getPost(context *gin.Context) {
	id := context.Param("id")
	currentPost, err := postById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, currentPost)
}

func init() {
	config.LoadEnv()
	db := database.Connect()
	fmt.Println(db)
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.POST("/post", controllers.CreatePost)
	router.PATCH("/post/:id", updatePost)
	router.DELETE("/post/:id", deletePost)
	router.GET("/post/:id", getPost)

	router.Run("localhost:5000")
}
