package Intern

type InternCollection interface {
	AddIntern(intern *Intern) (*Intern, error)
	GetInterns() ([]*Intern,error)
	GetIntern(id int64) (*Intern, error)
	UpdateIntern(intern *Intern) (*Intern, error)
	DeleteIntern(intern *Intern) error
	GetInternsFromCourses (id int64)  ([]*Intern, error)
}

type Intern struct {
	InternID        int64    `json:"intern_id"`
	Name            string   `json:"name"`
	Mail            string   `json:"mail"`
	QuestionnaireID int64     `json:"questionnaire_id"`
	Answers         []string `json:"answers"`
	ContestID       int64    `json:"contest_id"`
	CourseID        int64    `json:"course_id"`
	Status          string   `json:"status"`
	ContestScore   string   `json:"contest_score"`
	ContestUsername string   `json:"contest_username"`


}

