package helpers

import (
	"gin_todo/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromSession(c *gin.Context) *models.User {
	userID := c.GetUint64("user_id")
	if userID > 0 {
		return models.UserGetByID(userID)
	} else {
		return nil
	}
}

func SetPayload(c *gin.Context, h gin.H) gin.H {
	email := c.GetString("email")
	if len(email) > 0 {
		h["email"] = email
	}
	return h
}
