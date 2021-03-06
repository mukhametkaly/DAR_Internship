package Contest

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Endpoints interface {
	AddContest(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetContests(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetContest(idParam string) func(w http.ResponseWriter,r *http.Request)
	UpdateContest(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	DeleteContest(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetContestsFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	Contst ContestsinInternship
}

func NewEndpointsFactory(contst ContestsinInternship) Endpoints{
	return &endpointsFactory{
		Contst: contst,
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func (ef *endpointsFactory) GetContests(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		contest, err := ef.Contst.GetContests()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, contest)
	}
}

func (ef *endpointsFactory) AddContest(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return

		}
		contest:=&Contest{}
		if err:= json.Unmarshal(data,&contest);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.Contst.AddContest(contest)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetContest(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			paramid, paramerr := vars[idParam]
			if !paramerr {
				respondJSON(w, http.StatusBadRequest, "Не был передан аргумент")
				return
			}
			id, err := strconv.ParseInt(paramid, 10, 10)
			contest, err := ef.Contst.GetContest(id)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondJSON(w, http.StatusOK, contest)

	}
}


func (ef *endpointsFactory) DeleteContest(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}

		vars:=mux.Vars(r)
		paramid,paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		course,err:=ef.Contst.GetContest(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.Contst.DeleteContest(course)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Contest was deleted")
	}

}


func (ef *endpointsFactory) UpdateContest(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}

		vars:=mux.Vars(r)
		paramid,paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		contest,err:=ef.Contst.GetContest(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&contest);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_contest,err:=ef.Contst.UpdateContest(contest)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_contest)
	}
}

func (ef *endpointsFactory) GetContestsFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] == "HR" || roleAndId[0]== "L" {

			vars := mux.Vars(r)
			paramid, paramerr := vars[idParam]
			if !paramerr {
				respondJSON(w, http.StatusBadRequest, "Не был передан аргумент")
				return
			}
			id, err := strconv.ParseInt(paramid, 10, 10)
			interns, err := ef.Contst.GetContestsFromInternship(id)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondJSON(w, http.StatusOK, interns)
		}else {
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
	}
}



