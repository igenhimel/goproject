package handler

import (
    "net/http"
    "log"
    "os"
    "goproject/models"
    "strings"
)

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        id := strings.ToUpper(r.URL.Query().Get("id"))

        deleteChan := make(chan bool)

        go func() {
            student, exists := models.Students[id]
            if exists {
                // Delete the student's image file from the "image/" folder
                deleteStudentImage(student.ImageURL)

                // Delete the student from the models.Students map
                delete(models.Students, id)
            }
            log.Printf("Request: %s | Status: %v | URL: %s", r.Method, http.StatusFound, r.URL.Path)
            deleteChan <- exists
        }()

        success := <-deleteChan

        if success {
            http.Redirect(w, r, "/?deleted_student_id="+id, http.StatusFound)
        } else {
            http.Redirect(w, r, "/?student_not_found="+id, http.StatusFound)
        }
    }
}

func deleteStudentImage(imageName string) {
    imagePath := imageName

    // Check if the image file exists
    if _, err := os.Stat(imagePath); err == nil {
        // Delete the image file
        err := os.Remove(imagePath)
        if err != nil {
            log.Printf("Error deleting image: %v", err)
        }
    }
}
