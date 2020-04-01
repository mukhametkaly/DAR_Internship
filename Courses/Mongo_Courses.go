package Courses

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"Internship/internship"
	"log"
)
var (
	collection *mongo.Collection
)


type CourseCollectionClass struct{
	dbcon *mongo.Database
}


func NewCourseCollection(config Internship.MongoConfig) (CourseCollection, error){

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
	collection=db.Collection("Courses")
	return &CourseCollectionClass{dbcon:db,},nil
}


func(cocc *CourseCollectionClass) GetCourses() ([]*Courses, error){

	findOptions:=options.Find()
	var courses []*Courses
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var course Courses
		err:=cur.Decode(&course)
		if err!=nil{
			return nil,err
		}
		courses = append(courses,&course)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return courses,nil
}


func (cocc *CourseCollectionClass) AddCourse(cours *Courses) (*Courses,error){

	courses,err:=cocc.GetCourses()
	n:=len(courses)
	if n!=0{
		course:=courses[n-1]
		cours.CourseID = course.CourseID+1
	}else{
		cours.CourseID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), cours)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return cours,nil

}

func (cocc *CourseCollectionClass) GetCourse(id int64) (*Courses,error){

	filter:=bson.D{{"courseid",id}}
	course:=&Courses{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&course)
	if err!=nil{
		return nil,err
	}
	return course,nil

}


func (cocc *CourseCollectionClass) DeleteCourse(course *Courses) error{

	filter:=bson.D{{"coursid",course.CourseID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (cocc *CourseCollectionClass) UpdateCourse (course *Courses)  (*Courses, error){

	filter:=bson.D{{"courseid",course.CourseID}}
	update:=bson.D{{"$set",bson.D{
		{"title",course.Title},
		{"lecturer",course.LecturerID},
		{"lecturername",course.LecturerName},
		{"lecturermail", course.LecturerMail},
		{"internshipid", course.InternshipID},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return course,nil
}


func (cocc *CourseCollectionClass) GetCoursesFromInternship (id int64)  ([]*Courses, error)  {
	filter:=bson.D{{"internshipid",id}}
	options:=options.Find()
	var courses []*Courses
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var course Courses
		err:=cur.Decode(&course)
		if err!=nil{
			return nil,err
		}
		courses = append(courses,&course)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return courses,nil
}

