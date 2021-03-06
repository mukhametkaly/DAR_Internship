package Questionnaire

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
	GetQuestionnaires(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	AddQuestionnaire(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetQuestionnaire(idParam string) func(w http.ResponseWriter,r *http.Request)
	DeleteQuestionnaire(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	UpdateQuestionnaire(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetQuestionnaireFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	Questnr QuestionnaireInInternship
}

func NewEndpointsFactory(questnr QuestionnaireInInternship) Endpoints{
	return &endpointsFactory{
		Questnr: questnr,
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

func (ef *endpointsFactory) GetQuestionnaires(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		questionnaires, err := ef.Questnr.GetQuestionnaires()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, questionnaires)
	}
}


func (ef *endpointsFactory) AddQuestionnaire(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		questionnaire:=&Questionnaire{}
		if err:= json.Unmarshal(data,&questionnaire);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.Questnr.AddQuestionnaire(questionnaire)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetQuestionnaire(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars:=mux.Vars(r)
		paramid, paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		questionnnaire,err:=ef.Questnr.GetQuestionnaire(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,questionnnaire)
	}
}


func (ef *endpointsFactory) DeleteQuestionnaire(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		questionnaire,err:=ef.Questnr.GetQuestionnaire(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.Questnr.DeleteQuestionnaire(questionnaire)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Questionnaire was deleted")
	}

}


func (ef *endpointsFactory) UpdateQuestionnaire(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		questionnaire,err:=ef.Questnr.GetQuestionnaire(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&questionnaire);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_questionnaire,err:=ef.Questnr.UpdateQuestionnaire(questionnaire)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_questionnaire)
	}
}
func (ef *endpointsFactory) GetQuestionnaireFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		vars:=mux.Vars(r)
		paramid, paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		questionnaire,err:=ef.Questnr.GetQuestionnaireFromInternship(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,questionnaire)
	}
}

