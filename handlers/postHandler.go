package handlers

import (
	"checkin/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func POSTHandler(route string, r *gin.Engine) {

	handler := r.Group(route)
	//handler.Use(JSONMiddleware())
	handler.POST("/add", func(con *gin.Context) {

		currentList := models.GetInstance()
		var bindItem models.Checkin
		if err := con.ShouldBindJSON(&bindItem); err != nil {
			con.AbortWithError(400, err)
		} else {
			currentList.AddCheckin(&bindItem)
			for _, item := range currentList.GetCheckins() {
				fmt.Println(item.UserId)
			}
			con.JSON(200, gin.H{"message": "Checked in."})
		}
	})

}
