package main

import (
	"github.com/gin-gonic/gin"
)

var portNumber string = ":5001"

func main() {
	r := gin.Default()
	SetRoutesForEngine(r)
	r.Run(portNumber)
}
