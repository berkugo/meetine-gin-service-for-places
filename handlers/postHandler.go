package handlers

import (
	"checkin/models"
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
	handler.POST("/add", func(con *gin.Context) {

		var bindItem models.Checkin
		if err := con.ShouldBindJSON(&bindItem); err != nil {
			con.AbortWithError(400, err)
		} else {
			response, status := bindItem.AddCheckin()
			con.JSON(status, gin.H{"code": status, "result": map[string]string{"message": response}})
		}
	})

}
