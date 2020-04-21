package HR

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/mukhametkaly/DAR_Internship/Account"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


type Endpoints interface {
	AddHR(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	UpdateHR(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	DeleteHR(client *redis.Client) func(w http.ResponseWriter,r *http.Request)
	Authorization (client *redis.Client)  func(w http.ResponseWriter,r *http.Request)
}

type endpointsFactory struct {
	HRs HRCollection
}

func NewEndpointsFactory(hrs HRCollection) Endpoints{
	return &endpointsFactory{
		HRs: hrs,
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


func (ef *endpointsFactory) AddHR(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		hr:=&HR{}
		if err:= json.Unmarshal(data,&hr);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		st,err:=ef.HRs.AddHR(hr)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,st)
	}
}



func (ef *endpointsFactory) DeleteHR(username string,  client *redis.Client) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		reqToken := strings.Split(r.Header.Get("Authorization"), " ")
		data, _ := client.Get(reqToken[1]).Result()
		roleAndId := strings.Split(data, " ")
		if roleAndId[0] != "HR"{
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}
		vars:=mux.Vars(r)
		UName,paramerr:=vars[username]
		if !paramerr{
			respondJSON(w,http.StatusBadRequest,"Не был передан аргумент")
			return
		}

		hr,err:=ef.HRs.GetHRByUsername(UName)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		err=ef.HRs.DeleteHR(hr)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}
		respondJSON(w,http.StatusOK,"Intern was deleted")
	}

}


func (ef *endpointsFactory) UpdateHR(client *redis.Client) func(w http.ResponseWriter,r *http.Request){
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
		hr:=&HR{}
		if err:= json.Unmarshal(data,&hr);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}


		_,err=ef.HRs.GetHRByUsername(hr.UserName)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err.Error())
			return
		}

		UpdatedHR,err:=ef.HRs.UpdateHR(hr)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		respondJSON(w,http.StatusOK,UpdatedHR)
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
		err =ef.HRs.Authorization(account.UserName, account.Password, client)
		if err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}

		respondJSON(w,http.StatusOK, "Hello you are HR")

	}

}




