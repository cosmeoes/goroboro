package routing

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
    "fmt"
)

type Router struct {
    router *httprouter.Router
    routes []Route
}

func NewRouter() *Router{
    r := new(Router)
    r.router = httprouter.New()
    return r
}

func (r *Router) GET(path string, handler Handle) {
    route := Route{handler} 
    r.routes = append(r.routes, route)
    r.router.GET(path, route.Handle)
}


func (r *Router) Serve(addr string) error {
    fmt.Println("Starting server in address", addr)
    return http.ListenAndServe(addr, r.router)
}
