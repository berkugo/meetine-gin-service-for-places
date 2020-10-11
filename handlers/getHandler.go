package handlers

import (
	"github.com/gin-gonic/gin"
)

func GETHandler(route string, r *gin.Engine) {

	handler := r.Group(route)
	//handler.Use(JSONMiddleware())
	handler.GET("/get/:userid", func(con *gin.Context) {

	})
	handler.GET("/get/all", func(con *gin.Context) {

	})

}
