package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello go http server"))
}

func ping(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Pone"))
}

func main(){
	http.HandleFunc("/",hello)
	http.HandleFunc("/ping",ping)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatalln(err)
	}
}