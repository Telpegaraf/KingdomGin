package router

import (
	"github.com/gin-gonic/gin"
	"kingdom/api"
	"kingdom/config"
	"kingdom/database"
	gerror "kingdom/error"
)

func Create(db *database.GormDatabase, conf *config.Configuration) (*gin.Engine, func()) {
	g := gin.New()
	g.RemoteIPHeaders = []string{"X-Forwarded-For"}
	g.SetTrustedProxies(conf.Server.TrustedProxies)
	g.ForwardedByClientIP = true
	g.Use(func(ctx *gin.Context) {
		// Map sockets "@" to 127.0.0.1, because gin-gonic can only trust IPs.
		if ctx.Request.RemoteAddr == "@" {
			ctx.Request.RemoteAddr = "127.0.0.1:65535"
		}
	})

	userChangeNotifier := new(api.UserChangeNotifier)
	userHandler := api.UserApi{DB: db, PasswordStrength: conf.PassStrength, UserChangeNotifier: userChangeNotifier, Registration: conf.Registration}

	g.NoRoute(gerror.NotFound())

	userGroup := g.Group("/user")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("", userHandler.GetUsers)
	}
	return g, func() {}
}
