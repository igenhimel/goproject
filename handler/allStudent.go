package handler

import (
    "html/template"
    "net/http"
    "sort"
    "strconv"
    "goproject/models"
)

const studentsPerPage = 3 // Number of students per page

// CalculatePageNumbers calculates the list of page numbers to display
func CalculatePageNumbers(currentPage, totalPages int) []int {
    var pageNumbers []int

    for i := 1; i <= totalPages; i++ {
        pageNumbers = append(pageNumbers, i)
    }

    return pageNumbers
}

// ShowAllStudentsHandler handles displaying all students with pagination and sorting by CGPA.
func ShowAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        pageStr := r.URL.Query().Get("page")
        page, err := strconv.Atoi(pageStr)
        if err != nil || page <= 0 {
            page = 1
        }

        var students []models.Student
        for _, student := range models.Students {
            students = append(students, student)
        }

        // Create a channel to communicate the sorted students
        sortedChan := make(chan []models.Student)

        // Goroutine to sort students by CGPA in descending order
        go func() {
            sort.SliceStable(students, func(i, j int) bool {
                return students[i].CGPA > students[j].CGPA
            })
            sortedChan <- students
        }()

        sortedStudents := <-sortedChan

        startIndex := (page - 1) * studentsPerPage
        endIndex := startIndex + studentsPerPage

        if endIndex > len(sortedStudents) {
            endIndex = len(sortedStudents)
        }

        // Get students for the current page
        pagedStudents := sortedStudents[startIndex:endIndex]

        totalPages := (len(sortedStudents) + studentsPerPage - 1) / studentsPerPage

        data := struct {
            Students       []models.Student
            CurrentPage    int
            TotalPages     int
            HasNextPage    bool
            HasPrevPage    bool
            PrevPageNumber int
            NextPageNumber int
            PageNumbers    []int // Added PageNumbers field
        }{
            Students:       pagedStudents,
            CurrentPage:    page,
            TotalPages:     totalPages,
            HasNextPage:    endIndex < len(sortedStudents),
            HasPrevPage:    startIndex > 0,
            PrevPageNumber: page - 1,
            NextPageNumber: page + 1,
            PageNumbers:    CalculatePageNumbers(page, totalPages), // Calculate Page Numbers
        }

        tmpl := template.Must(template.ParseFiles("templates/show_students.html"))
        tmpl.Execute(w, data)
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
