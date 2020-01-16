package main

import (
	"fmt"
	"net/http"
)

//myHandler implements Handler now. Handler is
type myHandler struct {
	greeting string
}


//http.Handle(pattern string , Handler ) - handler interface with single ServeHTTP function
//implementing Handler interface must be done here
func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", h.greeting)))
}

func main(){
	//our handler implements ServeHTTP and method is called here
	http.Handle("/", &myHandler{greeting: "blaah"})

	//if nil DefaultServerMux is used and DSM gets registered handler
	//http.ListenAndServe(":8000",nil)
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000",nil)

}