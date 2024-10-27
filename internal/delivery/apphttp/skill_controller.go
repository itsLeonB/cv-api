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

func (c *SkillController) HandleCreateCategory() gin.HandlerFunc {
	methodName := "HandleCreateCategory()"
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

		response, err := c.skillUsecase.CreateCategory(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, model.NewSuccessResponse(response))
	}
}

func (c *SkillController) HandleGetAllCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, err := c.skillUsecase.GetAllCategories(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(categories))
	}
}

func (c *SkillController) HandleGetCategoryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := GetNumericPathParam(ctx, "id")
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		category, err := c.skillUsecase.GetCategoryByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(category))
	}
}
