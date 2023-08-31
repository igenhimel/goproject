package handler

import (
    "net/http"
    "log"
    "goproject/models"
    "strings"
)


func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        id := strings.ToUpper(r.URL.Query().Get("id"))

        deleteChan := make(chan bool)

        go func() {
            _, exists := models.Students[id] 
            if exists {
                delete(models.Students, id)
            }
            log.Printf("Request: %s | Status: %v | URL: %s", r.Method, http.StatusFound, r.URL.Path)
            deleteChan <- exists 
        }()

        success := <-deleteChan 

        if success {
            http.Redirect(w, r, "/?deleted_student_id="+id, http.StatusFound) // Redirect with deleted_student_id query parameter
        } else {
            http.Redirect(w, r, "/?student_not_found="+id, http.StatusFound) // Redirect with student_not_found query parameter
        }
    }
}
