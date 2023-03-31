package concurrentservice

import (
	"sync"

	services "github.com/Krypt9x/blog-backend/internal/api/service"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type ConcurrentAmountService interface {
	GetAmountCommentById(ctx *fiber.Ctx, amountCommentsCh chan uint64, wg *sync.WaitGroup)
	UpdateAmountViewsById(ctx *fiber.Ctx, amountViewsCh chan uint64, wg *sync.WaitGroup)
}

type ConcurrentAmountServiceImpl struct {
	AmountService services.AmountService
}

func NewAmountService(service services.AmountService) ConcurrentAmountService {
	return &ConcurrentAmountServiceImpl{
		AmountService: service,
	}
}

func (service *ConcurrentAmountServiceImpl) GetAmountCommentById(ctx *fiber.Ctx, amountCommentsCh chan uint64, wg *sync.WaitGroup) {
	amountCommentData, err := service.AmountService.GetAmountCommentById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		helper.PanicIfError(err)
	}

	amountCommentsCh <- amountCommentData
	wg.Done()
}

func (service *ConcurrentAmountServiceImpl) UpdateAmountViewsById(ctx *fiber.Ctx, amountViewsCh chan uint64, wg *sync.WaitGroup) {
	updatedAmountViewsData, err := service.AmountService.UpdateAmountViewsById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		helper.PanicIfError(err)
	}

	amountViewsCh <- updatedAmountViewsData
	wg.Done()
}
