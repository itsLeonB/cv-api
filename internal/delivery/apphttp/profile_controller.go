package apphttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/usecase"
)

type ProfileController struct {
	structName     string
	profileUsecase usecase.ProfileUsecase
}

func NewProfileController(profileUsecase usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{"profileUsecase", profileUsecase}
}

func (c *ProfileController) GetShortSummary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _, err := GetNumericQueryParam(ctx, "id", true)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		summary, err := c.profileUsecase.GetShortSummaryByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(summary))
	}
}

func (c *ProfileController) HandleSummary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _, err := GetNumericQueryParam(ctx, "id", true)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		summary, err := c.profileUsecase.GetSummaryByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(summary))
	}
}
