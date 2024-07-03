package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAPI(t *testing.T) {
	const PORT = "8080"
	api := NewApi()
	api.SetPort(PORT)

	assert.NotNil(t, api)
	assert.Equal(t, PORT, api.port)

}

func TestAPIIni(t *testing.T) {
	const PORT = "8080"
	api := NewApi()
	api.init()
	assert.NotNil(t, api)
	assert.Equal(t, PORT, api.port)
	assert.NotNil(t, api.router)
}
