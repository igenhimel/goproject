package handler

import (
    "html/template"
    "net/http"

)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

    data := struct {
        StudentIDExists   bool
        ExistingStudentID string
        StudentAdded      bool
        AddedStudentID    string
        StudentNotFound   bool
        DeleteStudent     bool
        DeleteStudentID   string
    }{
        StudentIDExists:   r.URL.Query().Get("existing_student_id") != "",
        ExistingStudentID: r.URL.Query().Get("existing_student_id"),
        StudentAdded:      r.URL.Query().Get("added_student_id") != "",
        AddedStudentID:    r.URL.Query().Get("added_student_id"),
        StudentNotFound:   r.URL.Query().Get("student_not_found") != "",
        DeleteStudent:     r.URL.Query().Get("deleted_student_id") != "",
        DeleteStudentID:   r.URL.Query().Get("deleted_student_id"),
    }

    tmpl.Execute(w, data)
}
