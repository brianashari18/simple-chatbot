package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang_backend/models"
	"testing"
)

func TestMarshal(t *testing.T) {
	request := models.QuestionRequest{Question: "Test Question"}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)

	fmt.Println(string(bytes))
}
