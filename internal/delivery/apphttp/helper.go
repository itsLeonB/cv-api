package apphttp

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
)

func QueryNumeric(ctx *gin.Context, key string, isRequired bool) (int, bool, error) {
	methodName := fmt.Sprintf("QueryNumeric(key: %s)", key)
	param := ctx.Query(key)
	if param == "" {
		if isRequired {
			return 0, false, apperror.NewAppError(
				httperror.BadRequestError(fmt.Sprintf("query param %s is required", key)),
				"", methodName, "param == \"\" && isRequired",
			)
		}
		return 0, false, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, false, apperror.NewAppError(
			err, "", methodName, fmt.Sprintf("strconv.Atoi(param: %s)", param),
		)
	}

	return val, true, nil
}
