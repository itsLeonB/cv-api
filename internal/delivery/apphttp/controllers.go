package apphttp

import (
	"github.com/itsLeonB/cv-api/internal/repository"
	"github.com/itsLeonB/cv-api/internal/usecase"
	"github.com/jackc/pgx/v5"
)

type Controllers struct {
	Controller *Controller
}

func SetupControllers(conn *pgx.Conn) *Controllers {
	repo := repository.NewRepository(conn)

	useCase := usecase.NewUseCase(repo)

	controller := NewController(useCase)

	return &Controllers{
		Controller: controller,
	}
}
