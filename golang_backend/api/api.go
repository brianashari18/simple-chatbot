package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang_backend/helper"
	"golang_backend/models"
	"io"
	"net/http"
	"strings"
)

func AskHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := models.QuestionRequest{}
	helper.ReadFromRequest(r, &request)

	questionResponse, err := callPythonAI(request)
	helper.PanicErr(err)

	response := models.Response{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponse(w, &response)
	helper.PanicErr(err)
}

func callPythonAI(question models.QuestionRequest) (models.QuestionResponse, error) {
	aiUrl := "http://localhost:5000/process_question" // URL API Python

	bytes, err := json.Marshal(question)
	helper.PanicErr(err)

	response, err := http.Post(aiUrl, "application/json", strings.NewReader(string(bytes)))
	helper.PanicErr(err)
	defer func() {
		err := response.Body.Close()
		helper.PanicErr(err)
	}()

	body, err := io.ReadAll(response.Body)
	helper.PanicErr(err)

	answerResponse := models.QuestionResponse{}
	err = json.Unmarshal(body, &answerResponse)
	helper.PanicErr(err)

	return answerResponse, nil
}
