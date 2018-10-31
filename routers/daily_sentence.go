package routers

import "github.com/zhanghe06/gin_project/controllers"

func RegisterDailySentence() {
	Ver.GET("/daily_sentences", controllers.ListsDailySentenceHandler)
	Ver.GET("/daily_sentence/:id", controllers.GetDailySentenceHandler)
	Ver.POST("/daily_sentence", controllers.CreateDailySentenceHandler)
	Ver.PUT("/daily_sentence/:id", controllers.UpdateDailySentenceHandler)
	Ver.DELETE("/daily_sentence/:id", controllers.DeleteDailySentenceHandler)
	Ver.PATCH("/daily_sentence/:id/score", controllers.ScoreDailySentenceHandler)
}
