package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/model"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				model.NewErrorResponse(ctx.Errors[0].Error()),
			)
			return
		}
	}
}
