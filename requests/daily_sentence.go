package requests

type ScoreDailySentenceRequests struct {
	//Score          int       `json:"score" binding:"required"`
	Score          int       `json:"score" binding:"exists,gte=0,lte=5"`
}


type ReTitleRequests struct {
	Title          string    `json:"title" binding:"required"`
}
