package Courses

import (
	"../internship"
	 "errors"
)



type CoursesInInternship interface {
	CheckCourse(course *Courses)    (*Courses, error)
	AddCourse (course *Courses)    (*Courses, error)
	GetCourses()                  ([]*Courses,error)
	GetCourse(id int64)           (*Courses, error)
	UpdateCourse(course *Courses) (*Courses, error)
	DeleteCourse(course *Courses)            error
}

type CoursesInInternshipClass struct {
	cours CourseCollection
	internship Internship.InternshipCollection
}

func NewCoursesInInternship(intcollection Internship.InternshipCollection, coursecollection CourseCollection )  CoursesInInternship {
	return &CoursesInInternshipClass{cours:coursecollection, internship: intcollection}
}

func(crs_intrnshp *CoursesInInternshipClass) CheckCourse (course *Courses)  (*Courses, error)  {

	if course.Title == "" {
		return nil, errors.New("No Title ")
	}
	if course.InternshipID == 0 {
		return nil, errors.New("No Internship ID ")
	}
	_, err:=crs_intrnshp.internship.GetInternship(course.InternshipID)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (crs_intrnshp *CoursesInInternshipClass) GetCourses() ([]*Courses,error) {
	courses, err:=crs_intrnshp.cours.GetCourses()
	if err != nil {
		return nil, err
	}
	return courses, err
}

func (crs_intrnshp *CoursesInInternshipClass) GetCourse(id int64)  (*Courses, error)  {

	course, err := crs_intrnshp.cours.GetCourse(id)
	if err!= nil {
		return nil, err
	}
	return course, nil
}

func (crs_intrnshp *CoursesInInternshipClass)AddCourse (course *Courses)    (*Courses, error) {
	_, err := crs_intrnshp.CheckCourse(course)
	if err != nil {
		return nil, err
	}
	_, err = crs_intrnshp.cours.AddCourse(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (crs_intrnshp *CoursesInInternshipClass) UpdateCourse(course *Courses) (*Courses, error) {
	_, err := crs_intrnshp.CheckCourse(course)
	if err != nil {
		return nil, err
	}
	_, err = crs_intrnshp.cours.UpdateCourse(course)
	if err != nil {
		return nil, err
	}
	return course, nil

}
func (crs_intrnshp *CoursesInInternshipClass) DeleteCourse(course *Courses)  error {
	if course.CourseID == 0 {
		return errors.New("NO ID")
	}
	err := crs_intrnshp.cours.DeleteCourse(course)
	if err != nil {
		return err
	}
	return err

}









