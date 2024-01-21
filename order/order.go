package order

import (
	"fmt"
	"net/http"
)

type Storer interface {
	Save(Order) error
}

type Handler struct {
	channel string
	store   Storer
}

func NewHandler(channel string, store Storer) *Handler {
	return &Handler{channel: channel, store: store}
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
		return
	}

	if order.SalesChannel != h.channel {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("%s is not accepted", order.SalesChannel),
		})
		return
	}

	if err := h.store.Save(order); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "successful",
	})
}
