package responseUtils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type StandardResponse struct {
	Code         int
	Message      string
	ResponseData interface{}
}

func ResponseNotImplemented() StandardResponse {
	return StandardResponse{
		Code:         http.StatusMethodNotAllowed,
		Message:      "Method not implemented",
		ResponseData: nil,
	}
}

func (s *StandardResponse) RenderResponse(w http.ResponseWriter) {
	status := false
	if s.Code == http.StatusOK {
		status = true
	}

	response := bson.M{
		"responseData": s.ResponseData,
		"message":      s.Message,
		"status":       status,
	}

	indentedResponse, _ := json.MarshalIndent(response, "", "")

	w.Header().Set("Content-Length", fmt.Sprint(len(indentedResponse)))
	w.WriteHeader(s.Code)
	w.Write(indentedResponse)
}
