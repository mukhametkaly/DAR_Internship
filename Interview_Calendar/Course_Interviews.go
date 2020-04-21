package Interview_Calendar

import (
	"errors"
	"github.com/mukhametkaly/DAR_Internship/Courses"
)


type CourseInterviewsCal interface {
	CheckInterviewCal(intrvw *InterviewCalendar)    (*InterviewCalendar, error)
	AddInterviewCal (intrvw *InterviewCalendar)    (*InterviewCalendar, error)
	GetInterviewCals()                  ([]*InterviewCalendar,error)
	GetInterviewCal(id int64)           (*InterviewCalendar, error)
	UpdateInterviewCal(intrvw *InterviewCalendar) (*InterviewCalendar, error)
	DeleteInterviewCal(intrvw *InterviewCalendar)            error
	GetInternviewCalendarFromCourses (id int64)  ([]*InterviewCalendar, error)
}

type CourseInterviewsCalClass struct {
	intervw InterviewCalendarCollection
	courses Courses.CourseCollection
}

func NewCourseIntern(courscollection Courses.CourseCollection, intrvwCal InterviewCalendarCollection )  CourseInterviewsCal {
	return &CourseInterviewsCalClass{intervw:intrvwCal, courses: courscollection}
}

func(Intrnvw *CourseInterviewsCalClass) CheckInterviewCal (interview *InterviewCalendar)  (*InterviewCalendar, error)  {

	if interview.ComeTime == "" {
		return nil, errors.New("No come time ")
	}
	if interview.ComeDate == "" {
		return nil, errors.New("No come date ")
	}
	if interview.LecturerMail == "" {
		return nil, errors.New("No Lecturer mail")
	}
	if interview.Duration == "" {
		return nil, errors.New("No Duration")
	}
	if interview.InternMail == "" {
		return nil, errors.New("No Intern mail")
	}
	if interview.CourseID == 0 {
		return nil, errors.New("No Course ID")
	}
	_, err:=Intrnvw.courses.GetCourse(interview.CourseID)
	if err != nil {
		return nil, err
	}
	return interview, nil
}

func (Intrnvw *CourseInterviewsCalClass) GetInterviewCals() ([]*InterviewCalendar,error) {
	intern, err:=Intrnvw.intervw.GetInterviewCalendars()
	if err != nil {
		return nil, err
	}
	return intern, err
}

func (Intrnvw *CourseInterviewsCalClass) GetInterviewCal(id int64)  (*InterviewCalendar, error)  {

	intern, err := Intrnvw.intervw.GetInterviewCalendar(id)
	if err!= nil {
		return nil, err
	}
	return intern, nil
}

func (Intrnvw *CourseInterviewsCalClass)AddInterviewCal (interview *InterviewCalendar)    (*InterviewCalendar, error) {
	_, err := Intrnvw.CheckInterviewCal(interview)
	if err != nil {
		return nil, err
	}
	_, err = Intrnvw.intervw.AddInterviewCalendar(interview)
	if err != nil {
		return nil, err
	}
	return interview, nil
}

func (Intrnvw *CourseInterviewsCalClass) UpdateInterviewCal (interview *InterviewCalendar) (*InterviewCalendar, error) {
	_, err := Intrnvw.CheckInterviewCal(interview)
	if err != nil {
		return nil, err
	}
	_, err = Intrnvw.intervw.UpdateInterviewCalendar(interview)
	if err != nil {
		return nil, err
	}
	return interview, nil

}
func (Intrnvw *CourseInterviewsCalClass) DeleteInterviewCal (interview *InterviewCalendar)  error {
	if interview.InterviewCalendarID == 0 {
		return errors.New("NO ID")
	}
	err := Intrnvw.intervw.DeleteInterviewCalendar(interview)
	if err != nil {
		return err
	}
	return err

}
func (Intrnvw *CourseInterviewsCalClass) GetInternviewCalendarFromCourses (id int64)  ([]*InterviewCalendar, error) {
	_, err := Intrnvw.courses.GetCourse(id)
	if err != nil {
		return nil,err
	}
	interviews, err:= Intrnvw.intervw.GetInternviewCalendarFromCourses(id)
	if err != nil {
		return nil, err
	}
	return interviews, nil

}





