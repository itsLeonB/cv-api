package apphttp

import (
	"github.com/itsLeonB/cv-api/internal/repository"
	"github.com/itsLeonB/cv-api/internal/usecase"
	"github.com/jackc/pgx/v5"
)

type Controllers struct {
	Controller *Controller
	Auth       *AuthController
}

func SetupControllers(conn *pgx.Conn) *Controllers {
	repo := repository.NewRepository(conn)
	userRepository := repository.NewUserRepository(conn)

	useCase := usecase.NewUseCase(repo)
	authUsecase := usecase.NewAuthUsecase(userRepository)

	controller := NewController(useCase)
	authController := NewAuthController(authUsecase)

	return &Controllers{
		Controller: controller,
		Auth:       authController,
	}
}
