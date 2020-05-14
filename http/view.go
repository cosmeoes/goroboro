package http

import (
	"html/template"
	"net/http"
	"strings"
)

type View struct { 
    Name string
    Model interface{}
}


func (v View) Respond(responseWriter http.ResponseWriter) error {
    t, err := template.ParseFiles(v.findViewFile())

    if err != nil {
        return err
    }

    responseWriter.Header().Add("Content-Type", "text/html")
    return t.Execute(responseWriter, v.Model)
}

func (v View) findViewFile() (filepath string) {
    filepath = "view/" + strings.Replace(v.Name, ".", "/", -1) + ".html"
    return filepath
}

