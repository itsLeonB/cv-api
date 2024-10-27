package apphttp

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
)

func GetNumericQueryParam(ctx *gin.Context, key string, isRequired bool) (int, bool, error) {
	methodName := fmt.Sprintf("GetNumericQueryParam(key: %s, isRequired: %t)", key, isRequired)
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

func GetNumericPathParam(ctx *gin.Context, key string) (int, error) {
	methodName := fmt.Sprintf("GetNumericPathParam(key: %s)", key)
	param := ctx.Param(key)
	if param == "" {
		return 0, apperror.NewAppError(
			httperror.BadRequestError(fmt.Sprintf("missing path parameter %s", key)),
			"", methodName, "param == \"\"",
		)
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, apperror.NewAppError(
			err, "", methodName, fmt.Sprintf("strconv.Atoi(param: %s)", param),
		)
	}

	return val, nil
}
