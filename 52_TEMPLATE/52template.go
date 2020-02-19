package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}

func main() {
	fmt.Println("listening on http://localhost:8000. Execute  template home.html by http://localhost:8000/home")
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//remove slash
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			http.NotFound(w, r)
			//w.WriteHeader(http.StatusFound)
		}
	})
	//these two are handled automatically by fileserver
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}
