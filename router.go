package main

import (
	

	"github.com/gorilla/mux"
)



func NewRouter() *mux.Router{
  r := mux.NewRouter()
  r.Use(loggerMiddlware)
  r.HandleFunc("/api/users",GetUsers).Methods("GET")
  r.HandleFunc("/api/user/{username}",GetUser).Methods("GET")
  r.HandleFunc("/api/users/create",CreateUpdateUser).Methods("POST")
  r.HandleFunc("/api/users/delete/{username}",DeleteUser).Methods("DELETE")
  return r

}
