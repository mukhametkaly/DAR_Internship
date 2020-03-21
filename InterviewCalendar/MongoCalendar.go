package InterviewCalendar

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"../internship"
)
var (
	collection *mongo.Collection
)

type InterviewCalendarCollectionClass struct{
	dbcon *mongo.Database
}

func NewInterviewCalendarCollection(config Internship.MongoConfig) (InterviewCalendarCollection, error){

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
	collection=db.Collection("InterviewCalendar")
	return &InterviewCalendarCollectionClass{dbcon:db,},nil
}


func(incc *InterviewCalendarCollectionClass) GetInterviewCalendars() ([]*InterviewCalendar, error){

	findOptions:=options.Find()
	var interviewCalendars []*InterviewCalendar
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var interviewCalendar InterviewCalendar
		err:=cur.Decode(&interviewCalendar)
		if err!=nil{
			return nil,err
		}
		interviewCalendars = append(interviewCalendars,&interviewCalendar)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return interviewCalendars,nil
}


func (incc *InterviewCalendarCollectionClass) AddInterviewCalendar(interCal *InterviewCalendar) (*InterviewCalendar,error){

	interviewCalendars,err:=incc.GetInterviewCalendars()
	n:=len(interviewCalendars)
	if n!=0{
		interviewCalendar:=interviewCalendars[n-1]
		interCal.InterviewCalendarID = interviewCalendar.InterviewCalendarID+1
	}else{
		interCal.InternshipID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), interCal)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return interCal,nil

}

func (incc *InterviewCalendarCollectionClass) GetInterviewCalendar(id int64) (*InterviewCalendar,error){

	filter:=bson.D{{"interviewcalendarid",id}}
	interviewCalendar:=&InterviewCalendar{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&interviewCalendar)
	if err!=nil{
		return nil,err
	}
	return interviewCalendar,nil

}


func (incc *InterviewCalendarCollectionClass) DeleteInterviewCalendar(interviewCalendar *InterviewCalendar) error{

	filter:=bson.D{{"interviewcalendarid",interviewCalendar.InterviewCalendarID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (incc *InterviewCalendarCollectionClass) UpdateInterviewCalendar (interviewCalendar *InterviewCalendar)  (*InterviewCalendar, error){

	filter:=bson.D{{"interviewcalendarid",interviewCalendar.InterviewCalendarID}}
	update:=bson.D{{"$set",bson.D{
		{"come_date",interviewCalendar.comeDate},
		{"cometime",interviewCalendar.comeTime},
		{"lecturer_mail",interviewCalendar.LecturerMail},
		{"duration",interviewCalendar.Duration},
		{"intern_mail",interviewCalendar.InternMail},
		{"internship_id",interviewCalendar.InternshipID},
		{"course_id",interviewCalendar.CourseID},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return interviewCalendar,nil
}

