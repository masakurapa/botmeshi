package http

import (
	"testing"

	"github.com/masakurapa/botmeshi/app/domain/model/http"
	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	res := http.NewResponse(200, "body")

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "body", res.Body)

	headers := map[string]string{
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "OPTIONS,POST",
		"Access-Control-Allow-Origin":  "*",
	}
	for key, val := range headers {
		v, ok := res.Headers[key]
		if assert.True(t, ok) {
			assert.Equal(t, val, v)
		}
	}
}
