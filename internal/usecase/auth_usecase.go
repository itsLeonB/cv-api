package usecase

import (
	"context"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/model/converter"
	"github.com/itsLeonB/cv-api/internal/repository"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *model.RegisterRequest) error
}

type authUsecase struct {
	structName     string
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecase{"authUsecase", userRepository}
}

func (u *authUsecase) Register(ctx context.Context, req *model.RegisterRequest) error {
	methodName := "Register()"
	existingUser, err := u.userRepository.SelectByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return apperror.NewAppError(
			httperror.ConflictError("email already registered"),
			u.structName, methodName, "existingUser != nil",
		)
	}

	insertingUser := converter.RegisterRequestToUser(req)
	err = u.userRepository.Insert(ctx, insertingUser)
	if err != nil {
		return err
	}

	return nil
}
