package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"kingdom/model"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthDatabase interface {
	GetUserByEmail(email string) (*model.User, error)
}

type Controller struct {
	DB AuthDatabase
}

func (a *Controller) LoginPage(ctx *gin.Context) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error loading template")
		return
	}
	err = tmpl.Execute(ctx.Writer, nil)
	if err != nil {
		return
	}
}

// Login godoc
//
// @Summary Login user for token
// @Description Authorization by email and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body model.UserLogin true "User data"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func (a *Controller) Login(ctx *gin.Context) {
	body := model.UserLogin{}
	if err := ctx.Bind(&body); err != nil {
		log.Println("Failed to bind request body:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user, err := a.DB.GetUserByEmail(body.Email)
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
	ctx.Redirect(http.StatusFound, "/character")
}

func (a *Controller) Validate(c *gin.Context) {
	user, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
