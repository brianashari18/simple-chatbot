package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang_backend/api" // Sesuaikan dengan path package asli

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAskHandler(t *testing.T) {
	// Setup Gin router dan API handler
	router := gin.Default()
	router.POST("/ask", api.AskHandler)

	// Mock pertanyaan dari Flutter
	requestBody, _ := json.Marshal(map[string]string{
		"question": "apa kabar?",
	})

	// Membuat request HTTP POST menggunakan httptest
	req, _ := http.NewRequest(http.MethodPost, "/ask", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Membuat recorder untuk menangkap response dari handler
	w := httptest.NewRecorder()

	// Menjalankan router dengan request
	router.ServeHTTP(w, req)

	// Membaca response body
	resBody, err := io.ReadAll(w.Body)
	assert.Nil(t, err)

	// Assert status code 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Pastikan response tidak kosong
	assert.NotEmpty(t, resBody)

	// Assert bahwa response berisi kunci "answer"
	var response map[string]string
	err = json.Unmarshal(resBody, &response)
	assert.Nil(t, err)

	// Pastikan response memiliki kunci "answer"
	if _, exists := response["answer"]; !exists {
		t.Fatalf("\nResponse tidak mengandung kunci 'answer'")
	}

	// (Opsional) Cetak response untuk debugging
	t.Logf("Response: %v", response)
}
