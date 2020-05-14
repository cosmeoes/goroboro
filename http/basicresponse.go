package http

import "net/http"

type BasicResponse struct {
    Message string
}

func (r BasicResponse) Respond(responseWriter http.ResponseWriter) error {
    _, error := responseWriter.Write([]byte(r.Message))
    return error
}

