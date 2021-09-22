package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imufiid/goabsen/app/core"
)

type Router interface {
	Router()
}

type router struct {
	Addrs string
	Handler *core.BaseHandler
}

func NewRouter(addrs string, handler *core.BaseHandler) *router {
	return &router{
		Addrs: addrs, Handler: handler,
	}
}

func (r *router) Router() {
	router := gin.Default()
	router.Use(cors.Default())
	
	// apiV1 := router.Group("/api/v1")
	
	// apiV1.GET("/users", r.Handler.UserHandler.CheckHandler)

	http.ListenAndServe(r.Addrs, router)
	router.Run()
}
