package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type AuthDatabase interface {
	GetUserByEmail(email string) (*model.User, error)
}

type Controller struct {
	DB AuthDatabase
}

func (a *Controller) Validate(c *gin.Context) {
	user, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
