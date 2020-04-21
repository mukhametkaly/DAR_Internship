package Intern

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/mukhametkaly/DAR_Internship/Courses"
)



type CourseIntern interface {
	CheckIntern(intern *Intern)    (*Intern, error)
	AddIntern (intern *Intern)    (*Intern, error)
	GetInterns()                  ([]*Intern,error)
	GetIntern(id int64)           (*Intern, error)
	UpdateIntern(intern *Intern) (*Intern, error)
	DeleteIntern(intern *Intern)            error
	GetInternsFromCourses (id int64)  ([]*Intern, error)
	Authorization(username string, password string, client *redis.Client) error
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
	if intern.CourseID == 0 {
		return nil, errors.New("No Course ID ")
	}
	if intern.Answers[0] == "" {
		return nil, errors.New("No Answers")
	}
	if intern.Mail == "" {
		return nil, errors.New("No mail")
	}
	if intern.ContestUsername == "" {
		return nil, errors.New("No Contest username")
	}
	if intern.QuestionnaireID == 0 {
	return nil, errors.New("No Questionnaire ID")
	}
	_, err:=CoursIntrn.courses.GetCourse(intern.CourseID)
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
	_, err = CoursIntrn.intern.GetInternByUsername(intern.UserName)
	if err == nil {
		return nil, errors.New("Intern wiht this username are exists")
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
func (CoursIntrn *CourseInternClass) GetInternsFromCourses (id int64)  ([]*Intern, error) {
	_, err := CoursIntrn.courses.GetCourse(id)
	if err != nil {
		return nil,err
	}
	interns, err:= CoursIntrn.intern.GetInternsFromCourses(id)
	if err != nil {
		return nil, err
	}
	return interns, nil

}

func (CoursIntern *CourseInternClass) Authorization(username string, password string, client *redis.Client) error  {
    if username == "" {
    	return errors.New("No username")
	}
	if password == "" {
		return errors.New("No password")
	}
	err := CoursIntern.intern.Authorization(username, password, client)
	return err

}









