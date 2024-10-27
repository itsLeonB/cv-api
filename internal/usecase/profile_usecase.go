package usecase

import (
	"context"

	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/model/converter"
	"github.com/itsLeonB/cv-api/internal/repository"
)

type ProfileUsecase interface {
	GetShortSummaryByID(ctx context.Context, id int) (*model.Summary, error)
	GetSummaryByID(ctx context.Context, id int) (*model.Summary, error)
}

type profileUsecase struct {
	structName        string
	profileRepository repository.ProfileRepository
}

func NewUseCase(profileRepository repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{"profileUsecase", profileRepository}
}

func (u *profileUsecase) GetShortSummaryByID(ctx context.Context, id int) (*model.Summary, error) {
	profile, err := u.profileRepository.GetShortSummaryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.ProfileToShortSummary(profile), nil
}

func (u *profileUsecase) GetSummaryByID(ctx context.Context, id int) (*model.Summary, error) {
	profile, err := u.profileRepository.GetSummaryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.ProfileToSummary(profile), nil
}
