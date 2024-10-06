package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler untuk menerima pertanyaan dari Flutter dan meneruskannya ke Python AI
func AskHandler(c *gin.Context) {
	var request struct {
		Question string `json:"question"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Kirim pertanyaan ke API Python
	aiResponse, err := callPythonAI(request.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": aiResponse})
}

// Fungsi untuk mengirimkan request ke Python AI
func callPythonAI(question string) (string, error) {
	aiUrl := "http://localhost:5000/process_question" // URL API Python

	// Format JSON untuk request ke Python
	jsonData, err := json.Marshal(map[string]string{"question": question})
	if err != nil {
		return "", err
	}

	// Kirim request POST ke Python API
	response, err := http.Post(aiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Baca response dari Python API
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result["answer"], nil
}
