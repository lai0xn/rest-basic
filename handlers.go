package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func GetUsers(w http.ResponseWriter,_ *http.Request) {
  users := getUsers()
  json.NewEncoder(w).Encode(&users)
    
}

func GetUser(w http.ResponseWriter,r *http.Request){
  params := mux.Vars(r)
  username := params["username"]
  user := getUser(username)
  json.NewEncoder(w).Encode(user)


}

func CreateUpdateUser(w http.ResponseWriter,r *http.Request){
  var user User
  json.NewDecoder(r.Body).Decode(&user)
  if err := SetUser(user); err!= nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
  }
  w.WriteHeader(http.StatusCreated)
  w.Write([]byte("user added"))
}


func DeleteUser(w http.ResponseWriter,r *http.Request){ 
  params := mux.Vars(r)
  username := params["username"]
  deleteUser(username)
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("user deleted"))

}
