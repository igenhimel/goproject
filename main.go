package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "goproject/handler"
)

func main() {
    log.SetOutput(os.Stdout)
    http.HandleFunc("/", handler.IndexHandler)
    http.HandleFunc("/add", handler.AddStudentHandler)
    http.HandleFunc("/view", handler.ViewStudentHandler)
    http.HandleFunc("/delete", handler.DeleteStudentHandler)
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
    port := "8080"
    fmt.Printf("Server listening on port %s...\n", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

