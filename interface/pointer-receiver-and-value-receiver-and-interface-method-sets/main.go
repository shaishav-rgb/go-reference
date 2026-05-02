package main

import (
	"log"
	"net/http"
)

type home1 struct {}
func (h *home1) ServeHTTP(w http.ResponseWriter, r *http.Request) {   //pointer receiver method
w.Write([]byte("This is my home page"))
}
// func (h home1) ServeHTTP(w http.ResponseWriter, r *http.Request) {    //value receiver method
// w.Write([]byte("This is my home page"))
// }

func main(){
	mux:=http.NewServeMux()

	mux.Handle("/", home1{})  // ServeHTTP is a pointer receiver method() so cannot use value receiver method here
	mux.Handle("/", &home1{})

	log.Print("Starting server on :4000")
	err:=http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}