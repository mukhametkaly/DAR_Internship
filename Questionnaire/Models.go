package Questionnaire

type QuestionnaireCollection interface {
	AddQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	GetQuestionnaires() ([]*Questionnaire,error)
	GetQuestionnaire(id int64) (*Questionnaire, error)
	UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	DeleteQuestionnaire(questionnaire *Questionnaire) error
}

type Questionnaire struct {

	QuestionnaireID int `json:"questionnaire_id"`
	InternshipID int `json:"internship_id"`
	Questions []string `json:"questions,omitempty"`
	StartTime string `json:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty"`
}

