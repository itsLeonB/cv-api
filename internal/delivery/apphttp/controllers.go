package apphttp

import (
	"github.com/itsLeonB/cv-api/internal/repository"
	"github.com/itsLeonB/cv-api/internal/usecase"
	"github.com/jackc/pgx/v5"
)

type Controllers struct {
	Controller *ProfileController
	Auth       *AuthController
}

func SetupControllers(conn *pgx.Conn) *Controllers {
	repo := repository.NewProfileRepository(conn)
	userRepository := repository.NewUserRepository(conn)

	useCase := usecase.NewUseCase(repo)
	authUsecase := usecase.NewAuthUsecase(userRepository)

	controller := NewProfileController(useCase)
	authController := NewAuthController(authUsecase)

	return &Controllers{
		Controller: controller,
		Auth:       authController,
	}
}
