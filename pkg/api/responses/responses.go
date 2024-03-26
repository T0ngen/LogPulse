package responses

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Description ErrorDescription
}

type ErrorDescription struct {
	ErrorCode           int    `json:"error_code"`
	DetailedDescription string `json:"detailed_description"`
	TagError            string
}


func WriteResponse(w http.ResponseWriter, dataJSON []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	_, err := w.Write(dataJSON)
	if err != nil {
		logrus.Errorf("error in writing response body: %s", err)
	}
}