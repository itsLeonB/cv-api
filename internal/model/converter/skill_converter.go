package converter

import (
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/itsLeonB/cv-api/internal/model"
)

func NewSkillCategoryRequestToEntity(request *model.NewSkillCategoryRequest) *entity.SkillCategory {
	return &entity.SkillCategory{Name: request.Name}
}

func SkillCategoryEntityToResponse(entity *entity.SkillCategory) *model.SkillCategoryResponse {
	response := &model.SkillCategoryResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
	}

	if entity.DeletedAt.Valid {
		response.DeletedAt = entity.DeletedAt.Time.String()
	}

	return response
}
