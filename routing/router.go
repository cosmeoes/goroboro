package routing

import (
	"fmt"
	ghttp "net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/cosmeoes/goroboro/http"
	"github.com/julienschmidt/httprouter"
)

var once sync.Once
var instance *Router

type Router struct {
    router *httprouter.Router
    routes []Route
}

func NewRouter() *Router{
    once.Do(func() {
        instance = new(Router)
        instance.router = httprouter.New()
    })

    return instance
}


func (r *Router) registerHandler(method string, path string, handler interface{}) {
    resolvedHandler, err := resolveHandler(handler)
    if err != nil {
        text, _ := fmt.Printf("Error ocurred registering route %s: %s\n", path, err.Error())
        panic(text)

    }
    route := Route{resolvedHandler}
    r.routes = append(r.routes, route)
    r.router.Handle(method, path, route.Handle)
}


// Register GET route
func (r *Router) GET(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodGet, path, handler)
}

// Register POST route
func (r *Router) POST(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodPost, path, handler)
}

// Register PUT route
func (r *Router) PUT(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodPut, path, handler)
}

// Register PATCH route
func (r *Router) PATCH(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodPatch, path, handler)
}

// Register DELETE route
func (r *Router) DELETE(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodDelete, path, handler)
}

// Register OPTIONS route
func (r *Router) OPTIONS(path string, handler interface{}) {
    r.registerHandler(ghttp.MethodOptions, path, handler)
}

func (r *Router) Serve(addr string) error {
    fmt.Println("Starting server in address", addr)
    return ghttp.ListenAndServe(addr, r.router)
}

func resolveHandler(handler interface{}) (Handle, error) {
    //Check if handler is of type handle
    t := reflect.TypeOf(handler)

    if t.AssignableTo(reflect.TypeOf((Handle)(nil))) {
        if f, ok := handler.(func(* http.Request) http.Response); ok {
            return Handle(f), nil
        }
    }

    if t == reflect.TypeOf("") {
        return handlerFromString(handler.(string))
    } 

    return nil, fmt.Errorf("Wrong type passed to route, second parameter must be of type Handler or a string")
}


func handlerFromString(handerString string) (Handle, error) {
    separeted := strings.Split(handerString, "@")

    if len(separeted) < 2 {
        return nil, fmt.Errorf("Invalid string %s, the string must be formated `ControllerName@MethodName`", handerString)
    }

    controllerName := separeted[0]
    methodName := separeted[1]
    for _, controller := range ControllersRegistry {
        v := reflect.ValueOf(controller)       
        if v.Type().Name() == controllerName {
            controllerType := reflect.New(v.Type()) 
            method := controllerType.MethodByName(methodName)
            if method.Kind() == reflect.Invalid {
                return nil, fmt.Errorf("%s does not have method %s\n", controllerName, methodName)
            }
            return resolveHandler(method.Interface())
        }
    }
    return nil, fmt.Errorf("Unregistered controller %s declared in route\n", controllerName)
}
