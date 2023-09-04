package main

import (
	"encoding/json"
	"fmt"
)

func ParseUsers(users []string)[]User{
  var response []User
  base := ""
  for i := 0;i < len(users) -1 ; i++ {
    base += users[i]
    base +=","
  }
  
  if err := json.Unmarshal([]byte(base), &response);err != nil{
    panic(err)
  }
  fmt.Println(base)
  
  return response
}
