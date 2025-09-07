package handler

import (
	"fmt"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/application"
)

func ConvertStringToUint(s string) (uint, error) {
	var id uint
	_, err := fmt.Sscan(s, &id)
	if err != nil {
		return 0, application.NewInvalidParamError("invalid ID format")
	}
	return id, nil
}
