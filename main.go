package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "goproject/handler"
    "github.com/gorilla/mux"

)


func main() {
    r := mux.NewRouter()

    log.SetOutput(os.Stdout)
    r.HandleFunc("/", handler.IndexHandler)
    r.HandleFunc("/add", handler.AddStudentHandler)
    r.HandleFunc("/view", handler.ViewStudentHandler)
    r.HandleFunc("/delete", handler.DeleteStudentHandler)
    r.HandleFunc("/all-student",handler.ShowAllStudentsHandler)

    r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
    http.Handle("/", r)

    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
    port := "8080"
    fmt.Printf("Server listening on port %s...\n", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

