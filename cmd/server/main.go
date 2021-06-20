package main


import (
	"log"
	"github.com/borjamunozf/proglog/internal/server"
)

func main(){
	log.Print("Starting http server...")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
