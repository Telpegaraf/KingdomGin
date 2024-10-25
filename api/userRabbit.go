package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kingdom/auth/password"
	"kingdom/model"
	"net/http"
	"net/mail"
)

type UserRabbitDatabase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	DeleteUserByID(id uint) error
	UpdateUser(user *model.User) error
	CreateUser(user *model.User) error
	CountUser(condition ...interface{}) (int, error)
}

type UserConsumer interface {
	Publish(username string, email string)
}

type UserRabbitApi struct {
	DB                 UserDatabase
	PasswordStrength   int
	UserChangeNotifier *UserChangeNotifier
	Registration       bool
	Consumer           UserConsumer
}

// CreateUserRabbit godoc
//
// @Summary Returns all users
// @Description Returns all users
// @Tags A
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "User data"
// @Success 200 {object} model.UserExternal "user details"
// @Router /a/rabbit [post]
func (a *UserRabbitApi) CreateUserRabbit(ctx *gin.Context) {
	user := model.CreateUser{}
	if err := ctx.ShouldBindJSON(&user); err == nil {
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		internal := &model.User{
			Username: user.Username,
			Email:    user.Email,
			Password: password.CreatePassword(user.Password, a.PasswordStrength),
		}
		existingUser, err := a.DB.GetUserByUsername(internal.Username)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if existingUser == nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.CreateUser(internal)); !success {
				return
			}
			ctx.JSON(201, toExternalUser(internal))

			a.Consumer.Publish(internal.Username, internal.Email)

			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}

	}
}
