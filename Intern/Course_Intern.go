package Intern

import (
"../Courses"
"errors"
)



type CourseIntern interface {
	CheckIntern(intern *Intern)    (*Intern, error)
	AddIntern (intern *Intern)    (*Intern, error)
	GetInterns()                  ([]*Intern,error)
	GetIntern(id int64)           (*Intern, error)
	UpdateIntern(intern *Intern) (*Intern, error)
	DeleteIntern(intern *Intern)            error
}

type CourseInternClass struct {
	intern InternCollection
	courses Courses.CourseCollection
}

func NewCourseIntern(courscollection Courses.CourseCollection, interncollection InternCollection )  CourseIntern {
	return &CourseInternClass{intern:interncollection, courses: courscollection}
}

func(CoursIntrn *CourseInternClass) CheckIntern (intern *Intern)  (*Intern, error)  {

	if intern.Name == "" {
		return nil, errors.New("No Name ")
	}
	if intern.courseID == 0 {
		return nil, errors.New("No Course ID ")
	}
	if intern.answers[0] == "" {
		return nil, errors.New("No Answers")
	}
	if intern.Mail == "" {
		return nil, errors.New("No mail")
	}
	if intern.contestUsername == "" {
		return nil, errors.New("No Contest username")
	}
	_, err:=CoursIntrn.courses.GetCourse(intern.courseID)
	if err != nil {
		return nil, err
	}
	return intern, nil
}

func (CoursIntrn *CourseInternClass) GetInterns() ([]*Intern,error) {
	intern, err:=CoursIntrn.intern.GetInterns()
	if err != nil {
		return nil, err
	}
	return intern, err
}

func (CoursIntrn *CourseInternClass) GetIntern(id int64)  (*Intern, error)  {

	intern, err := CoursIntrn.intern.GetIntern(id)
	if err!= nil {
		return nil, err
	}
	return intern, nil
}

func (CoursIntrn *CourseInternClass)AddIntern (intern *Intern)    (*Intern, error) {
	_, err := CoursIntrn.CheckIntern(intern)
	if err != nil {
		return nil, err
	}
	_, err = CoursIntrn.intern.AddIntern(intern)
	if err != nil {
		return nil, err
	}
	return intern, nil
}

func (CoursIntrn *CourseInternClass) UpdateIntern(intern *Intern) (*Intern, error) {
	_, err := CoursIntrn.CheckIntern(intern)
	if err != nil {
		return nil, err
	}
	_, err = CoursIntrn.intern.UpdateIntern(intern)
	if err != nil {
		return nil, err
	}
	return intern, nil

}
func (CoursIntrn *CourseInternClass) DeleteIntern(intern *Intern)  error {
	if intern.InternID == 0 {
		return errors.New("NO ID")
	}
	err := CoursIntrn.intern.DeleteIntern(intern)
	if err != nil {
		return err
	}
	return err

}










