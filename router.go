package main

import (
	"checkin/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutesForEngine(r *gin.Engine) {

	handlers.POSTHandler("/checkin", r)

}
