package handler

import (
	"fmt"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService *application.ImageService
}

func NewImageHandler(imageService *application.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) RegisterRoutes(g *gin.RouterGroup) {
	// Register image-related routes here
	image := g.Group("/images")
	{
		image.POST("/", h.UploadImage)
		image.GET("/:id", h.GetImage)
		image.DELETE("/:id", h.DeleteImage)
		image.GET("/", h.ListImages)
	}
}

// @Summary Upload an image
// @Description Upload an image
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file to upload"
// @Success 200 {object} map[string]string "Returns success message"
// @Router /images [post]
func (h *ImageHandler) UploadImage(ctx *gin.Context) {
	// Handle image upload logic here

	file, err := ctx.FormFile("image")
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	image, err := h.imageService.UploadImage(file)
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Image uploaded successfully",
		"data":    image,
	})
}

// @Summary Get an image by ID
// @Description Get an image by ID
// @Tags images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Success 200 {object} map[string]interface{} "Returns image data"
// @Router /images/{id} [get]
func (h *ImageHandler) GetImage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		application.HandleError(ctx, application.NewInvalidParamError("id"))
		return
	}

	// Convert string id to uint
	var id uint
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		application.HandleError(ctx, application.NewInvalidParamError("id"))
		return
	}

	image, err := h.imageService.GetImageByID(id)
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": image,
	})
}

// @Summary Delete an image by ID
// @Description Delete an image by ID
// @Tags images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Success 204 "No Content"
// @Router /images/{id} [delete]
func (h *ImageHandler) DeleteImage(ctx *gin.Context) {
	id, err := ConvertStringToUint(ctx.Param("id"))
	if err != nil {
		application.HandleError(ctx, application.NewInvalidParamError("id"))
		return
	}

	if err := h.imageService.DeleteImage(id); err != nil {
		application.HandleError(ctx, err)
		return
	}

	ctx.JSON(204, nil)
}

// @Summary List all images
// @Description List all images
// @Tags images
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Returns list of images"
// @Router /images [get]
func (h *ImageHandler) ListImages(ctx *gin.Context) {
	images, err := h.imageService.ListImages()
	if err != nil {
		application.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": images,
	})
}
