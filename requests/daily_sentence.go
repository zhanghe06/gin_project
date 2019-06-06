package requests

type ScoreDailySentenceJsonRequests struct {
	//Score          int       `json:"score" binding:"required"`
	Score int `json:"score" binding:"exists,gte=0,lte=5"`
}

type ReTitleJsonRequests struct {
	//Title string `json:"title" binding:"required,ValidatorUpdateTitleRepetition"`
	Title string `json:"title" binding:"required,validateTitle"`
}

type DeleteDailySentenceUriRequests struct {
	ID string `uri:"id" binding:"required"`
}

type ScoreDailySentenceUriRequests struct {
	ID string `uri:"id" binding:"required"`
}

type UpdateDailySentenceTransactionRequests struct {
	ID string `json:"id" binding:"required"`
	Author string `json:"author"`
	Title string `json:"title"`
	Classification string `json:"classification"`
	Score int `json:"score"`
}
