package handlers

import (
	"checkin/request"
	"github.com/gin-gonic/gin"
)

func GETHandler(route string, r *gin.Engine) {

	handler := r.Group(route)
	handler.Use(JSONMiddleware())
	handler.GET("/get/user/:userid", func(con *gin.Context) {
		if len(con.Params.ByName("userid")) > 0 {
			response, err, status := request.GetRequestForPersonalCheckins(con.Params.ByName("userid"))
			if err != nil {
				con.JSON(status, gin.H{"code": status, "result": "Error happened."})
			} else {
				con.JSON(status, response)
			}

		} else {
			con.JSON(404, gin.H{"code": 404, "result": "Error happened."})
		}
	})
	handler.GET("/get/all", func(con *gin.Context) {

		if len(con.Query("city")) > 0 {
			response, err, status := request.SendGetRequestForAll(con.Query("city"))
			if err != nil {
				con.JSON(status, gin.H{"code": status, "result": "Error happened."})
			} else {
				con.JSON(status, response)
			}

		} else {
			con.JSON(404, gin.H{"code": 404, "result": "Error happened."})
		}

	})

}
