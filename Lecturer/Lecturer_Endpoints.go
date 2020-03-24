package Lecturer

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Endpoints interface {
	AddLecturer() func(w http.ResponseWriter,r *http.Request)
	GetLecturers() func(w http.ResponseWriter,r *http.Request)
	GetLecturer(idParam string) func(w http.ResponseWriter,r *http.Request)
	UpdateLecturer(idParam string) func(w http.ResponseWriter,r *http.Request)
	DeleteLecturer(idParam string) func(w http.ResponseWriter,r *http.Request)

}

type endpointsFactory struct {
	Lectrs CourseLecturer
}

func NewEndpointsFactory(lectrs CourseLecturer) Endpoints{
	return &endpointsFactory{
		Lectrs: lectrs,
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

func (ef *endpointsFactory) GetLecturers() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		course, err := ef.Lectrs.GetCourseLecturers()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "Ошибка"+err.Error())
			return
		}
		respondJSON(w, http.StatusOK, course)
	}
}

func (ef *endpointsFactory) AddLecturer() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		lecturer:=&Lecturer{}
		if err:= json.Unmarshal(data,&lecturer);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.Lectrs.AddCourseLecturer(lecturer)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}

func (ef *endpointsFactory) GetLecturer(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars:=mux.Vars(r)
		paramid, paramerr:=vars[idParam]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}
		id,err:=strconv.ParseInt(paramid,10,10)
		lecturer,err:=ef.Lectrs.GetCourseLecturer(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,lecturer)
	}
}


func (ef *endpointsFactory) DeleteLecturer(idParam string) func(w http.ResponseWriter,r *http.Request){
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
		lecturer,err:=ef.Lectrs.GetCourseLecturer(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.Lectrs.DeleteCourseLecturer(lecturer)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Intern was deleted")
	}

}


func (ef *endpointsFactory) UpdateLecturer(idParam string) func(w http.ResponseWriter,r *http.Request){
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
		lecturer,err:=ef.Lectrs.GetCourseLecturer(id)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		if err:=json.Unmarshal(data,&lecturer);err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		updated_lecturer,err:=ef.Lectrs.UpdateCourseLecturer(lecturer)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,updated_lecturer)
	}
}

