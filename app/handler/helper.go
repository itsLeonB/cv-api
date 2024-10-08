package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryNumeric(ctx *gin.Context, key string) (int64, error) {
	param := ctx.Query(key)
	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return int64(val), nil
}
