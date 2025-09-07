package module

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(r *gin.Engine) {

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
