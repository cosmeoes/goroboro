package http

import "net/http"

type BasicResponse struct {
    Message string
}

func (r BasicResponse) Respond(responseWriter http.ResponseWriter) {
    responseWriter.Write([]byte(r.Message))
}

