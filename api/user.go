package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"kingdom/auth"
	"kingdom/auth/password"
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
}

type UserChangeNotifier struct {
	userDeletedCallbacks []func(uid uint) error
	userAddedCallbacks   []func(uid uint) error
}

func (c *UserChangeNotifier) OnUserDeleted(callback func(uid uint) error) {
	c.userDeletedCallbacks = append(c.userDeletedCallbacks, callback)
}

func (c *UserChangeNotifier) OnUserAdded(callback func(uid uint) error) {
	c.userAddedCallbacks = append(c.userAddedCallbacks, callback)
}

func (c *UserChangeNotifier) fireUserDeleted(uid uint) error {
	for _, callback := range c.userDeletedCallbacks {
		if err := callback(uid); err != nil {
			return err
		}
	}
	return nil
}

func (c *UserChangeNotifier) fireUserAdded(uid uint) error {
	for _, callback := range c.userAddedCallbacks {
		if err := callback(uid); err != nil {
			return err
		}
	}
	return nil
}

type UserApi struct {
	DB                 UserDatabase
	PasswordStrength   int
	UserChangeNotifier *UserChangeNotifier
	Registration       bool
}

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

func (a *UserApi) CreateUser(ctx *gin.Context) {
	user := model.CreateUser{}
	if err := ctx.Bind(&user); err != nil {
		internal := &model.User{
			Username: user.Username,
			Admin:    user.Admin,
			Password: password.CreatePassword(user.Password, a.PasswordStrength),
		}
		existingUser, err := a.DB.GetUserByUsername(internal.Username)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}

		var requestedBy *model.User
		uid := auth.TryGetUserID(ctx)
		if uid != nil {
			requestedBy, err = a.DB.GetUserByID(*uid)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not get user: %s", err))
				return
			}
		}

		if requestedBy == nil || !requestedBy.Admin {
			status := http.StatusUnauthorized
			if requestedBy != nil {
				status = http.StatusForbidden
			}
			if !a.Registration {
				ctx.AbortWithError(status, errors.New("you are not allowed to access this api"))
				return
			}
			if internal.Admin {
				ctx.AbortWithError(http.StatusUnauthorized, errors.New("you are not allowed to create an admin user"))
				return
			}
		}

		if existingUser == nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.CreateUser(internal)); !success {
				return
			}
			if err := a.UserChangeNotifier.fireUserAdded(internal.ID); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(200, toExternalUser(internal))
		} else {
			ctx.AbortWithError(400, errors.New("user already exists"))
		}
	}
}

func toExternalUser(internal *model.User) *model.UserExternal {
	return &model.UserExternal{
		Username: internal.Username,
		Admin:    internal.Admin,
		ID:       internal.ID,
	}
}
