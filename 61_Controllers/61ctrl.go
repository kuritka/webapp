package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"webapp/61_Controllers/controller"

	"webapp/common/guard"
)

const layout = "/_layoutvm.html"

func main() {
	fmt.Println("listening on http://localhost:8000/shop")
	fmt.Println("listening on http://localhost:8000/shop/0,http://localhost:8000/shop/1, http://localhost:8000/shop/, http://localhost:8000/shop/7, http://localhost:8000/shop/aaa")
	templates := populateTemplates()
	//now main looks as I expected how it should look like
	//kicks off controller layer
	controller.Startup(templates)
	err := http.ListenAndServe(":8000", nil)
	guard.FailOnError(err, "unable to run server on port :8000")
}

func populateTemplates() map[string]*template.Template {
	const basePath = "templates"

	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + layout))

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
