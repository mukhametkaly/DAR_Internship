package Questionnaire

import (
	"Internship/internship"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	collection *mongo.Collection
)

type QuestionnaireCollectionClass struct{
	dbcon *mongo.Database
}

func NewQuestionnaireCollection(config Internship.MongoConfig) (QuestionnaireCollection, error){

	clientOptions:=options.Client().ApplyURI("mongodb://"+config.Host+":"+config.Port)
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		return nil,err
	}
	err = client.Ping(context.TODO(),nil)
	if err!=nil{
		return nil,err
	}

	db:=client.Database(config.Database)
	collection=db.Collection("Questionnaire")
	return &QuestionnaireCollectionClass{dbcon:db,}, nil
}




func(incc *QuestionnaireCollectionClass) GetQuestionnaires() ([]*Questionnaire, error){

	findOptions:=options.Find()
	var questionnaires []*Questionnaire
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var questionnaire Questionnaire
		err:=cur.Decode(&questionnaire)
		if err!=nil{
			return nil,err
		}
		questionnaires = append(questionnaires,&questionnaire)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return questionnaires,nil
}



func (incc *QuestionnaireCollectionClass) AddQuestionnaire(quest *Questionnaire) (*Questionnaire,error){

	questionnaires,err:=incc.GetQuestionnaires()
	n:=len(questionnaires)
	if n!=0{
		questionnaire:=questionnaires[n-1]
		quest.QuestionnaireID = questionnaire.QuestionnaireID+1
	}else{
		quest.QuestionnaireID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), quest)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return quest,nil

}

func (incc *QuestionnaireCollectionClass) GetQuestionnaire(id int64) (*Questionnaire,error){

	filter:=bson.D{{"questionnaireid",id}}
	questionnnaire:=&Questionnaire{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&questionnnaire)
	if err!=nil{
		return nil,err
	}
	return questionnnaire,nil

}


func (incc *QuestionnaireCollectionClass) DeleteQuestionnaire(questionnaire *Questionnaire) error{

	filter:=bson.D{{"questionnaireid",questionnaire.QuestionnaireID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (incc *QuestionnaireCollectionClass) UpdateQuestionnaire (questionnaire *Questionnaire)  (*Questionnaire, error){

	filter:=bson.D{{"questionnaireid",questionnaire.QuestionnaireID}}
	update:=bson.D{{"$set",bson.D{
		{"questions",questionnaire.Questions},
		{"starttime",questionnaire.StartTime},
		{"endtime",questionnaire.EndTime},
		{"internship_id", questionnaire.InternshipID},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return questionnaire,nil
}



func (incc *QuestionnaireCollectionClass) GetQuestionnaireFromInternship (id int64)  (*Questionnaire, error)  {
	filter:=bson.D{{"internshipid",id}}
	questionnnaire:=&Questionnaire{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&questionnnaire)
	if err!=nil{
		return nil,err
	}
	return questionnnaire,nil
}






