package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultRoutePing(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/ping", nil)
	w := httptest.NewRecorder()
	defaultRoutes().ServeHTTP(w, req)

	assert.Equal(t, "pong", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDefaultRouteRoot(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	w := httptest.NewRecorder()
	defaultRoutes().ServeHTTP(w, req)

	assert.Equal(t, "hi", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDefaultRouteNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/someurl", nil)
	w := httptest.NewRecorder()
	defaultRoutes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
