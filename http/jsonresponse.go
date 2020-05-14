package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
    Data interface{} 
}

func (r JsonResponse) Respond(responseWriter http.ResponseWriter) error {
    return fmt.Errorf("B-baka")
    responseWriter.Header().Add("Content-Type", "application/json")
    json, err := json.Marshal(r.Data)
    if err != nil {
        return err
    }
    _, err = responseWriter.Write(json)
    return err
}
