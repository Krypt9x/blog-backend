package concurrentservice

type ConcurrentService struct {
	AmountService  ConcurrentAmountService
	CommentService ConcurrentCommentService
}

func NewConcurrentService(amountService ConcurrentAmountService, commentService ConcurrentCommentService) *ConcurrentService {
	return &ConcurrentService{
		AmountService:  amountService,
		CommentService: commentService,
	}
}
