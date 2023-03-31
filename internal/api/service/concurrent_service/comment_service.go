package concurrentservice

import (
	"sync"

	services "github.com/Krypt9x/blog-backend/internal/api/service"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/gofiber/fiber/v2"
)

type ConcurrentCommentService interface {
	GetCommentByIdPost(ctx *fiber.Ctx, commentsCh chan []web.CommentResponse, wg *sync.WaitGroup)
}

type ConcurrentCommentServiceImpl struct {
	CommentService services.CommentService
}

func NewCommentService(service services.CommentService) ConcurrentCommentService {
	return &ConcurrentCommentServiceImpl{
		CommentService: service,
	}
}

func (service *ConcurrentCommentServiceImpl) GetCommentByIdPost(ctx *fiber.Ctx, commentsCh chan []web.CommentResponse, wg *sync.WaitGroup) {
	dataComment := service.CommentService.GetById(ctx.Context(), ctx.Params("id"))
	commentsCh <- dataComment
	wg.Done()
}
