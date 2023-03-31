package controller

import (
	"sync"

	services "github.com/Krypt9x/blog-backend/internal/api/service"
	concurrentservice "github.com/Krypt9x/blog-backend/internal/api/service/concurrent_service"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type MainControllerImpl struct {
	Service           services.MainService
	ConcurrentService concurrentservice.ConcurrentService
}

func NewMainController(service services.MainService, concurrentService concurrentservice.ConcurrentService) MainController {
	return &MainControllerImpl{
		Service:           service,
		ConcurrentService: concurrentService,
	}
}

func (controller *MainControllerImpl) Create(ctx *fiber.Ctx) error {
	var data domain.MainDomain
	if err := ctx.BodyParser(data); err != nil {
		helper.PanicIfError(err)
	}
	postData := controller.Service.Create(ctx.Context(), data)

	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   postData,
	})
}

func (controller *MainControllerImpl) GetAll(ctx *fiber.Ctx) error {
	data := controller.Service.GetAll(ctx.Context())
	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func (controller *MainControllerImpl) GetByUsername(ctx *fiber.Ctx) error {
	data := controller.Service.GetByUsername(ctx.Context(), ctx.Params("username"))
	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func (controller *MainControllerImpl) GetById(ctx *fiber.Ctx) error {
	data, err := controller.Service.GetById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   fiber.StatusNotFound,
			Status: "not found",
			Data:   "id not found",
		})
	}

	var (
		waitGrup             sync.WaitGroup
		commentDataCh        = make(chan []web.CommentResponse)
		amountViewsDataCh    = make(chan uint64)
		amountCommentsDataCh = make(chan uint64)
	)

	waitGrup.Add(3)
	go controller.ConcurrentService.AmountService.UpdateAmountViewsById(ctx, amountViewsDataCh, &waitGrup)
	go controller.ConcurrentService.CommentService.GetCommentByIdPost(ctx, commentDataCh, &waitGrup)
	go controller.ConcurrentService.AmountService.GetAmountCommentById(ctx, amountCommentsDataCh, &waitGrup)

	fullData := helper.ToAllJoinTable(data, amountCommentsDataCh, amountViewsDataCh, commentDataCh)

	waitGrup.Wait()
	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   fullData,
	})
}

func (controller *MainControllerImpl) UpdateById(ctx *fiber.Ctx) error {
	var data domain.MainDomain
	if err := ctx.BodyParser(data); err != nil {
		helper.PanicIfError(err)
	}

	postData, err := controller.Service.UpdateById(ctx.Context(), data)
	if err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   fiber.StatusNotFound,
			Status: "error update",
			Data:   "id user wrong",
		})
	}

	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   postData,
	})
}

func (controller *MainControllerImpl) Delete(ctx *fiber.Ctx) error {
	err := controller.Service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
			Code:   500,
			Status: "error delete",
			Data:   "internal server error",
		})
	}
	return helper.SendJSONResponse(ctx, &helper.HttpResponseData{
		Code:   200,
		Status: "OK",
		Data:   "success delete",
	})
}
