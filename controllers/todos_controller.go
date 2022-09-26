package controllers

import (
	"gin_todo/controllers/helpers"
	"gin_todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodosIndex(c *gin.Context) {
	// Get user from the request
	currentUser := helpers.GetUserFromSession(c)
	if currentUser == nil {
		c.HTML(
			http.StatusUnauthorized,
			"/todos/index.html",
			helpers.SetPayload(c, gin.H{
				"Error": "Unauthorized Access!",
			}),
		)
		return
	}
	todos := models.TodosGetAll(currentUser)
	c.HTML(
		http.StatusOK,
		"/todos/index.html",
		helpers.SetPayload(c, gin.H{
			"todos": todos,
		}),
	)
}

func TodosNew(c *gin.Context) {
	currentUser := helpers.GetUserFromSession(c)
	if currentUser == nil {
		c.HTML(
			http.StatusUnauthorized,
			"/todos/index.html",
			helpers.SetPayload(c, gin.H{
				"Error": "Unauthorized Access!",
			}),
		)
		return
	}
	c.HTML(
		http.StatusOK,
		"todos/new.html",
		helpers.SetPayload(c, gin.H{}),
	)
}

type FormData struct {
	Title  string `form:"title"`
	Detail string `form:"detail"`
}

func TodosCreate(c *gin.Context) {
	currentUser := helpers.GetUserFromSession(c)
	if currentUser == nil {
		c.HTML(
			http.StatusUnauthorized,
			"/todos/index.html",
			helpers.SetPayload(c, gin.H{
				"Error": "Unauthorized Access!",
			}),
		)
		return
	}

	var data FormData
	c.Bind(&data)
	models.TodoCreate(currentUser, data.Title, data.Detail)
	c.Redirect(http.StatusMovedPermanently, "/todos")
}

type JsonData struct {
	ID    uint64 `json:"id"`
	State bool   `json:"state"`
}

func TodosComplete(c *gin.Context) {
	currentUser := helpers.GetUserFromSession(c)
	if currentUser == nil {
		c.HTML(
			http.StatusUnauthorized,
			"/todos/index.html",
			helpers.SetPayload(c, gin.H{
				"Error": "Unauthorized Access!",
			}),
		)
		return
	}
	var data JsonData
	c.Bind(&data)
	todo := models.TodoFind(currentUser, data.ID)

	todo.MarkComplete(data.State)

	c.HTML(
		http.StatusOK,
		"/todos/index.html",
		helpers.SetPayload(c, gin.H{}),
	)

}
