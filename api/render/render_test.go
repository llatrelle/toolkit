package render

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":{"tests":"is_test"},"status":"success","code":200,"message":""}`

	m := make(map[string]interface{})
	m["tests"] = "is_test"
	Success(w, m, nil, http.StatusOK)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessWithOutCode(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":{"tests":"is_test"},"status":"success","code":200,"message":""}`

	m := make(map[string]interface{})
	m["tests"] = "is_test"
	Success(w, m, nil)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":null,"status":"error","code":500,"message":"error"}`

	Error(w, "error", http.StatusInternalServerError)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestErrorWithOutCode(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":null,"status":"error","code":500,"message":"error"}`

	Error(w, "error")

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFail(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":null,"status":"fail","code":500,"message":"error"}`

	Fail(w, "error", http.StatusInternalServerError)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSuccessWithWarning(t *testing.T) {
	w := httptest.NewRecorder()
	expected := `{"data":{"tests":"is_test"},"status":"warning","code":202,"message":"Some warn...","metadata":"some metadata..."}`

	m := make(map[string]interface{})
	m["tests"] = "is_test"
	SuccessWithWarning(w, m, "Some warn...", "some metadata...", "some err...", http.StatusAccepted)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusAccepted, w.Code)
}
