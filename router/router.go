package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kingdom/api"
	"kingdom/auth"
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
		if ctx.Request.RemoteAddr == "@" {
			ctx.Request.RemoteAddr = "127.0.0.1:65535"
		}
	})
	authentication := auth.Auth{DB: db}

	userChangeNotifier := new(api.UserChangeNotifier)
	userHandler := api.UserApi{
		DB:                 db,
		PasswordStrength:   conf.PassStrength,
		UserChangeNotifier: userChangeNotifier,
		Registration:       conf.Registration}

	g.NoRoute(gerror.NotFound())

	g.Use(cors.New(auth.CorsConfig(conf)))

	userGroup := g.Group("/user").Use(authentication.Optional())
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUserByID)
	}
	return g, func() {}
}
