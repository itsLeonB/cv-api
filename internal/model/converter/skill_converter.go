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

func NewSkillRequestToEntity(request *model.NewSkillRequest) *entity.Skill {
	return &entity.Skill{
		ProfileID:   request.ProfileID,
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Description: request.Description,
	}
}

func SkillEntityToResponse(entity *entity.Skill) *model.SkillResponse {
	response := &model.SkillResponse{
		ID:          entity.ID,
		ProfileID:   entity.ProfileID,
		CategoryID:  entity.CategoryID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.String(),
		UpdatedAt:   entity.UpdatedAt.String(),
	}

	if entity.DeletedAt.Valid {
		response.DeletedAt = entity.DeletedAt.Time.String()
	}

	if entity.Category != nil {
		response.Category = SkillCategoryEntityToResponse(entity.Category)
	}

	return response
}
