package Courses

import (
	"Internship/internship"
	 "errors"
)



type CoursesInInternship interface {
	CheckCourse(course *Courses)    (*Courses, error)
	AddCourse (course *Courses)    (*Courses, error)
	GetCourses()                  ([]*Courses,error)
	GetCourse(id int64)           (*Courses, error)
	UpdateCourse(course *Courses) (*Courses, error)
	DeleteCourse(course *Courses)            error
	GetCoursesFromInternship (id int64)  ([]*Courses, error)
}

type CoursesInInternshipClass struct {
	cours CourseCollection
	internship Internship.InternshipCollection
}

func NewCoursesInInternship(intcollection Internship.InternshipCollection, coursecollection CourseCollection )  CoursesInInternship {
	return &CoursesInInternshipClass{cours:coursecollection, internship: intcollection}
}

func(CrsIntrnship *CoursesInInternshipClass) CheckCourse (course *Courses)  (*Courses, error)  {

	if course.Title == "" {
		return nil, errors.New("No Title ")
	}
	if course.InternshipID == 0 {
		return nil, errors.New("No Internship ID ")
	}
	_, err:=CrsIntrnship.internship.GetInternship(course.InternshipID)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (CrsIntrnship *CoursesInInternshipClass) GetCourses() ([]*Courses,error) {
	courses, err:=CrsIntrnship.cours.GetCourses()
	if err != nil {
		return nil, err
	}
	return courses, err
}

func (CrsIntrnship *CoursesInInternshipClass) GetCourse(id int64)  (*Courses, error)  {

	course, err := CrsIntrnship.cours.GetCourse(id)
	if err!= nil {
		return nil, err
	}
	return course, nil
}

func (CrsIntrnship *CoursesInInternshipClass)AddCourse (course *Courses)    (*Courses, error) {
	_, err := CrsIntrnship.CheckCourse(course)
	if err != nil {
		return nil, err
	}
	_, err = CrsIntrnship.cours.AddCourse(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (CrsIntrnship *CoursesInInternshipClass) UpdateCourse(course *Courses) (*Courses, error) {
	_, err := CrsIntrnship.CheckCourse(course)
	if err != nil {
		return nil, err
	}
	_, err = CrsIntrnship.cours.UpdateCourse(course)
	if err != nil {
		return nil, err
	}
	return course, nil

}
func (CrsIntrnship *CoursesInInternshipClass) DeleteCourse(course *Courses)  error {
	if course.CourseID == 0 {
		return errors.New("NO ID")
	}
	err := CrsIntrnship.cours.DeleteCourse(course)
	if err != nil {
		return err
	}
	return err

}
func (CrsIntrnship *CoursesInInternshipClass) GetCoursesFromInternship (id int64)  ([]*Courses, error) {
	_, err := CrsIntrnship.internship.GetInternship(id)
	if err != nil {
		return nil,err
	}
	courses, err:= CrsIntrnship.cours.GetCoursesFromInternship(id)
	if err != nil {
		return nil, err
	}
	return courses, nil

}









