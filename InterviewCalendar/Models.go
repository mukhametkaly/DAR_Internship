package InterviewCalendar

type InterviewCalendarCollection interface {
	AddInterviewCalendar(interviewCalendar *InterviewCalendar) (*InterviewCalendar, error)
	GetInterviewCalendars() ([]*InterviewCalendar,error)
	GetInterviewCalendar(id int64) (*InterviewCalendar, error)
	UpdateInterviewCalendar(interviewCalendar *InterviewCalendar) (*InterviewCalendar, error)
	DeleteInterviewCalendar(interviewCalendar *InterviewCalendar) error
}

type InterviewCalendar struct {
	InterviewCalendarID int64 `json:"interviewcalendar_id"`
	comeDate string `json:"come_date"`
	comeTime string `json:"cometime"`
	LecturerMail string `json:"lecturer_mail"`
	Duration string `json:"duration"`
	InternMail string `json:"intern_mail"`
	InternshipID int `json:"internship_id"`
	CourseID int64 `json:"course_id"`

}
