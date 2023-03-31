package helper

import (
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

// after call this function, must be call [sync.WaitGrup.Wait()] to avoid deadlock goroutine
func ToAllJoinTableConcurrent(data web.MainModelResponse, amountCommentsDataCh chan uint64, amountViewsDataCh chan uint64, commentDataCh chan []web.CommentResponse) web.AllTableJoinResponse {
	return web.AllTableJoinResponse{
		Id:             data.Id,
		Title:          data.Title,
		Username:       data.Username,
		Date:           data.Date,
		TrailerContent: data.TrailerContent,
		Content:        data.Content,
		AmountComments: <-amountCommentsDataCh,
		AmountViews:    <-amountViewsDataCh,
		Comments:       <-commentDataCh,
	}
}
