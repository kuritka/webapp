package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"webapp/common/guard"
)

func main() {
	fmt.Println("listening on http://localhost:8000. Execute  template home.html by http://localhost:8000/53")
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//remove slash
		requestedFile := r.URL.Path[1:]
		t := templates[requestedFile+".html"]
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
	err := http.ListenAndServe(":8000", nil)
	guard.FailOnError(err, "unable to run server on port :8000")
}

func populateTemplates() map[string]*template.Template {
	const basePath = "templates"

	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	//load subtemplates
	subtemplates, err := layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html")
	//just check that files exists
	template.Must(subtemplates, err)
	//also can be called like that: template.Must(layout.ParseFiles(basePath + "/_header.html", basePath + "/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	guard.FailOnError(err, "failed to read /content directory")

	fis, err := dir.Readdir(-1)
	guard.FailOnError(err, "failed to read content of /content directory")

	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		guard.FailOnError(err, "failed to open template %s", fi.Name())
		content, err := ioutil.ReadAll(f)
		guard.FailOnError(err, "cannot read content of %s", f.Name())
		err = f.Close()
		guard.FailOnError(err, "unable to close file %s", fi.Name())
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		guard.FailOnError(err, "failed to parse content %s", fi.Name())
		result[fi.Name()] = tmpl
	}
	return result
}
