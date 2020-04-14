package Lecturer

import (
	"Internship/Courses"
	"errors"
	"Internship/Veryfing"

)


type CourseLecturer interface {
	CheckCourseLecturer(lecturer *Lecturer)    (*Lecturer, error)
	AddCourseLecturer (lecturer *Lecturer)    (*Lecturer, error)
	GetCourseLecturers()                  ([]*Lecturer,error)
	GetCourseLecturer(id int64)           (*Lecturer, error)
	UpdateCourseLecturer(lecturer *Lecturer) (*Lecturer, error)
	DeleteCourseLecturer(lecturer *Lecturer)            error
	GetLecturerFromCourses (id int64)  (*Lecturer, error)
	Authorization (username string, password string) error

}

type CourseLecturerClass struct {
	lecturers Lecturers
	courses Courses.CourseCollection
}

func NewCourseLecturer(courscollection Courses.CourseCollection, lectrs Lecturers )  CourseLecturer {
	return &CourseLecturerClass{lecturers:lectrs, courses: courscollection}
}

func(CoursLec *CourseLecturerClass) CheckCourseLecturer (lecturer *Lecturer)  (*Lecturer, error)  {

	if lecturer.LecturerName == "" {
		return nil, errors.New("No Lecturer Name ")
	}
	if lecturer.Mail == "" {
		return nil, errors.New("No Lecturer mail")
	}
	if lecturer.Password == "" {
		return nil, errors.New("No Password")
	}
	errs := Veryfing.VerifyPassword(lecturer.Password)
	if errs != nil {
		return nil, errs
	}
	if lecturer.CourseID == 0 {
		return nil, errors.New("No Course ID")
	}
	_, err:=CoursLec.courses.GetCourse(lecturer.CourseID)
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
	_, err = CoursLec.lecturers.GetLecturerByUsername(lecturer.UserName)
	if err == nil {
		return nil, errors.New("Lecturer wiht this username are exists")
	}
	err = CoursLec.SetLecturerInCourse(lecturer.CourseID, lecturer)
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
	if lecturer.LecturerID == 0 {
		return errors.New("NO ID")
	}
	err := CoursLec.lecturers.DeleteLecturer(lecturer)
	if err != nil {
		return err
	}
	return err

}
func (CoursIntrn *CourseLecturerClass) GetLecturerFromCourses (id int64)  (*Lecturer, error) {
	_, err := CoursIntrn.courses.GetCourse(id)
	if err != nil {
		return nil,err
	}
	lecturer, err:= CoursIntrn.lecturers.GetLecturerFromCourses(id)
	if err != nil {
		return nil, err
	}
	return lecturer, nil
}

func (CoursIntrn *CourseLecturerClass) SetLecturerInCourse (id int64, lecturer *Lecturer) error{
	course, err := CoursIntrn.courses.GetCourse(id)
	if err != nil {
		return err
	}
	course.LecturerMail = lecturer.Mail
	course.LecturerName = lecturer.LecturerName
	course.LecturerID = lecturer.LecturerID
	_, err = CoursIntrn.courses.UpdateCourse(course)
	return err
}

func (CoursLec *CourseLecturerClass) Authorization (username string, password string) error {
	if username == "" {
		return errors.New("No Username")
	}
	if password == "" {
		return errors.New("No Password")
	}

	err := CoursLec.lecturers.Authorization(username, password)
	if err!= nil {
		return err
	}
	return nil
}

