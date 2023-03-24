package domain

type (
	Amount struct {
		No             int
		Id             string
		AmountComments int64
		AmountViews    int64
	}

	Comments struct {
		Id          string
		Username    string
		Comment     string
		DateComment string
	}

	MainDomain struct {
		No             int
		Id             string
		Title          string
		Username       string
		Date           string
		TrailerContent string
		Content        string
	}

	User struct {
		Username string
		Email    string
		Password string
	}
)
