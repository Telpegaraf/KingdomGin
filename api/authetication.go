package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"kingdom/model"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthDatabase interface {
	GetUserByUsername(username string) (*model.User, error)
}

type Controller struct {
	DB AuthDatabase
}

// Login godoc
//
// @Summary Login user for token
// @Description Авторизация пользователя по логину и паролю
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body model.UserLogin true "User data"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func (a *Controller) Login(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.Bind(&body); err != nil {
		log.Println("Failed to bind request body:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user, err := a.DB.GetUserByUsername(body.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"isAdmin": user.Admin,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})
}

func (a *Controller) Validate(c *gin.Context) {
	user, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
