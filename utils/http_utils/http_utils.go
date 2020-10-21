package http_utils

import (
	"encoding/json"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err errors.RestErr){
	ResponseJson(w, err.Status, err)
}