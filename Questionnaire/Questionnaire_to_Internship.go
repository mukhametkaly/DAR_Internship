package Questionnaire

import (
	"Internship/internship"
    "errors"
)


type QuestionnaireInInternship interface {
	AddQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	GetQuestionnaires() ([]*Questionnaire,error)
	GetQuestionnaire(id int64) (*Questionnaire, error)
	UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error)
	DeleteQuestionnaire(questionnaire *Questionnaire) error
	GetQuestionnaireFromInternship (id int64) (*Questionnaire, error)
}

type QuestionnaireInInternshipClass struct {
	Quest QuestionnaireCollection
	internship Internship.InternshipCollection
}

func NewQuestionnaireInInternship(intcollection Internship.InternshipCollection, quest QuestionnaireCollection )  QuestionnaireInInternship {
	return &QuestionnaireInInternshipClass{Quest:quest, internship: intcollection}
}

func(QuestIntrnshp *QuestionnaireInInternshipClass) CheckQuestionnaire (questionnaire *Questionnaire)  (*Questionnaire, error)  {

	if questionnaire.Questions[0] == "" {
		return nil, errors.New("No Questions ")
	}
	if questionnaire.InternshipID == 0 {
		return nil, errors.New("No Internship ID ")
	}
	if questionnaire.StartTime == "" {
		return nil, errors.New("No start time ")
	}
	if questionnaire.EndTime == "" {
		return nil, errors.New("No end time ")
	}
	_, err:=QuestIntrnshp.internship.GetInternship(questionnaire.InternshipID)
	if err != nil {
		return nil, err
	}
	return questionnaire, nil
}

func (QuestIntrnshp *QuestionnaireInInternshipClass) GetQuestionnaires() ([]*Questionnaire,error) {
	questionnaires, err:=QuestIntrnshp.Quest.GetQuestionnaires()
	if err != nil {
		return nil, err
	}
	return questionnaires, err
}

func (QuestIntrnshp *QuestionnaireInInternshipClass) GetQuestionnaire(id int64)  (*Questionnaire, error)  {

	questionnaire, err := QuestIntrnshp.Quest.GetQuestionnaire(id)
	if err!= nil {
		return nil, err
	}
	return questionnaire, nil
}

func (QuestIntrnshp *QuestionnaireInInternshipClass)AddQuestionnaire (questionnaire *Questionnaire)    (*Questionnaire, error) {
	_, err := QuestIntrnshp.CheckQuestionnaire(questionnaire)
	if err != nil {
		return nil, err
	}
	_, err = QuestIntrnshp.Quest.AddQuestionnaire(questionnaire)
	if err != nil {
		return nil, err
	}
	return questionnaire, nil
}

func (QuestIntrnshp *QuestionnaireInInternshipClass) UpdateQuestionnaire(questionnaire *Questionnaire) (*Questionnaire, error) {
	_, err := QuestIntrnshp.CheckQuestionnaire(questionnaire)
	if err != nil {
		return nil, err
	}
	_, err = QuestIntrnshp.Quest.UpdateQuestionnaire(questionnaire)
	if err != nil {
		return nil, err
	}
	return questionnaire, nil

}
func (QuestIntrnshp *QuestionnaireInInternshipClass) DeleteQuestionnaire(questionnaire *Questionnaire)  error {
	if questionnaire.QuestionnaireID == 0 {
		return errors.New("NO ID")
	}
	err := QuestIntrnshp.Quest.DeleteQuestionnaire(questionnaire)
	if err != nil {
		return err
	}
	return err

}
func (QuestIntrnshp *QuestionnaireInInternshipClass) GetQuestionnaireFromInternship (id int64)  (*Questionnaire, error) {
	_, err := QuestIntrnshp.internship.GetInternship(id)
	if err != nil {
		return nil,err
	}
	questionnaires, err:= QuestIntrnshp.Quest.GetQuestionnaireFromInternship(id)
	if err != nil {
		return nil, err
	}
	return questionnaires, nil

}










