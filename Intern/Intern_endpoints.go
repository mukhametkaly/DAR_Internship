package Intern

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/mukhametkaly/DAR_Internship/Account"
	"github.com/mukhametkaly/DAR_Internship/Courses"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Endpoints interface {
	AddIntern() func(w http.ResponseWriter,r *http.Request)
	GetInterns(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetIntern(idParam string, client *redis.Client, cc Courses.CourseCollection) func(w http.ResponseWriter,r *http.Request)
	UpdateIntern(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	DeleteIntern(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	GetInternsFromCourses (idParam string, client *redis.Client, cc Courses.CourseCollection)  func(w http.ResponseWriter,r *http.Request)
	Authorization (client *redis.Client)  func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	Intrn CourseIntern
}

func NewEndpointsFactory(intrn CourseIntern) Endpoints{
	return &endpointsFactory{
		Intrn: intrn,
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

func (ef *endpointsFactory) GetInterns(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		course, err := ef.Intrn.GetInterns()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, course)
	}
}

func (ef *endpointsFactory) AddIntern() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return

		}
		intern:=&Intern{}
		if err:= json.Unmarshal(data,&intern);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.Intrn.AddIntern(intern)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetIntern(idParam string, client *redis.Client, cc Courses.CourseCollection) func(w http.ResponseWriter,r *http.Request) {
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
		intern,err:=ef.Intrn.GetIntern(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}

		course, err := cc.GetCourse(intern.InternID)
		if err != nil {
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if roleAndId[0] != "HR"{
			if (roleAndId[0] == "L" && roleAndId[1] != strconv.FormatInt(course.LecturerID, 10)) || (roleAndId[0] == "I" && roleAndId[1] != paramid)  {
					http.Error(w, "StatusBadRequest", http.StatusBadRequest)
					return

			}

		}

		respondJSON(w,http.StatusOK,intern)
	}
}


func (ef *endpointsFactory) DeleteIntern(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		intern,err:=ef.Intrn.GetIntern(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.Intrn.DeleteIntern(intern)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Intern was deleted")
	}

}


func (ef *endpointsFactory) UpdateIntern(idParam string, client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		RedisData, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(RedisData, " ")


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

		if roleAndId[0] != "HR"{
			if (roleAndId[0] == "I" && roleAndId[1] != paramid) || roleAndId[0] != "I" {
				http.Error(w, "StatusBadRequest", http.StatusBadRequest)
				return
			}
		}

		intern,err:=ef.Intrn.GetIntern(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&intern);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_intern,err:=ef.Intrn.UpdateIntern(intern)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_intern)
	}
}

func (ef *endpointsFactory) GetInternsFromCourses (idParam string, client *redis.Client, cc Courses.CourseCollection)  func(w http.ResponseWriter,r *http.Request) {
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
		interns,err:=ef.Intrn.GetInternsFromCourses(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		course, err := cc.GetCourse(id)
		if err != nil {
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if roleAndId[0] != "HR"{
			if (roleAndId[0] == "L" && roleAndId[1] != strconv.FormatInt(course.LecturerID, 10)) || roleAndId[0] != "L"  {
				http.Error(w, "StatusBadRequest", http.StatusBadRequest)
				return
			}
		}

		respondJSON(w,http.StatusOK,interns)
	}
}

func (ef *endpointsFactory) Authorization (client *redis.Client) func(w http.ResponseWriter,r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		account:=&Account.Account{}
		if err:= json.Unmarshal(data,&account);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		err =ef.Intrn.Authorization(account.UserName, account.Password, client)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}

		respondJSON(w,http.StatusOK, "Hello you are intern")


	}

}



