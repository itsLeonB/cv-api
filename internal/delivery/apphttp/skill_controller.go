package apphttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/model"
	"github.com/itsLeonB/cv-api/internal/usecase"
)

type SkillController struct {
	structName   string
	skillUsecase usecase.SkillUsecase
}

func NewSkillController(skillUsecase usecase.SkillUsecase) *SkillController {
	return &SkillController{"SkillController", skillUsecase}
}

func (c *SkillController) HandleInsertCategory() gin.HandlerFunc {
	methodName := "HandleInsertCategory()"
	return func(ctx *gin.Context) {
		request := new(model.NewSkillCategoryRequest)
		err := ctx.ShouldBindJSON(request)
		if err != nil {
			_ = ctx.Error(apperror.NewAppError(
				err, c.structName, methodName,
				"ctx.ShouldBindJSON()",
			))
			return
		}

		response, err := c.skillUsecase.InsertCategory(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, model.NewSuccessResponse(response))
	}
}
