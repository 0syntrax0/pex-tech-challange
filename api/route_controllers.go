package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) current(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"current number (sequence moved up)": a.Fibonacci.GetCurrent(),
	})
}

func (a *API) next(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"next number in current sequence": a.Fibonacci.GetNext(),
	})
}

func (a *API) previous(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"previous number in current sequence": a.Fibonacci.GetPrevious(),
	})
}
