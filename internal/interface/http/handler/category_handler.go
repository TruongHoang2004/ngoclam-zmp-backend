package handler

import (
	"log"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/interface/http/dto"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *application.CategoryService
}

func NewCategoryHandler(categoryService *application.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.RouterGroup) {
	category := r.Group("/categories")
	{
		category.GET("/", h.FindAllCategories)
		category.POST("/", h.CreateCategory)
		category.GET("/:id", h.FindCategoryByID)
	}
}

// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Category to create"
// @Success 200 {object} dto.CategoryResponse "Returns the created category"
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var createCategoryRequest dto.CreateCategoryRequest

	// Validate the request
	if err := ctx.ShouldBindJSON(&createCategoryRequest); err != nil {
		application.HandleError(ctx, err)
		return
	}

	category, err := h.categoryService.CreateCategory(ctx, createCategoryRequest.ToDomain())
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	log.Println("Created category:", category)

	ctx.JSON(201, gin.H{
		"message": "Category created successfully",
		"data":    dto.NewCategoryResponse(*category),
	})
}

// @Summary Find category by ID
// @Description Get a category by its ID
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryResponse "Returns the category"
// @Router /categories/{id} [get]
func (h *CategoryHandler) FindCategoryByID(ctx *gin.Context) {
	id, err := ConvertStringToUint(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.categoryService.GetCategoryByID(ctx, id)
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": dto.NewCategoryResponse(*category),
	})

}

// @Summary Find all categories
// @Description Get a list of all categories
// @Tags categories
// @Success 200 {object} map[string]interface{} "Returns a list of categories"
// @Router /categories [get]
func (h *CategoryHandler) FindAllCategories(ctx *gin.Context) {
	categories, err := h.categoryService.GetAllCategories(ctx)
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	respone := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		respone = append(respone, dto.NewCategoryResponse(*category))
	}

	ctx.JSON(200, gin.H{
		"data": respone,
	})
}
