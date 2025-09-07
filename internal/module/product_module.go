package module

import (
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/infrastructure/persistence/repository"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/interface/http/handler"
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var ProductModule = fx.Options(
	// Provide repository
	fx.Provide(repository.NewProductRepository),

	// Provide service
	fx.Provide(application.NewProductService),

	// Provide handler
	fx.Provide(handler.NewProductHandler),

	fx.Invoke(func(h *handler.ProductHandler, g *gin.RouterGroup) {
		h.RegisterRoutes(g)
	}),
)
