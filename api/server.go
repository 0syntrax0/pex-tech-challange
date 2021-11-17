package api

import (
	"pex/api/controllers/fibonacci"

	"github.com/gin-gonic/gin"
)

type API struct {
	*gin.Engine

	Fibonacci *fibonacci.Client
}

// get everything started
func Run() {
	var api API

	// init Fibboanacci
	api.initFibboanacci()

	// creates a router without any middleware by default
	r := gin.New()

	// global middleware
	r.Use(gin.Logger())

	// recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// init url routes
	api.initRoutes(r)

	// run server in port 8080
	r.Run(":8080")
}

func (a *API) initFibboanacci() {
	a.Fibonacci = fibonacci.New()
}
