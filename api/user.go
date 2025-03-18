package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kingdom/auth"
	"kingdom/model"
	"net/http"
)

type UserDatabase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	DeleteUserByID(id uint) error
	UpdateUser(user *model.User) error
	CreateUser(user *model.User) error
	CountUser(condition ...interface{}) (int, error)
	GetUserCodeByEmail(email string) (*model.UserCode, error)
	UpdateUserVerification(user *model.User) error
}

type UserConsumer interface {
	Publish(email string)
}

type UserApi struct {
	DB               UserDatabase
	PasswordStrength int
	Registration     bool
	Consumer         UserConsumer
}

// GetCurrentUser godoc
//
// @Summary Returns current user
// @Description Returns current user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.UserExternal "User current"
// @Failure 500
// @Router /user/current [get]
func (a *UserApi) GetCurrentUser(ctx *gin.Context) {
	user, err := a.DB.GetUserByID(auth.GetUserID(ctx))
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	ctx.JSON(200, toExternalUser(user))
}

// GetUsers godoc
//
// @Summary Returns all users
// @Description Returns all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.UserExternal "User list"
// @Failure 500
// @Router /user [get]
func (a *UserApi) GetUsers(ctx *gin.Context) {
	users, err := a.DB.GetUsers()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.UserExternal
	for _, user := range users {
		resp = append(resp, toExternalUser(user))
	}
	ctx.JSON(200, resp)
}

// GetUserByID godoc
//
// @Summary returns User by ID
// @Description Retrieve User details using its ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} model.UserExternal "user details"
// @Failure 404 {string} string "User not found"
// @Router /user/{id} [get]
func (a *UserApi) GetUserByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		user, err := a.DB.GetUserByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(404, gin.H{"error": "User not found"})
		}
		if user != nil {
			ctx.JSON(200, toExternalUser(user))
		} else {
			ctx.JSON(404, gin.H{"error": "User not found"})
		}
	})
}

func (a *UserApi) GetUserByUsername(ctx *gin.Context) {
	user, err := a.DB.GetUserByUsername("username")
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	if user != nil {
		ctx.JSON(200, toExternalUser(user))
	} else {
		ctx.JSON(404, errors.New("User not found"))
	}
}

// CreateUser godoc
//
// @Summary Create and returns user or nil
// @Description Create new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "User data"
// @Success 200 {object} model.UserExternal "user details"
// @Failure 404 {string} string "User already exist"
// @Router /user [post]
func (a *UserApi) CreateUser(ctx *gin.Context) {
	user := model.CreateUser{}
	if err := ctx.ShouldBindJSON(&user); err == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
	}
}

// DeleteUserByID godoc
//
// @Summary Returns and delete User by ID if you're admin
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 201 {string} string "Ok"
// @Failure 400 {string} string "User doesn't exist"
// @Failure 401 {string} string "You need to provide a valid access token or user credentials to access this api"
// @Router /user/{id} [delete]
func (a *UserApi) DeleteUserByID(ctx *gin.Context) {
	currentUserID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser, _ := a.DB.GetUserByID(currentUserID.(uint))
	if !currentUser.Admin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
		return
	}

	withID(ctx, "id", func(id uint) {
		user, err := a.DB.GetUserByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if user != nil {
			adminCount, err := a.DB.CountUser(&model.User{Admin: true})
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if user.Admin && adminCount == 1 {
				ctx.AbortWithError(400, errors.New("can't delete last admin"))
				return
			}
			SuccessOrAbort(ctx, 500, a.DB.DeleteUserByID(id))
		} else {
			ctx.AbortWithError(400, errors.New("user doesn't exist"))
		}
	})
}

func toExternalUser(internal *model.User) *model.UserExternal {
	return &model.UserExternal{
		Username:   internal.Username,
		Admin:      internal.Admin,
		ID:         internal.ID,
		Characters: internal.Characters,
	}
}
