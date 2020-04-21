package Intern

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mukhametkaly/DAR_Internship/Account"
	"github.com/mukhametkaly/DAR_Internship/internship"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	collection *mongo.Collection
)


type InternCollectionClass struct{
	dbcon *mongo.Database
}


func NewInternCollection(config Internship.MongoConfig) (InternCollection, error){

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
	collection=db.Collection("Interns")
	return &InternCollectionClass{dbcon:db,},nil
}


func(cocc *InternCollectionClass) GetInterns() ([]*Intern, error){

	findOptions:=options.Find()
	var interns []*Intern
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var intern Intern
		err:=cur.Decode(&intern)
		if err!=nil{
			return nil,err
		}
		interns = append(interns,&intern)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return interns,nil
}


func (cocc *InternCollectionClass) AddIntern(intrn *Intern) (*Intern,error){

	interns,err:=cocc.GetInterns()
	n:=len(interns)
	if n!=0{
		intern:=interns[n-1]
		intrn.InternID = intern.InternID+1
	}else{
		intrn.InternID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), intrn)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return intrn,nil

}

func (cocc *InternCollectionClass) GetIntern(id int64) (*Intern,error){

	filter:=bson.D{{"internid",id}}
	intern:=&Intern{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&intern)
	if err!=nil{
		return nil,err
	}
	return intern,nil

}


func (cocc *InternCollectionClass) DeleteIntern(intern *Intern) error{

	filter:=bson.D{{"internid",intern.InternID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (cocc *InternCollectionClass) UpdateIntern (intern *Intern)  (*Intern, error){

	filter:=bson.D{{"internid",intern.InternID}}
	update:=bson.D{{"$set",bson.D{
		{"name",intern.Name},
		{"mail",intern.Mail},
		{"password", intern.Password},
		{"contestid",intern.ContestID},
		{"questionnaireid", intern.QuestionnaireID},
		{"courseid", intern.CourseID},
		{"status", intern.Status},
		{"contestscore", intern.ContestScore},
		{"contestusername", intern.ContestUsername},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return intern,nil
}
func (cocc *InternCollectionClass) GetInternsFromCourses (id int64)  ([]*Intern, error)  {
	filter:=bson.D{{"courseid",id}}
	options:=options.Find()
	var interns []*Intern
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var intern Intern
		err:=cur.Decode(&intern)
		if err!=nil{
			return nil,err
		}
		interns = append(interns,&intern)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return interns,nil
}

func (cocc *InternCollectionClass) GetInternByUsername (username string) (*Intern, error)  {

	filter:=bson.D{{"username",username}}
	intern:=&Intern{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&intern)
	if err!=nil{
		return nil, err
	}
	return intern, nil

}

func (cocc *InternCollectionClass) Authorization (username string, password string, client *redis.Client )  error{
	intern, err := cocc.GetInternByUsername(username)
	if err!=nil{
		return  errors.New("Invalid username")
	}
	if intern.Password != password {
		return errors.New("Invalid password")
	}
	tokenString := Account.CreateToken()
	value := "I " + strconv.FormatInt(intern.InternID, 10)
	client.Set(tokenString, value, time.Minute * 3)
	data := strings.Split(client.Get(tokenString).String(), " ")
	fmt.Println(data[2])
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}

	fmt.Println(pong, err)
	fmt.Println(tokenString+" "+value)
	return nil
}


