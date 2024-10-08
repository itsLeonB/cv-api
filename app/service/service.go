package service

import (
	"context"
	"fmt"

	"github.com/itsLeonB/cv-api/dto"
	"github.com/itsLeonB/cv-api/repository"
)

type Service interface {
	GetShortSummary(context.Context) (*dto.ShortSummary, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(r repository.Repository) *serviceImpl {
	return &serviceImpl{r}
}

func (s *serviceImpl) GetShortSummary(ctx context.Context) (*dto.ShortSummary, error) {
	profile, err := s.repo.GetShortSummary(ctx)
	if err != nil {
		return nil, err
	}

	summary := dto.ShortSummary{
		Header: fmt.Sprintf("Hi, I'm %s 👋", profile.Nickname),
		Body:   fmt.Sprintf("I'm a %s based in %s. %s", profile.Occupation, profile.Location, profile.ShortSummary),
	}

	return &summary, nil
}
