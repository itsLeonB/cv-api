package apphttp

import (
	"github.com/itsLeonB/cv-api/internal/repository"
	"github.com/itsLeonB/cv-api/internal/usecase"
	"github.com/jackc/pgx/v5"
)

type Controllers struct {
	Controller *ProfileController
	Auth       *AuthController
	Skill      *SkillController
}

func SetupControllers(conn *pgx.Conn) *Controllers {
	repo := repository.NewProfileRepository(conn)
	userRepository := repository.NewUserRepository(conn)
	skillRepository := repository.NewSkillRepository(conn)

	useCase := usecase.NewUseCase(repo)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	skillUsecase := usecase.NewSkillUsecase(skillRepository)

	controller := NewProfileController(useCase)
	authController := NewAuthController(authUsecase)
	skillController := NewSkillController(skillUsecase)

	return &Controllers{
		Controller: controller,
		Auth:       authController,
		Skill:      skillController,
	}
}
