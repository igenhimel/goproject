package handler

import (
    "net/http"
    "html/template"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.WriteHeader(http.StatusNotFound)
        tmpl := template.Must(template.ParseFiles("templates/404.html"))
        tmpl.Execute(w, nil)
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
