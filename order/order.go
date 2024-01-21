package order

import (
	"fmt"
	"net/http"
)

type Handler struct {
	channel string
}

type Context interface {
	Order() (Order, error)
	JSON(int, interface{})
}

func (h *Handler) Order(c Context) {
	order, err := c.Order()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if order.SalesChannel != h.channel {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("%s is not accepted", order.SalesChannel),
		})
	}
}