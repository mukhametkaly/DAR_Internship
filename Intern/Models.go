package Intern

type InternCollection interface {
	AddIntern(intern *Intern) (*Intern, error)
	GetInterns() ([]*Intern,error)
	GetIntern(id int64) (*Intern, error)
	UpdateIntern(intern *Intern) (*Intern, error)
	DeleteIntern(intern *Intern) error
}

type Intern struct {
	InternID        int64  `json:"intern_id"`
	Name            string `json:"name"`
	Mail            string `json:"mail"`
	answers       []string `json:"answers"`
	contestID       int64  `json:"contest_id"`
	courseID        int64  `json:"course_id"`
	status          string `json:"status"`
	contest_score   string `json:"contest_score"`
	contestUsername string `json:"contest_username"`


}

