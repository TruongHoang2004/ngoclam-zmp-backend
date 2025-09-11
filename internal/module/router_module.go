package module

import (
	"time"

	. "github.com/TruongHoang2004/ngoclam-zmp-backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(r *gin.Engine, config *Config) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", config.ZaloMiniAppHost},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/uploads", "./uploads")

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}

var RouterModule = fx.Options(
	fx.Provide(func(g *gin.Engine) *gin.RouterGroup {
		return g.Group("/api")
	}),

	fx.Invoke(RegisterRoutes),
)
