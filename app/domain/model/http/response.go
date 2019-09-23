package http

const (
	// StatusOK status code
	StatusOK = 200
)

// Response struct
type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

// NewResponse returns Response instance
func NewResponse(status int, body string) Response {
	return Response{
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "OPTIONS,POST",
			"Access-Control-Allow-Origin":  "*",
			"X-Slack-No-Retry":             "1",
		},
		StatusCode: status,
		Body:       body,
	}
}
