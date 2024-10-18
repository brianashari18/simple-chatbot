package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_backend/app"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	router := app.NewRouter()
	requestBody := strings.NewReader(`{"question":"Gimana cuaca hari ini?"}`)
	request := httptest.NewRequest(http.MethodPost, "https://localhost:3000/api/ask", requestBody)
	recorder := httptest.NewRecorder()
	request.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	var responseBody map[string]interface{}
	res, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	err = json.Unmarshal(res, &responseBody)
	assert.Nil(t, err)

	assert.Equal(t, float64(200), responseBody["code"])
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Hari ini cerah.", responseBody["data"].(map[string]interface{})["answer"])
}
