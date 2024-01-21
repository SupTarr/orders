package router

import (
	"github.com/SupTarr/orders/order"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(order.Context)

type Router struct {
	*gin.Engine
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.Engine.POST(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

type Context struct {
	*gin.Context
}

func (c *Context) Order() (o order.Order, err error) {
	err = c.Context.ShouldBindJSON(&o)
	return
}

func (c *Context) JSON(code int, v interface{}) {
	c.Context.JSON(code, v)
}
