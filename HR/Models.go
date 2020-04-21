package HR

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mukhametkaly/DAR_Internship/Account"
	Internship "github.com/mukhametkaly/DAR_Internship/internship"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type HR struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type HRCollection interface {
	AddHR(hr *HR) (*HR, error)
	UpdateHR(hr *HR) (*HR, error)
	DeleteHR(hr *HR) error
	GetHRByUsername (username string)  (*HR, error)
	Authorization (username string, password string, client *redis.Client) error

}

var (
	collection *mongo.Collection
)


type HRCollectionClass struct{
	dbcon *mongo.Database
}


func NewHRCollection(config Internship.MongoConfig) (HRCollection, error){

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
	return &HRCollectionClass{dbcon:db,},nil
}



func (hrcc *HRCollectionClass) AddHR(hr *HR) (*HR,error){
	_, err := hrcc.GetHRByUsername(hr.UserName)
	if err == nil {
		return nil, errors.New("HR with the same username are exist")
	}

	insertResult,err:=collection.InsertOne(context.TODO(), hr)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return hr,nil

}



func (hrcc *HRCollectionClass) DeleteHR(hr *HR) error{

	filter:=bson.D{{"username",hr.UserName}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (hrcc *HRCollectionClass) UpdateHR (hr *HR)  (*HR, error){
	filter:=bson.D{{"username",hr.UserName}}
	update:=bson.D{{"$set",bson.D{
		{"username",hr.UserName},
		{"password", hr.Password},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return hr,nil
}

func (hrcc *HRCollectionClass) GetHRByUsername (username string) (*HR, error)  {

	filter:=bson.D{{"username",username}}
	hr:=&HR{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&hr)
	if err!=nil{
		return nil, err
	}
	return hr, nil

}

func (hrcc *HRCollectionClass) Authorization (username string, password string, client *redis.Client)  error{
	hr, err := hrcc.GetHRByUsername(username)
	if err!=nil{
		return  errors.New("Invalid username")
	}
	if hr.Password != password {
		return errors.New("Invalid password")
	}

	tokenString := Account.CreateToken()
	value := "HR"
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





