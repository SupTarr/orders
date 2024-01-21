package order

import (
	"errors"
	"net/http"
	"testing"
)

type MockContext struct {
	channel  string
	code     int
	response map[string]string
}

func (c *MockContext) Order() (Order, error) {
	return Order{
		SalesChannel: c.channel,
	}, nil
}

func (c *MockContext) JSON(code int, v interface{}) {
	c.code = code
	c.response = v.(map[string]string)
}

func TestOnlyAcceptOnlineChannel(t *testing.T) {
	handler := &Handler{
		channel: "Online",
	}

	c := &MockContext{channel: "Offline"}
	handler.Order(c)

	want := "Offline is not accepted"

	if want != c.response["message"] {
		t.Errorf("%q is expected but got %q\n", want, c.response["message"])
	}
}

func TestOnlyAcceptOfflineChannel(t *testing.T) {
	handler := &Handler{
		channel: "Offine",
	}

	c := &MockContext{channel: "Online"}
	handler.Order(c)

	want := "Online is not accepted"

	if want != c.response["message"] {
		t.Errorf("%q is expected but got %q\n", want, c.response["message"])
	}
}

type MockContextBadRequest struct {
	code     int
	response map[string]string
}

func (c *MockContextBadRequest) Order() (Order, error) {
	return Order{}, errors.New("Order went wrong")
}

func (c *MockContextBadRequest) JSON(code int, v interface{}) {
	c.code = code
	c.response = v.(map[string]string)
}

func TestBadRequestOrderWentWrong(t *testing.T) {
	handler := &Handler{}

	c := &MockContextBadRequest{}
	handler.Order(c)

	want := http.StatusBadRequest
	if want != c.code {
		t.Errorf("%d status code is expected but got %d\n", want, c.code)
	}
}

type MockContextBadRequestWithChannel struct {
	channel         string
	jsonCalledCount int
}

func (c *MockContextBadRequestWithChannel) Order() (Order, error) {
	return Order{SalesChannel: c.channel}, errors.New("Order went wrong")
}

func (c *MockContextBadRequestWithChannel) JSON(code int, v interface{}) {
	c.jsonCalledCount++
}

func TestOnlyCalledJSONOneTime(t *testing.T) {
	handler := &Handler{
		channel: "Offline",
	}

	c := &MockContextBadRequestWithChannel{}
	handler.Order(c)

	want := 1
	if want != c.jsonCalledCount {
		t.Errorf("it should called one time but got %d times\n", c.jsonCalledCount)
	}
}

type MockStore struct {
	wasCalled bool
}

func (m *MockStore) Save(Order) error {
	m.wasCalled = true
	return nil
}

func TestOrderWasSaved(t *testing.T) {
	store := &MockStore{}
	handler := &Handler{
		channel: "Online",
		store:   store,
	}

	c := &MockContext{channel: "Online"}
	handler.Order(c)

	want := true
	if want != store.wasCalled {
		t.Error("It should store order data")
	}
}

type MockFailStore struct{}

func (m *MockFailStore) Save(Order) error {
	return errors.New("Saving order to DB went wrong")
}

func TestOrderFailAtSave(t *testing.T) {
	store := &MockFailStore{}
	handler := &Handler{
		channel: "Online",
		store:   store,
	}

	c := &MockContext{channel: "Online"}
	handler.Order(c)

	want := http.StatusInternalServerError
	if want != c.code {
		t.Errorf("%d status code is expected but got %d\n", want, c.code)
	}
}

func TestOrderIsOk(t *testing.T) {
	store := &MockStore{}
	handler := &Handler{
		channel: "Online",
		store:   store,
	}

	c := &MockContext{channel: "Online"}
	handler.Order(c)

	_, ok := c.response["message"]
	if !ok {
		t.Error("Message key is expected")
	}
}
