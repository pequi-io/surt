package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	r := New()

	w1 := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w1, req)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w1.Body.String())

	w2 := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/healthz", nil)
	r.ServeHTTP(w2, req)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, "{\"status\":\"UP\"}", w2.Body.String())
}
