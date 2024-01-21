package order

import "testing"

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
