package api

import (
	"github.com/gin-gonic/gin"
)

func (a *API) initRoutes(r *gin.Engine) {
	r.GET("/current", a.current)
	r.GET("/next", a.next)
	r.GET("/previous", a.previous)
}
