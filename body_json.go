package transport

import (
	"bytes"
	"encoding/json"
	"io"
)

type jsonBody struct {
	value any
}

// JSON creates a Body that serializes the provided value as JSON.
//
// The returned Body automatically declares the Content-Type
// "application/json".
//
// Example:
//
//	req := &transport.Request{
//	    Method: http.MethodPost,
//	    Path:   "/orders",
//	    Body:   transport.JSON(order),
//	}
func JSON(v any) Body {
	return jsonBody{value: v}
}

func (b jsonBody) Reader() (io.Reader, error) {
	data, err := json.Marshal(b.value)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (b jsonBody) ContentType() string {
	return "application/json"
}
