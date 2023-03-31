package controller

import (
	services "github.com/Krypt9x/blog-backend/internal/api/service"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type CommentControllerImpl struct {
	CommentService services.CommentService
}

func (service *CommentControllerImpl) Create(ctx *fiber.Ctx) error {
	var dataReq domain.Comments
	if err := ctx.BodyParser(dataReq); err != nil {
		helper.PanicIfError(err)
	}

	if err := service.CommentService.Create(ctx.Context(), dataReq); err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   fiber.StatusInternalServerError,
			Status: "failed create",
			Data:   "internal server error",
		})
	}

	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "success create comment",
	})

}

func (service *CommentControllerImpl) Update(ctx *fiber.Ctx) error {
	var dataReq domain.Comments
	if err := ctx.BodyParser(dataReq); err != nil {
		helper.PanicIfError(err)
	}

	if err := service.CommentService.UpdateById(ctx.Context(), dataReq); err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   fiber.StatusInternalServerError,
			Status: "failed update comment",
			Data:   "internal server error",
		})
	}

	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "success update comment",
	})
}

func (service *CommentControllerImpl) Delete(ctx *fiber.Ctx) error {
	if err := service.CommentService.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   fiber.StatusInternalServerError,
			Status: "failed delete comment",
			Data:   "internal server error",
		})
	}

	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "success delete comment",
	})
}
