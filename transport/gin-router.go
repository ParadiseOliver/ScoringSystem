package transport

import "github.com/gin-gonic/gin"

type ginRouter struct{}

type Router interface {
	GET(uri string, f func(c *gin.Context))
	POST(uri string, f func(c *gin.Context))
	SERVE(port string)
}

func NewGinRouter() Router {
	return &ginRouter{}
}

func (r ginRouter) GET(uri string, f func(c *gin.Context)) {

}

func (r ginRouter) POST(uri string, f func(c *gin.Context)) {

}

func (r ginRouter) SERVE(port string) {

}
