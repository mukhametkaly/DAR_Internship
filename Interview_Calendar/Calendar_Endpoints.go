package Interview_Calendar

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
	AddInterviewCalendar(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetInterviewCalendars(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	UpdateInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	DeleteInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetInternviewCalendarFromCourses (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	InterCal CourseInterviewsCal
}

func NewEndpointsFactory(interCal CourseInterviewsCal) Endpoints{
	return &endpointsFactory{
		InterCal: interCal,
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

func (ef *endpointsFactory) GetInterviewCalendars(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		course, err := ef.InterCal.GetInterviewCals()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, course)
	}
}

func (ef *endpointsFactory) AddInterviewCalendar(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		interviewCalendar:=&InterviewCalendar{}
		if err:= json.Unmarshal(data,&interviewCalendar);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.InterCal.AddInterviewCal(interviewCalendar)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		vars:=mux.Vars(r)
		paramid, paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		interviewCalendar,err:=ef.InterCal.GetInterviewCal(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if roleAndId[0] != "HR"{
			if roleAndId[0] == "L" {
				if roleAndId[1] != strconv.FormatInt(interviewCalendar.LecturerID, 10)  {
					http.Error(w, "StatusBadRequest", http.StatusBadRequest)
					return
				}
			} else if roleAndId[1] == "I" {
				if roleAndId[1] != strconv.FormatInt(interviewCalendar.InternID, 10)  {
					http.Error(w, "StatusBadRequest", http.StatusBadRequest)
					return
				}

			} else {
				http.Error(w, "StatusBadRequest", http.StatusBadRequest)
				return
			}
		}
		respondJSON(w,http.StatusOK,interviewCalendar)

	}
}


func (ef *endpointsFactory) DeleteInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
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
		interviewCalendar,err:=ef.InterCal.GetInterviewCal(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.InterCal.DeleteInterviewCal(interviewCalendar)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Calendar was deleted")
	}

}


func (ef *endpointsFactory) UpdateInterviewCalendar(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		interviewCalendar,err:=ef.InterCal.GetInterviewCal(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&interviewCalendar);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_interviewCalendar,err:=ef.InterCal.UpdateInterviewCal(interviewCalendar)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_interviewCalendar)
	}
}

func (ef *endpointsFactory) GetInternviewCalendarFromCourses (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
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
		interns,err:=ef.InterCal.GetInternviewCalendarFromCourses(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,interns)
	}
}




