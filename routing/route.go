package routing

import (
    "fmt"
    ghttp "net/http"
    "github.com/cosmeoes/goroboro/http"
	"github.com/julienschmidt/httprouter"
)

type Route struct {
    handler Handle
}

// Adds a handler function for this route
func (r *Route) AddHandler(handler Handle) {
    r.handler = handler
}

func (r *Route) Handle(responseWriter ghttp.ResponseWriter, grequest *ghttp.Request, params httprouter.Params) {
    fmt.Println("Receved GET request", grequest.URL)
    request := http.Request{Request: grequest, Params: params}
    response := r.handler(&request)
    err := response.Respond(responseWriter)
    if err != nil {
        responseWriter.WriteHeader(ghttp.StatusInternalServerError)
        responseWriter.Write([]byte(err.Error()))
    }
}

type Handle func(*http.Request) http.Response

