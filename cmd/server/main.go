package main


import (
	"log"
	"github.com/borjamunozf/proglog/interval/server"
)

func main(){
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
