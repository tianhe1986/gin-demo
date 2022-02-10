package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterBase(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World!", w.Body.String())
}

func TestRouterAdd(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	var req *http.Request

	req, _ = http.NewRequest(http.MethodGet, "/add/gg/6", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/add/4/ok", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/add/7/6", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "13", w.Body.String())
}

func TestRouterMul(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	var req *http.Request
	var jsonStr []byte
	var jsonMap map[string]interface{}

	jsonStr = []byte(`{"a": 11, "b": "no"}`)
	req, _ = http.NewRequest(http.MethodPost, "/mul", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &jsonMap)
	assert.Equal(t, "error", jsonMap["result"])

	w = httptest.NewRecorder()
	jsonStr = []byte(`{"a": "ok", "b": 4}`)
	req, _ = http.NewRequest(http.MethodPost, "/mul", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &jsonMap)
	assert.Equal(t, "error", jsonMap["result"])

	w = httptest.NewRecorder()
	jsonStr = []byte(`{"a": 12, "b": 6}`)
	req, _ = http.NewRequest(http.MethodPost, "/mul", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &jsonMap)
	assert.Equal(t, 72, int(jsonMap["result"].(float64)))
}
