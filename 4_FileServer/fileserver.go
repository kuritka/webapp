package main

import (
	"fmt"
	"net/http"
)


func main(){
	//http.ListenAndServe(":8000",nil)
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000",http.FileServer(http.Dir("public")))

}
