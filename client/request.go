package client

// Request describes a transport request executed
// by a communication client.
type Request struct {
	Connection string

	Method string
	Path   string

	Headers map[string]string
	Query   map[string]string

	Body any
}
