package routing

import (
    "fmt"
    ghttp "net/http"
    "github.com/cosmeoes/goroboro/http"
	"github.com/julienschmidt/httprouter"
)

type Route struct {
    Handler Handle
}


func (r *Route) AddHandler(handler Handle) {
    r.Handler = handler
}


func (r *Route) Handle(responseWriter ghttp.ResponseWriter, grequest *ghttp.Request, params httprouter.Params) {
    fmt.Println("Receved GET request", grequest.URL)
    request := &http.Request{Request: grequest, Params: params}
    response := r.Handler(request)
    response.Respond(responseWriter)
}

type Handle func(*http.Request) http.Response

