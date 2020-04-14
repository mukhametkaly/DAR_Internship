package Lecturer

import (
	"Internship/internship"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
)


type LecturersClass struct{
	dbcon *mongo.Database
}


func NewLecturerCollection(config Internship.MongoConfig) (Lecturers, error){

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
	collection=db.Collection("Lecturers")
	return &LecturersClass{dbcon:db,}, nil
}


func(lec *LecturersClass) GetLecturers() ([]*Lecturer, error){

	findOptions:=options.Find()
	var lecturers []*Lecturer
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var lecturer Lecturer
		err:=cur.Decode(&lecturer)
		if err!=nil{
			return nil,err
		}
		lecturers = append(lecturers,&lecturer)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return lecturers,nil
}


func (lec *LecturersClass) AddLecturer(lectrr *Lecturer) (*Lecturer,error){

	lecturers,err:=lec.GetLecturers()
	n:=len(lecturers)
	if n!=0{
		lecturer:=lecturers[n-1]
		lectrr.LecturerID = lecturer.LecturerID+1
	}else{
		lectrr.LecturerID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), lectrr)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return lectrr,nil

}

func (lec *LecturersClass) GetLecturer(id int64) (*Lecturer,error){

	filter:=bson.D{{"lecturerid",id}}
	lecturer:=&Lecturer{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&lecturer)
	if err!=nil{
		return nil,err
	}
	return lecturer,nil

}


func (lec *LecturersClass) DeleteLecturer(lecturer *Lecturer) error{

	filter:=bson.D{{"lecturerid",lecturer.LecturerID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (lec *LecturersClass) UpdateLecturer (lecturer *Lecturer)  (*Lecturer, error){

	filter:=bson.D{{"lecturerid",lecturer.LecturerID}}
	update:=bson.D{{"$set",bson.D{
		{"name",lecturer.LecturerName},
		{"mail",lecturer.Mail},
		{"courseid", lecturer.CourseID},
		{"password", lecturer.Password},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return lecturer,nil
}
func (cocc *LecturersClass) GetLecturerFromCourses (id int64)  (*Lecturer, error)  {
	filter:=bson.D{{"courseid",id}}
	lecturer:=&Lecturer{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&lecturer)
	if err!=nil{
		return nil,err
	}
	return lecturer,nil

}
func (cocc *LecturersClass) GetLecturerByUsername (username string) (*Lecturer, error){
	filter:=bson.D{{"username",username}}
	lecturer:=&Lecturer{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&lecturer)
	if err!=nil{
		return nil, err
	}
	return lecturer,nil
}

func (cocc *LecturersClass) Authorization (username string, password string) error  {
	lecturer, err := cocc.GetLecturerByUsername(username)
	if err!=nil{
		return errors.New("Invalid username")
	}
	if lecturer.Password != password {
		return errors.New("Invalid password")
	}
	return nil
}






