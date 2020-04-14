package Intern


import (
	"Internship/Account"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Endpoints interface {
	AddIntern() func(w http.ResponseWriter,r *http.Request)
	GetInterns() func(w http.ResponseWriter,r *http.Request)
	GetIntern(idParam string) func(w http.ResponseWriter,r *http.Request)
	UpdateIntern(idParam string) func(w http.ResponseWriter,r *http.Request)
	DeleteIntern(idParam string) func(w http.ResponseWriter,r *http.Request)
	GetInternsFromCourses (idParam string)  func(w http.ResponseWriter,r *http.Request)
	Authorization ()  func(w http.ResponseWriter,r *http.Request)

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

func (ef *endpointsFactory) GetInterns() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
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

func (ef *endpointsFactory) GetIntern(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		respondJSON(w,http.StatusOK,intern)
	}
}


func (ef *endpointsFactory) DeleteIntern(idParam string) func(w http.ResponseWriter,r *http.Request){
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


func (ef *endpointsFactory) UpdateIntern(idParam string) func(w http.ResponseWriter,r *http.Request){
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

func (ef *endpointsFactory) GetInternsFromCourses (idParam string)  func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		respondJSON(w,http.StatusOK,interns)
	}
}

func (ef *endpointsFactory) Authorization () func(w http.ResponseWriter,r *http.Request)  {
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
		err =ef.Intrn.Authorization(account.UserName, account.Password)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		log.Println("Hello you are intern --->", account.UserName)
		w.Header().Add("Content-Type", "application/json")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":  time.Now().Add(time.Minute * time.Duration(2)).Unix(),
			"iat":  time.Now().Unix(),
		})
		tokenString, err := token.SignedString([]byte("intern"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"error":"token_generation_failed"}`)
			return
		}
		io.WriteString(w, `{"token":"`+tokenString+`"}`)

		respondJSON(w,http.StatusOK, "Hello you are intern")


	}

}



