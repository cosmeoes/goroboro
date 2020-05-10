package routing

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/julienschmidt/httprouter"
)

var once sync.Once
var instance Router

type Router struct {
    router *httprouter.Router
    routes []Route
}

func NewRouter() *Router{

    once.Do(func() {
            instance := new(Router)
            instance.router = httprouter.New()
    })

    return &instance
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
