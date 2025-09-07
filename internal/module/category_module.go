package module

import (
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/infrastructure/persistence/repository"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/interface/http/handler"
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var CategoryModule = fx.Options(
	// Provide repository
	fx.Provide(repository.NewCategoryRepository),

	// Provide service
	fx.Provide(application.NewCategoryService),

	// Provide handler
	fx.Provide(handler.NewCategoryHandler),

	fx.Invoke(func(h *handler.CategoryHandler, g *gin.RouterGroup) {
		h.RegisterRoutes(g)
	}),
)
