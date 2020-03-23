package Questionnaire

type QuestionnairePerformer interface {
	AddQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	GetQuestionnaires() ([]*Questionnaire,error)
	GetQuestionnaire(id int64) (*Questionnaire, error)
	UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	DeleteQuestionnaire(questionnaire *Questionnaire) error
}

type questionnairePerformer struct {
	collection QuestionnaireCollection
	// import InternshipCollection
	internshipCol InternshipCollection
}

func NewQuestionnairePerformer(collection QuestionnaireCollection) QuestionnairePerformer {
	return &questionnairePerformer{collection:collection}
}

func (s *questionnairePerformer) AddQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error) {
	//Here should be validation of fields and other checks
	// If you need some methods from another interface (e.g. for InternshipCollection) use
	// s.internshipCol.GetInternship(questionnaire.InternshipID)
	return s.collection.AddQuestionnaire(questionnaire)
}
func (s *questionnairePerformer) GetQuestionnaires() ([]*Questionnaire,error) {
	return s.collection.GetQuestionnaires()
}
func (s *questionnairePerformer) GetQuestionnaire(id int64) (*Questionnaire, error) {
	return s.collection.GetQuestionnaire(id)
}
func (s *questionnairePerformer) UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error) {
	return s.collection.UpdateQuestionnaire(questionnaire)
}
func (s *questionnairePerformer) DeleteQuestionnaire(questionnaire *Questionnaire) error {
	return s.collection.DeleteQuestionnaire(questionnaire)
}