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
	characterClassHandler := api.CharacterClassApi{
		DB: db,
	}
	itemHandler := api.ItemApi{
		DB: db,
	}

	godHandler := api.GodApi{DB: db}
	domainHandler := api.DomainApi{DB: db}

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
		userGroup.PATCH("/:id", userHandler.UpdateUser)
		userGroup.PATCH("/password", userHandler.ChangePassword)
	}
	characterGroup := g.Group("/character").Use(authentication.RequireJWT)
	{
		characterGroup.POST("", characterHandler.CreateCharacter)
		characterGroup.GET("/:id", characterHandler.GetCharacterByID)
		characterGroup.GET("", characterHandler.GetCharacters)
		characterGroup.PATCH("/:id", characterHandler.UpdateCharacter)
		characterGroup.DELETE("/:id", characterHandler.DeleteCharacter)
	}
	godGroup := g.Group("/god").Use(authentication.RequireAdmin)
	{
		godGroup.POST("", godHandler.CreateGod)
		godGroup.PATCH("/:id", godHandler.UpdateGod)
		godGroup.DELETE("/:id", godHandler.DeleteGod)
	}
	godGroup.GET("/:id", godHandler.GetGodById).Use(authentication.RequireJWT)
	godGroup.GET("", godHandler.GetGods).Use(authentication.RequireJWT)

	domainGroup := g.Group("/domain").Use(authentication.RequireAdmin)
	{
		domainGroup.POST("", domainHandler.CreateDomain)
		domainGroup.PATCH("/:id", domainHandler.UpdateDomain)
		domainGroup.DELETE("/:id", domainHandler.DeleteDomain)
	}
	g.GET("/domain/:id", domainHandler.GetDomainByID).Use(authentication.RequireJWT)
	g.GET("/domain", domainHandler.GetDomains).Use(authentication.RequireJWT)

	characterClassGroup := g.Group("/class").Use(authentication.RequireJWT)
	{
		characterClassGroup.POST("", characterClassHandler.CreateCharacterClass)
	}
	itemGroup := g.Group("/item").Use(authentication.RequireJWT)
	{
		itemGroup.GET("", itemHandler.GetItems)
		itemGroup.GET(":id", itemHandler.GetItemByID)
		itemGroup.GET("/armor", itemHandler.GetArmors)
		itemGroup.GET("/armor/:id", itemHandler.GetArmorByID)
		itemGroup.GET("/weapon", itemHandler.GetWeapons)
		itemGroup.GET("/weapon/:id", itemHandler.GetWeaponByID)
		itemGroup.GET("/gear", itemHandler.GetGears)
		itemGroup.GET("/gear/:id", itemHandler.GetGearByID)
	}
	itemGroup.DELETE("/:id", itemHandler.DeleteItem).Use(authentication.RequireAdmin)
	itemGroup.POST("/armor", itemHandler.CreateArmor).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/armor/:id", itemHandler.UpdateArmor).Use(authentication.RequireAdmin)
	itemGroup.POST("/weapon", itemHandler.CreateWeapon).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/weapon/:id", itemHandler.UpdateWeapon).Use(authentication.RequireAdmin)
	itemGroup.POST("/gear", itemHandler.CreateGear).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/gear/:id", itemHandler.UpdateGear).Use(authentication.RequireAdmin)

	return g, func() {}
}
