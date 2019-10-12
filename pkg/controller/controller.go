package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping
// @Description ping
// @Success 200
// @Failure 400
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"pong": "ok",
	})
}
