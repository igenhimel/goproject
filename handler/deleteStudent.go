package handler

import (
    "net/http"
    "log"
	"goproject/models"
)


func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        id := r.URL.Query().Get("id")

        // Create a channel to communicate the success of deletion
        deleteChan := make(chan bool)

        go func() {
          
            _, exists := models.Students[id] // Check if student with given ID exists
            if exists {
                delete(models.Students, id) // Delete the student if exists
            }
            log.Printf("Request: %s | Status: %v | URL: %s", r.Method, http.StatusFound, r.URL.Path)
            deleteChan <- exists
        }()

        success := <-deleteChan

        if success {
            http.Redirect(w, r, "/?deleted_student_id="+id, http.StatusFound) // Redirect with deleted_student_id query parameter
        } else {
            http.Redirect(w, r, "/?student_not_found="+id, http.StatusFound) // Redirect with delete_failed_student_id query parameter
        }
    }
}