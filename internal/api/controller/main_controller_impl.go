package controller

import (
	"sync"

	services "github.com/Krypt9x/blog-backend/internal/api/service"
	"github.com/Krypt9x/blog-backend/internal/model"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type MainControllerImpl struct {
	Service        services.MainService
	CommentService services.CommentService
	AmountService  services.AmountService
}

func NewMainController(service services.MainService, commentService services.CommentService, amountService services.AmountService) MainController {
	return &MainControllerImpl{
		Service:        service,
		CommentService: commentService,
		AmountService:  amountService,
	}
}

func (controller *MainControllerImpl) Create(ctx *fiber.Ctx) error {
	var data domain.MainDomain
	if err := ctx.BodyParser(data); err != nil {
		helper.PanicIfError(err)
	}
	postData := controller.Service.Create(ctx.Context(), data)

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   postData,
	})
}

func (controller *MainControllerImpl) GetAll(ctx *fiber.Ctx) error {
	data := controller.Service.GetAll(ctx.Context())
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func (controller *MainControllerImpl) GetByUsername(ctx *fiber.Ctx) error {
	data := controller.Service.GetByUsername(ctx.Context(), ctx.Params("username"))
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

// TODO : fix bug in this function
func (controller *MainControllerImpl) GetById(ctx *fiber.Ctx) error {
	data, err := controller.Service.GetById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "not found",
			Data:   "id not found",
		})
	}

	var (
		waitGrup             sync.WaitGroup
		commentDataCh        = make(chan []web.CommentResponse)
		amountViewsDataCh    = make(chan uint64)
		AmountCommentsDataCh = make(chan uint64)
	)

	waitGrup.Add(3)
	go func(wg *sync.WaitGroup) {
		commentData := controller.CommentService.GetById(ctx.Context(), ctx.Params("id"))
		commentDataCh <- commentData
		wg.Done()
	}(&waitGrup)

	go func(wg *sync.WaitGroup) {
		amountViewsData, err := controller.AmountService.UpdateAmountViewsById(ctx.Context(), ctx.Params("id"))
		if err != nil {
			wg.Done()
			helper.PanicIfError(err)
		}
		amountViewsDataCh <- amountViewsData
		wg.Done()
	}(&waitGrup)

	go func(wg *sync.WaitGroup) {
		amountCommentsData, err := controller.AmountService.GetAmountCommentById(ctx.Context(), ctx.Params("id"))
		if err != nil {
			wg.Done()
			helper.PanicIfError(err)
		}
		AmountCommentsDataCh <- amountCommentsData
		wg.Done()
	}(&waitGrup)

	fullData := web.AllTableJoinResponse{
		Id:             data.Id,
		Title:          data.Title,
		Username:       data.Username,
		Date:           data.Date,
		TrailerContent: data.TrailerContent,
		Content:        data.Content,
		AmountComments: <-AmountCommentsDataCh,
		AmountViews:    <-amountViewsDataCh,
		Comments:       <-commentDataCh,
	}

	waitGrup.Wait()
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
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
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "error update",
			Data:   "id user wrong",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   postData,
	})
}

func (controller *MainControllerImpl) Delete(ctx *fiber.Ctx) error {
	err := controller.Service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   fiber.StatusBadRequest,
			Status: "error",
			Data:   "error delete post",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Success Delete Post",
	})
}
