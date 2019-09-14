package http

const (
	// StatusOK status code
	StatusOK = 200
)

// Response struct
type Response struct {
	Headers    map[string]string
	StatusCode int
	Body       string
}

// NewResponse returns Response instance
func NewResponse(status int, body string) Response {
	return Response{
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "OPTIONS,POST",
			"Access-Control-Allow-Origin":  "*",
		},
		StatusCode: status,
		Body:       body,
	}
}
