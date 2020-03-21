package Internship

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	collection *mongo.Collection
)

type MongoConfig struct {
	Host string
	Database string
	Port string
}

type InternsipCollectionClass struct{
	dbcon *mongo.Database
}

func NewInternshipCollection(config MongoConfig) (InternshipCollection, error){

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
	collection=db.Collection("Internship")
	return &InternsipCollectionClass{dbcon:db,},nil
}


func(incc *InternsipCollectionClass) GetInternships() ([]*Internship, error){

	findOptions:=options.Find()
	var internships []*Internship
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var internship Internship
		err:=cur.Decode(&internship)
		if err!=nil{
			return nil,err
		}
		internships = append(internships,&internship)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return internships,nil
}


func (incc *InternsipCollectionClass) AddInternship(intern *Internship) (*Internship,error){

	internships,err:=incc.GetInternships()
	n:=len(internships)
	if n!=0{
		internship:=internships[n-1]
		intern.InternshipID = internship.InternshipID+1
	}else{
		intern.InternshipID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), intern)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return intern,nil

}

func (incc *InternsipCollectionClass) GetInternship(id int64) (*Internship,error){

	filter:=bson.D{{"internshipid",id}}
	internship:=&Internship{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&internship)
	if err!=nil{
		return nil,err
	}
	return internship,nil

}


func (incc *InternsipCollectionClass) DeleteInternship(internship *Internship) error{

	filter:=bson.D{{"internshipid",internship.InternshipID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (incc *InternsipCollectionClass) UpdateInternship (internship *Internship)  (*Internship, error){

	filter:=bson.D{{"internshipid",internship.InternshipID}}
	update:=bson.D{{"$set",bson.D{
		{"title",internship.Title},
		{"starttime",internship.StartTime},
		{"endtime",internship .EndTime},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return internship,nil
}








