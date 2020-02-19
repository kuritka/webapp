package main

import (
	"fmt"
	"net/http"
)

func main() {
	//our handler implements ServeHTTP and method is called here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//serve file built-in handler. We have also StripPrefix Handler (decorator), TimeOutHandler (decorator), NotFound and Redirect handlers...
		//3_ServeFile also sets contentType of particular resources, so CSS is CSS and HTML is HTML
		http.ServeFile(w, r, "public"+r.URL.Path)
	})
	//if nil DefaultServerMux is used and DSM gets registered handler
	//http.ListenAndServe(":8000",nil)
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)

}
