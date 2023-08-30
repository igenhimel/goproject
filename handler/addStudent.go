package handler

import (
    "io"
    "net/http"
    "os"
    "strconv"
	"goproject/models"
	"time"
	"fmt"
)


func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
		start := time.Now()
        r.ParseMultipartForm(10 << 20) // 10 MB limit for uploaded files

        file, handler, err := r.FormFile("image")
        if err != nil {
            http.Error(w, "Error uploading image", http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Create the "images" directory if it doesn't exist
        err = os.MkdirAll("images", os.ModePerm)
        if err != nil {
            http.Error(w, "Error creating directory", http.StatusInternalServerError)
            return
        }

        imagePath := "images/" + handler.Filename
        f, err := os.Create(imagePath)
        if err != nil {
            http.Error(w, "Error saving image", http.StatusInternalServerError)
            return
        }
        defer f.Close()

        _, err = io.Copy(f, file)
        if err != nil {
            http.Error(w, "Error saving image", http.StatusInternalServerError)
            return
        }

        id := r.FormValue("id")
        name := r.FormValue("name")
        cgpa := r.FormValue("cgpa")
        careerInterest := r.FormValue("career_interest")

        cgpaFloat, _ := strconv.ParseFloat(cgpa, 64)
        student := models.Student{
            ID:            id,
            Name:          name,
            CGPA:          cgpaFloat,
            CareerInterest: careerInterest,
            ImageURL:      imagePath,
        }

        // Create a channel to communicate the success of adding the student
        addChan := make(chan bool)

        go func() {
            _, exists := models.Students[id]
            if !exists {
                models.Students[id] = student
            }
            addChan <- !exists
        }()

        added := <-addChan
        

        if added {
            http.Redirect(w, r, "/?added_student_id="+id, http.StatusFound)
			elapsed := time.Since(start)
            fmt.Println("Execution time:", elapsed)
			
        } else {
            http.Redirect(w, r, "/?existing_student_id="+id, http.StatusFound)
        }
    }
}
