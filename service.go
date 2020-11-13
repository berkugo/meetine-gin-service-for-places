package main

import (
	"github.com/gin-gonic/gin"
)

var portNumber string = ":5011"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	SetRoutesForEngine(r)
	r.Run(portNumber)
}
