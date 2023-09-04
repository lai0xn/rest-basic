package main


import (
	"flag"
	"log"
	"net/http"
)



 func main() {
  port := flag.String("p", "5000", "port")
  flag.Parse()
  r := NewRouter()
  log.Fatal(http.ListenAndServe(":" + *port,r))
}

