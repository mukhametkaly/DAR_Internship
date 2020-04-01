package Courses

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Endpoints interface {
	AddCourse() func(w http.ResponseWriter,r *http.Request)
	GetCourses() func(w http.ResponseWriter,r *http.Request)
	GetCourse(idParam string) func(w http.ResponseWriter,r *http.Request)
	UpdateCourse(idParam string) func(w http.ResponseWriter,r *http.Request)
	DeleteCourse(idParam string) func(w http.ResponseWriter,r *http.Request)
	GetCoursesFromInternship (idParam string)  func(w http.ResponseWriter,r *http.Request)

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

func (ef *endpointsFactory) GetCourses() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		course, err := ef.CrsInIntrnshp.GetCourses()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, course)
	}
}

func (ef *endpointsFactory) AddCourse() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
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

func (ef *endpointsFactory) GetCourse(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		respondJSON(w,http.StatusOK,course)
	}
}


func (ef *endpointsFactory) DeleteCourse(idParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
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


func (ef *endpointsFactory) UpdateCourse(idParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
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


func (ef *endpointsFactory) GetCoursesFromInternship (idParam string)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

