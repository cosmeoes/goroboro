package http

import (
	"encoding/json"
	"net/http"
)

type Json struct {
    Data interface{} 
}

func (r Json) Respond(responseWriter http.ResponseWriter) error {
    responseWriter.Header().Add("Content-Type", "application/json")
    json, err := json.Marshal(r.Data)
    if err != nil {
        return err
    }
    _, err = responseWriter.Write(json)
    return err
}
