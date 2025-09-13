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
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"*"},
		MaxAge:          12 * time.Hour,
	}))

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
