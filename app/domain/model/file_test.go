package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	f := NewFile("path/to", "name.json", "body")
	assert.Equal(t, "path/to", f.GetPath())
	assert.Equal(t, "name.json", f.GetName())
	assert.Equal(t, "body", f.GetData())
}
