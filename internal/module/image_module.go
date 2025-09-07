package module

import (
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/infrastructure/persistence/repository"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/interface/http/handler"
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var ImageModule = fx.Options(
	// Provide repository
	fx.Provide(repository.NewImageRepository),

	// Provide service
	fx.Provide(application.NewImageService),

	// Provide handler
	fx.Provide(handler.NewImageHandler),

	fx.Invoke(func(h *handler.ImageHandler, g *gin.RouterGroup) {
		h.RegisterRoutes(g)
	}),
)
