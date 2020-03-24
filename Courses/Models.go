package Courses

type CourseCollection interface {
	AddCourse(course *Courses)    (*Courses, error)
	GetCourses()                  ([]*Courses,error)
	GetCourse(id int64)           (*Courses, error)
	UpdateCourse(course *Courses) (*Courses, error)
	DeleteCourse(course *Courses)            error
}

type Courses struct {
	CourseID     int64  `json:"coures_id"`
	Title        string `json:"title,omitempty"`
	LecturerID   int64 `json:"lecturerid,omitempty"`
	LecturerName string `json:"lecturer_name, omitempty"`
	LecturerMail string `json:"lecturer_mail,omitempty"`
	InternshipID int64  `json:"internship_id"`
}

