package converter

import (
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/itsLeonB/cv-api/internal/model"
)

func RegisterRequestToUser(request *model.RegisterRequest) *entity.User {
	return &entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
}
