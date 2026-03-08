package api

// Response represents the result of a transport execution.
type Response struct {
	Status  int
	Headers map[string]string
	Body    []byte
}
