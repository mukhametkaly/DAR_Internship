package Questionnaire

type QuestionnaireCollection interface {
	AddQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	GetQuestionnaires() ([]*Questionnaire,error)
	GetQuestionnaire(id int64) (*Questionnaire, error)
	UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	DeleteQuestionnaire(questionnaire *Questionnaire) error
	GetQuestionnaireFromInternship (id int64)  (*Questionnaire, error)
}

type Questionnaire struct {

	QuestionnaireID int `json:"questionnaire_id"`
	InternshipID int64 `json:"internship_id"`
	Questions []string `json:"questions"`
	StartTime string `json:"starttime"`
	EndTime string `json:"endtime"`
}

