package Lecturer

import "github.com/go-redis/redis"

type Lecturers interface {
	AddLecturer(lecturer *Lecturer) (*Lecturer, error)
	GetLecturers() ([]*Lecturer,error)
	GetLecturer(id int64) (*Lecturer, error)
	UpdateLecturer(lecturer *Lecturer) (*Lecturer, error)
	DeleteLecturer(lecturer *Lecturer) error
	GetLecturerFromCourses (id int64)  (*Lecturer, error)
	GetLecturerByUsername (username string) (*Lecturer, error)
	Authorization (username string, password string, client *redis.Client) error
}

type Lecturer struct {
	UserName      string `json:"user_name"`
	LecturerID    int64  `json:"lecturer_id,omitempty"`
	LecturerName string `json:"name,omitempty"`
	Mail         string `json:"mail,omitempty"`
	CourseID     int64  `json:"course_id,omitempty"`
	Password     string `json:"passwd, omitempty"`
}

