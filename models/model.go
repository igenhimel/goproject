package models

type Student struct {
    ID            string  `json:"student_id"`
    Name          string  `json:"student_name"`
    CGPA          float64 `json:"student_cgpa"`
    CareerInterest string  `json:"career_interest"`
    ImageURL      string  `json:"image_url"`
}

var (
    Students = make(map[string]Student) // Exported map variable
)
