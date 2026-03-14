package apitest

import (
	"github.com/entiqon/transport"
)

func NewRequest(method, url string) *transport.Request {
	return &transport.Request{
		Method: method,
		Path:   url,
	}
}
