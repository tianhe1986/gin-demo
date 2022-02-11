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

func TestRouterSub(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	var req *http.Request

	// 有一个参数不对
	req, _ = http.NewRequest(http.MethodGet, "/sub?a=99&b=no", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/sub?a=no&b=10", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())

	// 两个参数都不填，默认是0
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/sub", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0", w.Body.String())

	// 一个默认参数
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/sub?a=33", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "33", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/sub?b=21", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "-21", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/sub?a=63&b=62", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", w.Body.String())
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
