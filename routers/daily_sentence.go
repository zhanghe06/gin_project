package routers

import "github.com/zhanghe06/gin_project/controllers"

func RegisterDailySentence() {
	RouterGroupVer.GET("/daily_sentences", controllers.ListsDailySentenceHandler)
	RouterGroupVer.GET("/daily_sentence/:id", controllers.GetDailySentenceHandler)
	RouterGroupVer.POST("/daily_sentence", controllers.CreateDailySentenceHandler)
	RouterGroupVer.POST("/daily_sentence/transaction", controllers.UpdateDailySentenceTransactionHandler)
	RouterGroupVer.PUT("/daily_sentence/:id", controllers.UpdateDailySentenceHandler)
	RouterGroupVer.PUT("/daily_sentence/:id/title", controllers.ReTitleDailySentenceHandler)
	RouterGroupVer.DELETE("/daily_sentence/:id", controllers.DeleteDailySentenceHandler)
	RouterGroupVer.PATCH("/daily_sentence/:id/score", controllers.ScoreDailySentenceHandler)
}
