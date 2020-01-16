package main

import (
	"fmt"
	"net/http"
)

func main(){
	//our handler implements ServeHTTP and method is called here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello from HandlerFunc"))

	})
	//if nil DefaultServerMux is used and DSM gets registered handler
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000",nil)

}