package Lecturer

type Lecturers interface {
	AddLecturer(lecturer *Lecturer) (*Lecturer, error)
	GetLecturers() ([]*Lecturer,error)
	GetLecturer(id int64) (*Lecturer, error)
	UpdateLecturer(lecturer *Lecturer) (*Lecturer, error)
	DeleteLecturer(lecturer *Lecturer) error
}

type Lecturer struct {
	LectureID int64 `json:"lecturer_id,omitempty"`
	LecturerName string `json:"name,omitempty"`
	mail string `json:"mail,omitempty"`
	courseID int64 `json:"course_id,omitempty"`
	password string `json:"passwd, omitempty"`
}

