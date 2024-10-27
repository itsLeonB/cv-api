package usecase

import (
	"context"
	"fmt"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/model/converter"
	"github.com/itsLeonB/cv-api/internal/repository"
)

type SkillUsecase interface {
	InsertCategory(ctx context.Context, request *model.NewSkillCategoryRequest) (*model.SkillCategoryResponse, error)
}

type skillUsecase struct {
	structName      string
	skillRepository repository.SkillRepository
}

func NewSkillUsecase(skillRepository repository.SkillRepository) SkillUsecase {
	return &skillUsecase{"skillUsecase", skillRepository}
}

func (u *skillUsecase) InsertCategory(ctx context.Context, request *model.NewSkillCategoryRequest) (*model.SkillCategoryResponse, error) {
	methodName := "InsertCategory()"
	existingCategory, err := u.skillRepository.SelectCategoryByName(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	if existingCategory != nil {
		return nil, apperror.NewAppError(
			httperror.ConflictError(fmt.Sprintf("category %s already exists", request.Name)),
			u.structName, methodName, "existingCategory != nil",
		)
	}

	insertingCategory := converter.NewSkillCategoryRequestToEntity(request)
	err = u.skillRepository.InsertCategory(ctx, insertingCategory)
	if err != nil {
		return nil, err
	}

	return converter.SkillCategoryEntityToResponse(insertingCategory), nil
}
