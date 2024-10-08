package apphttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/usecase"
)

type Controller struct {
	uc usecase.UseCase
}

func NewController(uc usecase.UseCase) *Controller {
	return &Controller{uc}
}

func (c *Controller) GetShortSummary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := QueryNumeric(ctx, "id")
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		summary, err := c.uc.GetShortSummaryByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(summary))
	}
}
