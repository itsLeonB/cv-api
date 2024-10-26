package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
	"github.com/itsLeonB/cv-api/internal/model"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if err := ctx.Errors.Last(); err != nil {
			switch e := err.Err.(type) {
			case *apperror.AppError:
				switch wrappedErr := e.Err.(type) {
				case *httperror.HttpError:
					abortWithError(ctx, wrappedErr.StatusCode, wrappedErr)
				default:
					log.Printf("unhandled error: %T: %s\n", wrappedErr, wrappedErr.Error())
					abortWithError(ctx, http.StatusInternalServerError, httperror.InternalServerError())
				}
			default:
				log.Printf("unwrapped error: %T: %s\n", e, e.Error())
				abortWithError(ctx, http.StatusInternalServerError, httperror.InternalServerError())
			}
			return
		}
	}
}

func abortWithError(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, model.NewErrorResponse(err))
}
