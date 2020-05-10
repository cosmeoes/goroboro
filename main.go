package main

import (
	"github.com/cosmeoes/goroboro/http"
	"github.com/cosmeoes/goroboro/routing"
)
func main() {
    r := routing.NewRouter()
    r.GET("/", func (_ *http.Request) http.Response {
        return http.BasicResponse{Message: "General knobi"}
    });
    r.Serve(":8080")
}
