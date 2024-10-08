package usecase

import (
	"context"

	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/model/converter"
	"github.com/itsLeonB/cv-api/internal/repository"
)

type UseCase interface {
	GetShortSummaryByID(ctx context.Context, id int) (*model.ShortSummary, error)
}

type useCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{repo}
}

func (uc *useCase) GetShortSummaryByID(ctx context.Context, id int) (*model.ShortSummary, error) {
	profile, err := uc.repo.GetShortSummaryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.ProfileToShortSummary(profile), nil
}
