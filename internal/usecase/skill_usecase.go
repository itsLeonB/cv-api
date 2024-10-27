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
	CreateCategory(ctx context.Context, request *model.NewSkillCategoryRequest) (*model.SkillCategoryResponse, error)
	GetAllCategories(ctx context.Context) ([]*model.SkillCategoryResponse, error)
	GetCategoryByID(ctx context.Context, id int) (*model.SkillCategoryResponse, error)
	Create(ctx context.Context, request *model.NewSkillRequest) (*model.SkillResponse, error)
	GetAll(ctx context.Context) ([]*model.SkillResponse, error)
	GetByID(ctx context.Context, id int) (*model.SkillResponse, error)
}

type skillUsecase struct {
	structName      string
	skillRepository repository.SkillRepository
}

func NewSkillUsecase(skillRepository repository.SkillRepository) SkillUsecase {
	return &skillUsecase{"skillUsecase", skillRepository}
}

func (u *skillUsecase) CreateCategory(ctx context.Context, request *model.NewSkillCategoryRequest) (*model.SkillCategoryResponse, error) {
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

func (u *skillUsecase) GetAllCategories(ctx context.Context) ([]*model.SkillCategoryResponse, error) {
	categories, err := u.skillRepository.SelectAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*model.SkillCategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = converter.SkillCategoryEntityToResponse(category)
	}

	return responses, nil
}

func (u *skillUsecase) GetCategoryByID(ctx context.Context, id int) (*model.SkillCategoryResponse, error) {
	methodName := fmt.Sprintf("GetCategoryByID(id: %d)", id)
	category, err := u.skillRepository.SelectCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, apperror.NewAppError(
			httperror.NotFoundError(fmt.Sprintf("skill category with id: %d is not found", id)),
			u.structName, methodName, "category == nil",
		)
	}

	return converter.SkillCategoryEntityToResponse(category), nil
}

func (u *skillUsecase) Create(ctx context.Context, request *model.NewSkillRequest) (*model.SkillResponse, error) {
	insertingSkill := converter.NewSkillRequestToEntity(request)
	err := u.skillRepository.Insert(ctx, insertingSkill)
	if err != nil {
		return nil, err
	}

	return converter.SkillEntityToResponse(insertingSkill), nil
}

func (u *skillUsecase) GetAll(ctx context.Context) ([]*model.SkillResponse, error) {
	skills, err := u.skillRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*model.SkillResponse, len(skills))
	for i := range skills {
		responses[i] = converter.SkillEntityToResponse(skills[i])
	}

	return responses, nil
}

func (u *skillUsecase) GetByID(ctx context.Context, id int) (*model.SkillResponse, error) {
	methodName := fmt.Sprintf("GetByID(id: %d)", id)
	skill, err := u.skillRepository.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if skill == nil {
		return nil, apperror.NewAppError(
			httperror.NotFoundError(fmt.Sprintf("skill with id: %d is not found", id)),
			u.structName, methodName, "skill == nil",
		)
	}

	return converter.SkillEntityToResponse(skill), nil
}
