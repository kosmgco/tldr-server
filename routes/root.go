package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kosmgco/tldr/global"
	"time"
)

func setUpRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Sc-Sso-Token", "X-Request-Id", "Dmp-Management-Token"},
		MaxAge:           time.Hour * 24,
		AllowCredentials: true,
	}))

	rootRouter := router.Group("/tldr")
	{
		rootRouter.GET("/search", Search)
		rootRouter.GET("/get", Get)
		rootRouter.GET("/conf", GetConf)
		rootRouter.GET("/hot", Hot)
	}
}

func Start() {
	global.Config.GinApp.Register(setUpRoutes)
	global.Config.GinApp.Start()
}
