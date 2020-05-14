package http

import "net/http"

type Response interface {
   Respond(http.ResponseWriter) error
}

