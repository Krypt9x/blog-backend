package web

type (
	CommentRequest struct {
		Username    string `json:"username"`
		Comment     string `json:"comment"`
		DateComment string `json:"date_comment"`
	}

	MainModelRequest struct {
		Title          string `json:"title"`
		Username       string `json:"username"`
		Date           string `json:"date"`
		TrailerContent string `json:"trailer_content"`
		Content        string `json:"content"`
	}

	UserRequest struct {
		Username string `json:"username" validate:"required, min=6, max=32"`
		Email    string `json:"email" validate:"required, email, min=6"`
		Password string `json:"password" validate:"required, min=6, max=32"`
	}
)
