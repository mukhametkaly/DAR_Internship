package Intern

type InternCollection interface {
	AddIntern(intern *Intern) (*Intern, error)
	GetInterns() ([]*Intern,error)
	GetIntern(id int64) (*Intern, error)
	UpdateIntern(intern *Intern) (*Intern, error)
	DeleteIntern(intern *Intern) error
}

type Intern struct {
	InternID int64 `json:"intern_id"`
	Name string `json:"name, omitempty"`
	LecturerName string `json:"lecture_name, omitempty"`
	contestID int64 `json:"contest_id, omitempty"`
	courseID int64 `json:"course_id, omitempty"`
	status string `json:"status, omitempty"`
	contest_score string `json:"contest_score, omitempty"`
	contestUsername string `json:"contest_username,omitempty"`


}

