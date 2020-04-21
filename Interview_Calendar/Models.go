package Interview_Calendar

type InterviewCalendarCollection interface {
	AddInterviewCalendar(interviewCalendar *InterviewCalendar) (*InterviewCalendar, error)
	GetInterviewCalendars() ([]*InterviewCalendar,error)
	GetInterviewCalendar(id int64) (*InterviewCalendar, error)
	UpdateInterviewCalendar(interviewCalendar *InterviewCalendar) (*InterviewCalendar, error)
	DeleteInterviewCalendar(interviewCalendar *InterviewCalendar) error
	GetInternviewCalendarFromCourses (id int64)  ([]*InterviewCalendar, error)
}

type InterviewCalendar struct {
	InterviewCalendarID int64  `json:"interviewcalendar_id"`
	ComeDate            string `json:"come_date"`
	ComeTime            string `json:"come_time"`
	LecturerID          int64   `json:"lecturer_id"`
	LecturerMail        string `json:"lecturer_mail"`
	Duration            string `json:"duration"`
	InternID            int64  `json:"intern_id"`
	InternMail          string `json:"intern_mail"`
	CourseID            int64  `json:"course_id"`

}
