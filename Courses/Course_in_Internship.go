package Courses

import (
	"../internship"
	 "errors"
)



type Courses_in_Internship interface {
	CheckCourse(course *Courses)    (*Courses, error)
	AddCourse (course *Courses)    (*Courses, error)
	GetCourses()                  ([]*Courses,error)
	GetCourse(id int64)           (*Courses, error)
	UpdateCourse(course *Courses) (*Courses, error)
	DeleteCourse(course *Courses)            error
}

type courses_in_Internship struct {
	cours CourseCollection
	internship Internship.InternshipCollection
}

func New_Cours_in_Internship_Collection(intcollection Internship.InternshipCollection, coursecollection CourseCollection )  Courses_in_Internship {
	return &courses_in_Internship{cours:coursecollection, internship: intcollection}
}

func(crs_intrnshp *courses_in_Internship) CheckCourse (course *Courses)  (*Courses, error)  {

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

func (crs_intrnshp *courses_in_Internship) GetCourses() ([]*Courses,error) {
	courses, err:=crs_intrnshp.cours.GetCourses()
	if err != nil {
		return nil, err
	}
	return courses, err
}

func (crs_intrnshp *courses_in_Internship) GetCourse(id int64)  (*Courses, error)  {

	course, err := crs_intrnshp.cours.GetCourse(id)
	if err!= nil {
		return nil, err
	}
	return course, nil
}

func (crs_intrnshp *courses_in_Internship)AddCourse (course *Courses)    (*Courses, error) {
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

func (crs_intrnshp *courses_in_Internship) UpdateCourse(course *Courses) (*Courses, error) {
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
func (crs_intrnshp *courses_in_Internship) DeleteCourse(course *Courses)  error {
	if course.CourseID == 0 {
		return errors.New("NO ID")
	}
	err := crs_intrnshp.cours.DeleteCourse(course)
	if err != nil {
		return err
	}
	return err

}









