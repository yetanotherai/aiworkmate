package api

import "github.com/gin-gonic/gin"

func RegisterRoute(r *gin.Engine) {
	api := r.Group("/api")

	docAPI := api.Group("/doc")
	docAPI.POST("")
	docAPI.GET("")
	docAPI.GET("/:docid/fragments")
	docAPI.PUT("/:docid/fragment/:fragmentid")

	botAPI := api.Group("/bot")
	botAPI.POST("")
	botAPI.PUT("")
	botAPI.GET("/list")
	botAPI.GET("")
}