package web

type (
	CommentRequest struct {
		Username    string `json:"username"`
		Comment     string `json:"comment" validate:"required"`
		DateComment string `json:"date_comment"`
	}

	MainModelRequest struct {
		Title          string `json:"title" validate:"required"`
		Username       string `json:"username" validate:"required"`
		Date           string `json:"date"`
		TrailerContent string `json:"trailer_content"`
		Content        string `json:"content" validate:"required"`
	}

	UserRequest struct {
		Username string `json:"username" validate:"required, min=6, max=32"`
		Email    string `json:"email" validate:"required, email, min=6"`
		Password string `json:"password" validate:"required, min=6, max=32"`
	}
)
