package handler

import "github.com/gin-gonic/gin"

type ret struct {
	Username string `json:"username" example:"myUsername"`
	Password string `json:"password" example:"myPassword"`
}

// GetApi godoc
// @Summary TEST Swagger
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Router / [get]
// @Success 200 {object} handler.ret
func GetApi() gin.HandlerFunc {
	return func(c *gin.Context) {

		var x ret

		x.Username = "testUser"
		x.Password = "testPassword"

		c.JSON(200, x)
	}
}
