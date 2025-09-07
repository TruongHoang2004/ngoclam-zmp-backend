package repository

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/config"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/domain/entity"
	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/infrastructure/persistence/model"
)

type ImageRepositoryImpl struct {
	db       *gorm.DB
	basePath string // ví dụ: "./uploads"
}

func NewImageRepository(db *gorm.DB, config *config.Config) entity.ImageRepository {
	// đảm bảo thư mục tồn tại
	_ = os.MkdirAll(config.BasePath, 0755)
	return &ImageRepositoryImpl{db: db, basePath: config.BasePath}
}

// --- CRUD Image ---

// SaveFile lưu file upload vào thư mục uploads và insert record
func (r *ImageRepositoryImpl) SaveFile(file *multipart.FileHeader) (*entity.Image, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer src.Close()

	// đọc vài byte đầu để detect content-type
	buf := make([]byte, 512)
	n, _ := src.Read(buf)
	ct := http.DetectContentType(buf[:n])

	// chỉ cho phép các loại file ảnh
	allowed := map[string]bool{
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/webp":    true,
		"image/svg+xml": true,
	}
	if !allowed[ct] {
		return nil, fmt.Errorf("unsupported content type: %s", ct)
	}

	// reset về đầu
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("seek error: %w", err)
	}

	hash := sha256.New()
	if _, err := io.Copy(hash, src); err != nil {
		return nil, fmt.Errorf("cannot compute hash: %w", err)
	}
	fileHash := fmt.Sprintf("%x", hash.Sum(nil))
	log.Printf("File hash: %s", fileHash)

	// kiểm tra trùng lặp
	var existing model.Image
	if err := r.db.Where("hash = ?", fileHash).First(&existing).Error; err == nil {
		// đã tồn tại file giống hệt
		return existing.ToDomain(), nil
	} else if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("database error: %w", err)
	}

	// reset về đầu
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("seek error: %w", err)
	}

	// tạo tên file ngẫu nhiên
	random := make([]byte, 16)
	_, _ = rand.Read(random)
	exts, _ := mime.ExtensionsByType(ct)
	ext := ".bin"
	if len(exts) > 0 {
		ext = exts[0]
	}
	filename := hex.EncodeToString(random) + ext
	dstPath := filepath.Join(r.basePath, filename)

	// ghi file xuống uploads/
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, fmt.Errorf("cannot create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("cannot save file: %w", err)
	}

	// insert DB record
	image := &model.Image{
		Path: filename,
		Hash: fileHash,
	}
	if err := r.db.Create(image).Error; err != nil {
		// rollback: xoá file nếu DB fail
		_ = os.Remove(dstPath)
		return nil, err
	}

	return image.ToDomain(), nil
}

func (r *ImageRepositoryImpl) FindByID(id uint) (*entity.Image, error) {
	var img model.Image
	if err := r.db.First(&img, id).Error; err != nil {
		return nil, err
	}

	return img.ToDomain(), nil
}

func (r *ImageRepositoryImpl) FindAll() ([]*entity.Image, error) {
	var imgs []model.Image
	if err := r.db.Find(&imgs).Error; err != nil {
		return nil, err
	}
	result := make([]*entity.Image, len(imgs))
	for i, img := range imgs {
		result[i] = img.ToDomain()
	}
	return result, nil
}

func (r *ImageRepositoryImpl) Delete(id uint) error {
	var img model.Image
	if err := r.db.First(&img, id).Error; err != nil {
		return err
	}
	filePath := filepath.Join(r.basePath, img.Path)

	if err := r.db.Delete(&img).Error; err != nil {
		return err
	}
	// xoá file luôn
	_ = os.Remove(filePath)
	return nil
}
