package Lecturer

import (
"../Courses"
"errors"
)


type CourseLecturer interface {
	CheckCourseLecturer(lecturer *Lecturer)    (*Lecturer, error)
	AddCourseLecturer (lecturer *Lecturer)    (*Lecturer, error)
	GetCourseLecturers()                  ([]*Lecturer,error)
	GetCourseLecturer(id int64)           (*Lecturer, error)
	UpdateCourseLecturer(lecturer *Lecturer) (*Lecturer, error)
	DeleteCourseLecturer(lecturer *Lecturer)            error
}

type CourseLecturerClass struct {
	lecturers Lecturers
	courses Courses.CourseCollection
}

func NewCourseIntern(courscollection Courses.CourseCollection, lectrs Lecturers )  CourseLecturer {
	return &CourseLecturerClass{lecturers:lectrs, courses: courscollection}
}

func(CoursLec *CourseLecturerClass) CheckCourseLecturer (lecturer *Lecturer)  (*Lecturer, error)  {

	if lecturer.LecturerName == "" {
		return nil, errors.New("No Lecturer Name ")
	}
	if lecturer.mail == "" {
		return nil, errors.New("No Lecturer mail")
	}
	if lecturer.password == "" {
		return nil, errors.New("No Password")
	}
	if lecturer.courseID == 0 {
		return nil, errors.New("No Course ID")
	}
	_, err:=CoursLec.courses.GetCourse(lecturer.courseID)
	if err != nil {
		return nil, err
	}
	return lecturer, nil
}

func (CoursLec *CourseLecturerClass) GetCourseLecturers() ([]*Lecturer,error) {
	lecturer, err:=CoursLec.lecturers.GetLecturers()
	if err != nil {
		return nil, err
	}
	return lecturer, err
}

func (CoursLec *CourseLecturerClass) GetCourseLecturer(id int64)  (*Lecturer, error)  {

	lecturer, err := CoursLec.lecturers.GetLecturer(id)
	if err!= nil {
		return nil, err
	}
	return lecturer, nil
}

func (CoursLec *CourseLecturerClass)AddCourseLecturer (lecturer *Lecturer)    (*Lecturer, error) {
	_, err := CoursLec.CheckCourseLecturer(lecturer)
	if err != nil {
		return nil, err
	}
	_, err = CoursLec.lecturers.AddLecturer(lecturer)
	if err != nil {
		return nil, err
	}
	return lecturer, nil
}

func (CoursLec *CourseLecturerClass) UpdateCourseLecturer (lecturer *Lecturer) (*Lecturer, error) {
	_, err := CoursLec.CheckCourseLecturer(lecturer)
	if err != nil {
		return nil, err
	}
	_, err = CoursLec.lecturers.UpdateLecturer(lecturer)
	if err != nil {
		return nil, err
	}
	return lecturer, nil

}
func (CoursLec *CourseLecturerClass) DeleteCourseLecturer (lecturer *Lecturer)  error {
	if lecturer.LectureID == 0 {
		return errors.New("NO ID")
	}
	err := CoursLec.lecturers.DeleteLecturer(lecturer)
	if err != nil {
		return err
	}
	return err

}
