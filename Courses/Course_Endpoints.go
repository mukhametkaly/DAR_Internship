package Courses

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
	AddCourse(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetCourses(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	UpdateCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	DeleteCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetCoursesFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	CrsInIntrnshp CoursesInInternship
}

func NewEndpointsFactory(crsinintrnshp CoursesInInternship) Endpoints{
	return &endpointsFactory{
		CrsInIntrnshp: crsinintrnshp,
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

func (ef *endpointsFactory) GetCourses(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		course, err := ef.CrsInIntrnshp.GetCourses()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}




		respondJSON(w, http.StatusOK, course)
	}
}

func (ef *endpointsFactory) AddCourse(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		courses:=&Courses{}
		if err:= json.Unmarshal(data,&courses);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.CrsInIntrnshp.AddCourse(courses)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")

		vars:=mux.Vars(r)
		paramid, paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		course,err:=ef.CrsInIntrnshp.GetCourse(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if roleAndId[0] != "HR"{
			if roleAndId[0] != "L" {
				http.Error(w, "StatusBadRequest", http.StatusBadRequest)
				return
			} else if roleAndId[1] != strconv.FormatInt(course.LecturerID, 10) {
				http.Error(w, "StatusBadRequest", http.StatusBadRequest)
				return
			}
		}


		respondJSON(w,http.StatusOK,course)
	}
}


func (ef *endpointsFactory) DeleteCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		course,err:=ef.CrsInIntrnshp.GetCourse(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.CrsInIntrnshp.DeleteCourse(course)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Course was deleted")
	}

}


func (ef *endpointsFactory) UpdateCourse(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		course,err:=ef.CrsInIntrnshp.GetCourse(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&course);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_course,err:=ef.CrsInIntrnshp.UpdateCourse(course)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_course)
	}
}


func (ef *endpointsFactory) GetCoursesFromInternship (idParam string, client *redis.Client)  func(w http.ResponseWriter,r *http.Request) {
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
		course,err:=ef.CrsInIntrnshp.GetCoursesFromInternship(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,course)
	}
}

