package requests

type ScoreDailySentenceJsonRequests struct {
	//Score          int       `json:"score" binding:"required"`
	Score int `json:"score" binding:"exists,gte=0,lte=5"`
}

type ReTitleJsonRequests struct {
	//Title string `json:"title" binding:"required,ValidatorUpdateTitleRepetition"`
	Title string `json:"title" binding:"required"`
}

type DeleteDailySentenceUriRequests struct {
	ID string `uri:"id" binding:"required"`
}

type ScoreDailySentenceUriRequests struct {
	ID string `uri:"id" binding:"required"`
}
