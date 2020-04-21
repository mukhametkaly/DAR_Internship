package Contest

import (
	"context"
	"fmt"
	"github.com/mukhametkaly/DAR_Internship/internship"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
var (
	collection *mongo.Collection
)



type ContestCollectionClass struct{
	dbcon *mongo.Database
}


func NewContestCollection(config Internship.MongoConfig) (ContestCollection, error){

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
	collection=db.Collection("Contest")
	return &ContestCollectionClass{dbcon:db,},nil
}


func(cocc *ContestCollectionClass) GetContests() ([]*Contest, error){

	findOptions:=options.Find()
	var contests []*Contest
	con,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for con.Next(context.TODO()){
		var contest Contest
		err:=con.Decode(&contest)
		if err!=nil{
			return nil,err
		}
		contests = append(contests,&contest)
	}
	if err:=con.Err();err!=nil{
		return nil,err
	}
	con.Close(context.TODO())
	return contests,nil
}


func (cocc *ContestCollectionClass) AddContest(contst *Contest) (*Contest,error){

	contests,err:=cocc.GetContests()
	n:=len(contests)
	if n!=0{
		course:=contests[n-1]
		contst.ContestID = course.ContestID+1
	}else{
		contst.ContestID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), contst)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document",insertResult.InsertedID)
	return contst,nil

}

func (cocc *ContestCollectionClass) GetContest(id int64) (*Contest,error){

	filter:=bson.D{{"contestid",id}}
	contest:=&Contest{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&contest)
	if err!=nil{
		return nil,err
	}
	return contest,nil

}


func (cocc *ContestCollectionClass) DeleteContest(contest *Contest) error{

	filter:=bson.D{{"contestid",contest.ContestID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (cocc *ContestCollectionClass) UpdateContest (contest *Contest)  (*Contest, error){

	filter:=bson.D{{"contestid",contest.ContestID}}
	update:=bson.D{{"$set",bson.D{
		{"title",contest.Title},
		{"url", contest.URL},
		{"message", contest.Message},
		{"starttime",contest.StartTime},
		{"endtime",contest .EndTime},
		{"internshipid", contest.InternshipID},

	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return contest,nil
}
func (cocc *ContestCollectionClass) GetContestFromInternship (id int64)  ([]*Contest, error)  {
	filter:=bson.D{{"internshipid",id}}
	options:=options.Find()
	var contests []*Contest
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var contest Contest
		err:=cur.Decode(&contest)
		if err!=nil{
			return nil,err
		}
		contests = append(contests,&contest)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return contests,nil
}
