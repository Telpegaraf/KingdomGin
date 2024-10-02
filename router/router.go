package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kingdom/api"
	"kingdom/auth"
	"kingdom/config"
	"kingdom/database"
	"kingdom/docs"
	gerror "kingdom/error"
)

func Create(db *database.GormDatabase, conf *config.Configuration) (*gin.Engine, func()) {
	g := gin.New()
	g.RemoteIPHeaders = []string{"X-Forwarded-For"}
	g.SetTrustedProxies(conf.Server.TrustedProxies)
	g.ForwardedByClientIP = true
	g.Use(func(ctx *gin.Context) {
		if ctx.Request.RemoteAddr == "@" {
			ctx.Request.RemoteAddr = "localhost:8080"
		}
	})
	authentication := auth.Auth{DB: db}

	userChangeNotifier := new(api.UserChangeNotifier)
	userHandler := api.UserApi{
		DB:                 db,
		PasswordStrength:   conf.PassStrength,
		UserChangeNotifier: userChangeNotifier,
		Registration:       conf.Registration}

	characterHandler := api.CharacterApi{
		DB: db,
	}

	godHandler := api.GodApi{DB: db}

	authHandler := api.Controller{DB: db}

	g.NoRoute(gerror.NotFound())

	g.Use(cors.New(auth.CorsConfig(conf)))

	docs.SwaggerInfo.BasePath = ""
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST("/user", userHandler.CreateUser)
	g.POST("/login", authHandler.Login)
	g.GET("/validate", authHandler.Validate)

	userGroup := g.Group("/user").Use(authentication.RequireJWT)
	{
		userGroup.GET("", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUserByID)
	}
	characterGroup := g.Group("/character").Use(authentication.RequireJWT)
	{
		characterGroup.POST("", characterHandler.CreateCharacter)
		characterGroup.GET("/:id", characterHandler.GetCharacterByID)
		characterGroup.GET("", characterHandler.GetCharacters)
		characterGroup.PATCH("/:id", characterHandler.UpdateCharacter)
		characterGroup.DELETE("/:id", characterHandler.DeleteCharacter)
	}
	godGroup := g.Group("/god").Use(authentication.RequireJWT)
	{
		godGroup.GET("/:id", godHandler.GetGodById)
		godGroup.POST("", godHandler.CreateGod)
	}
	return g, func() {}
}
