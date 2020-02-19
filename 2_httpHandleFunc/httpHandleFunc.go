package main

import (
	"fmt"
	"html"
	"net/http"
)

func main(){
	//our handler implements ServeHTTP and method is called here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello from HandlerFunc"))

	})


	http.HandleFunc("/blag/aaa", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))


	})



	//if nil DefaultServerMux is used and DSM gets registered handler
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000",nil)

}