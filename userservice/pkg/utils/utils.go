package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}

func SendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Error("Error Encoding Response", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Error encoding response")
	}
}


func SendErrorResponse(w http.ResponseWriter, code int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil {
        logrus.Error("Error encoding error response: ", err)
    }
}