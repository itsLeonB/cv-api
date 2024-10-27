package apphttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/usecase"
)

type AuthController struct {
	structName  string
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{"AuthController", authUsecase}
}

func (c *AuthController) HandleRegister() gin.HandlerFunc {
	methodName := "HandleRegister()"
	return func(ctx *gin.Context) {
		request := new(model.RegisterRequest)
		err := ctx.ShouldBindJSON(request)
		if err != nil {
			_ = ctx.Error(apperror.NewAppError(
				err, c.structName, methodName,
				"ctx.ShouldBindJSON()",
			))
			return
		}

		err = c.authUsecase.Register(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
	}
}
