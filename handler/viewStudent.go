package handler

import (
    "html/template"
    "net/http"
    "log"
	"goproject/models"
    "strings"
)


func ViewStudentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl := template.Must(template.ParseFiles("templates/view.html"))
        id := strings.ToUpper(r.URL.Query().Get("id"))

        
        // Create a channel to receive the student and whether it exists
        resultChan := make(chan struct {
            student models.Student
            exists  bool
        })

        go func() {
            student, exists := models.Students[id]
            resultChan <- struct {
                student models.Student
                exists  bool
            }{student, exists}
        }()

        result := <-resultChan

        log.Printf("Request: %s | Status: %v | URL: %s", r.Method, http.StatusOK, r.URL.Path)
        if result.exists {
            tmpl.Execute(w, result.student)
        } else {
            http.Redirect(w, r, "/?student_not_found=true", http.StatusFound)
        }
    }
}