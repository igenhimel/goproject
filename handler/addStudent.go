package handler

import (
	"fmt"
	"goproject/models"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// AddStudentHandler handles the addition of a new student.
func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		start := time.Now() 
		id := strings.ToUpper(r.FormValue("id"))

		r.ParseMultipartForm(10 << 20) // 10 MB limit for uploaded files

		// Retrieve the uploaded image file
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

		// Define the path for saving the uploaded image
		imagePath := "images/" +id+ handler.Filename
		f, err := os.Create(imagePath)
		if err != nil {
			http.Error(w, "Error saving image", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Copy the uploaded image data to the file
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "Error saving image", http.StatusInternalServerError)
			return
		}

		// Extract student information from the form
		
		name := r.FormValue("name")
		cgpa := r.FormValue("cgpa")
		careerInterest := r.FormValue("career_interest")

		if id == "" || name == "" || cgpa == "" || careerInterest == "" {
            http.Redirect(w, r, "/?empty_fields=true", http.StatusFound)
            return
        }

		cgpaFloat, err := strconv.ParseFloat(cgpa, 64)
		if err != nil || cgpaFloat < 0.00 || cgpaFloat > 4.00 {
           http.Redirect(w, r, "/?invalid_cgpa="+cgpa, http.StatusFound)
		   return
		}


		// Create a Student object with the extracted information
		student := models.Student{
			ID:            id,
			Name:          name,
			CGPA:          cgpaFloat,
			CareerInterest: careerInterest,
			ImageURL:      imagePath,
		}

		// Create a channel to communicate the success of adding the student
		addChan := make(chan bool)

		// Goroutine to add student information to the models
		go func() {
			_, exists := models.Students[id]
			if !exists {
				models.Students[id] = student
			}
			addChan <- !exists
		}()

		added := <-addChan

		if added {
			// Redirect after successfully adding the student
			http.Redirect(w, r, "/?added_student_id="+id, http.StatusFound)
			elapsed := time.Since(start)
			fmt.Println("Execution time:", elapsed)
		} else {
			// Redirect if student already exists
			http.Redirect(w, r, "/?existing_student_id="+id, http.StatusFound)
		}
		
	} else if r.Method == http.MethodGet {
        tmpl := template.Must(template.ParseFiles("templates/404.html"))
        tmpl.Execute(w, nil)
	} else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }

}
