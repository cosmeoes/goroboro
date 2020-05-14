package http

import "net/http"

type Text struct {
    Message string
}

func (r Text) Respond(responseWriter http.ResponseWriter) error {
    _, error := responseWriter.Write([]byte(r.Message))
    return error
}

