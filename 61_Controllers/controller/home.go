package controller

import (
	"html/template"
	"net/http"
	"webapp/54_ViewModel/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h *home) setTemplate(template *template.Template) {
	h.homeTemplate = template
}

func (h *home) registerRoutes() {
	http.HandleFunc("/home", h.handle)
	http.HandleFunc("/", h.handle)
}

func (h *home) handle(writer http.ResponseWriter, request *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(writer, vm)
}
