package http

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Request struct {
    *http.Request
    Params httprouter.Params
}





