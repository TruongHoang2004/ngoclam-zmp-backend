package module

import (
	"github.com/TruongHoang2004/ngoclam-zmp-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/fx"
)

var SwaggerModule = fx.Options(
	fx.Invoke(registerSwaggerRoutes),
)

func registerSwaggerRoutes(r *gin.Engine) {
	// Swagger UI sẽ có ở endpoint: /swagger/index.html
	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
