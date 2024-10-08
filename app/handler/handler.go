package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/appcontext"
	"github.com/itsLeonB/cv-api/dto"
	"github.com/itsLeonB/cv-api/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) GetShortSummary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := QueryNumeric(ctx, "id")
		if err != nil {
			ctx.Error(err)
			return
		}

		newCtx := context.WithValue(ctx, appcontext.KeyProfileID, id)

		summary, err := h.svc.GetShortSummary(newCtx)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(summary))
	}
}
