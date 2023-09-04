package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)
var (
  ctx context.Context = context.TODO()
)

func Connect() *redis.Client{
  return redis.NewClient(&redis.Options{
    Addr:"localhost:6379",
    Password:"",
    DB:0,
  }) 
}
func SetUser( user User) error{
  r := Connect()
  bytes,err := json.Marshal(&user)
  if err != nil {
    return err
  }
  fmt.Println(string(bytes))
  r.HSet(ctx, "users",user.Username ,string(bytes))
  return nil

}

func getUsers()map[string]string {
  
  r := Connect()
  result := r.HGetAll(ctx, "users").Val()
  fmt.Println(result)
  
  return result
}
func getUser(username string) string {
  r := Connect()
  result := r.HGet(ctx, "users", username).Val()
  return result
  
}


func deleteUser(username string){
  r := Connect()
  r.HDel(ctx, "users", username)
}
