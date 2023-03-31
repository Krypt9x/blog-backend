package helper

import (
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

// must be call waitgrup.wait to avoid deadlock goroutine
func ToAllJoinTable(data web.MainModelResponse, amountCommentsDataCh chan uint64, amountViewsDataCh chan uint64, commentDataCh chan []web.CommentResponse) web.AllTableJoinResponse {
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
