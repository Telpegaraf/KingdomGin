package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"kingdom/auth/password"
	"kingdom/model"
	"net/http"
	"os"
	"strings"
)

const (
	headerName = "Kingdom-Key"
)

type Database interface {
	GetUserByUsername(name string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByToken(token string) (*model.User, error)
}

type Auth struct {
	DB Database
}

type authenticate func(tokenID string, user *model.User) (authenticated, success bool, userId uint, err error)

//func (a *Auth) RequireAdmin() gin.HandlerFunc {
//	return a.requireToken(func(tokenID string, user *model.User) (bool, bool, uint, error) {
//		if user != nil {
//			return true, user.Admin, user.ID, nil
//		}
//		if token, err := a.DB.GetUserByToken(tokenID); err == nil {
//			return false, false, 0, err
//		} else if token != nil {
//			user, err := a.DB.GetUserByID(token.ID)
//			if err != nil {
//				return false, false, token.ID, err
//			}
//			return true, user.Admin, token.ID, nil
//		}
//		return false, false, 0, nil
//	})
//}

func (a *Auth) tokenFromQueryOrHeader(ctx *gin.Context) string {
	if token := a.tokenFromQuery(ctx); token != "" {
		return token
	} else if token := a.tokenFromKingdomHeader(ctx); token != "" {
		return token
	} else if token := a.tokenFromAuthorizationHeader(ctx); token != "" {
		return token
	}
	return ""
}

func (a *Auth) tokenFromQuery(ctx *gin.Context) string {
	return ctx.Request.URL.Query().Get("token")
}

func (a *Auth) tokenFromKingdomHeader(ctx *gin.Context) string {
	return ctx.Request.Header.Get(headerName)
}

func (a *Auth) tokenFromAuthorizationHeader(ctx *gin.Context) string {
	const prefix = "Bearer "
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	if len(authHeader) < len(prefix) || !strings.EqualFold(prefix, authHeader[:len(prefix)]) {
		return ""
	}

	return authHeader[len(prefix):]
}

func (a *Auth) userFromBasicAuth(ctx *gin.Context) (*model.User, error) {
	if name, pass, ok := ctx.Request.BasicAuth(); ok {
		if user, err := a.DB.GetUserByUsername(name); err != nil {
			return nil, err
		} else if user != nil && password.ComparePassword(user.Password, []byte(pass)) {
			return user, nil
		}
	}
	return nil, nil
}

func (a *Auth) requireToken(auth authenticate) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if strings.HasPrefix(ctx.Request.URL.Path, "/swagger/") {
			ctx.Next()
			return
		}

		token := a.tokenFromKingdomHeader(ctx)
		user, err := a.userFromBasicAuth(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(500, errors.New("an error occurred while authenticating user"))
			return
		}

		if user == nil || token != "" {
			authenticated, ok, userID, err := auth(token, user)
			if err != nil {
				ctx.AbortWithError(500, errors.New("an error occurred while authenticating user"))
				return
			} else if ok {
				RegisterAuthentication(ctx, user, userID, token)
				ctx.Next()
				return
			} else if authenticated {
				ctx.AbortWithStatusJSON(403, errors.New("you are not allowed to access this api"))
				return
			}
		}
		ctx.AbortWithError(401, errors.New("you need to provide a valid access token or user credentials to access this api"))
	}
}

func (a *Auth) Optional() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := a.tokenFromQueryOrHeader(ctx)
		user, err := a.userFromBasicAuth(ctx)
		if err != nil {
			RegisterAuthentication(ctx, nil, 0, "")
			ctx.Next()
			return
		}

		if user != nil {
			RegisterAuthentication(ctx, user, user.ID, token)
			ctx.Next()
			return
		} else if token != "" {
			if tokenUser, err := a.DB.GetUserByToken(token); err == nil && tokenUser != nil {
				RegisterAuthentication(ctx, user, tokenUser.ID, token)
				ctx.Next()
				return
			}
		}
		RegisterAuthentication(ctx, nil, 0, "")
		ctx.Next()
	}
}

func (a *Auth) RequireJWT(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat64, ok := claims["sub"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		userIdUint := uint(userIdFloat64)
		c.Set("userID", userIdUint)
		c.Set("isAdmin", claims["isAdmin"])
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	c.Next()
}

func (a *Auth) RequireAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat64, ok := claims["sub"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		userIdUint := uint(userIdFloat64)
		c.Set("userID", userIdUint)
		c.Set("isAdmin", claims["isAdmin"])
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	user, _ := a.DB.GetUserByID(GetUserID(c))
	if !user.Admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
		c.Abort()
		return
	}
	c.Next()
}
