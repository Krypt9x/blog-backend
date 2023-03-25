package web

type (
	AmountResponse struct {
		AmountComment uint64
		AmountView    uint64
	}

	CommentResponse struct {
		Id          string
		Username    string
		Comment     string
		DateComment string
	}

	AllTableJoinResponse struct {
		Id             string
		Title          string
		Username       string
		Date           string
		TrailerContent string
		Content        string
		AmountComments int64
		AmountViews    int64
		Comments       []CommentResponse
	}

	MainModelResponse struct {
		Id             string
		Title          string
		Username       string
		Date           string
		TrailerContent string
		Content        string
	}

	UserResponse struct {
		Username string
		Email    string
		Password string
	}
)
