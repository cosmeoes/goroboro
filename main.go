package main

import (
	// "github.com/cosmeoes/goroboro/http"
	// "github.com/cosmeoes/goroboro/routing"
	"github.com/cosmeoes/goroboro/moneta"
)

type Costumer struct {
    name string
    password string
}

func main() {
    moneta.Find(Costumer{}).Where("name", "=", "Cosme").WhereNotNull("password").Get("name", "password")
    // r := routing.NewRouter()
    // r.GET("/", func (_ *http.Request) http.Response {
    //     return http.BasicResponse{Message: "General knobi"}
    // });
    // r.Serve(":8080")
}
