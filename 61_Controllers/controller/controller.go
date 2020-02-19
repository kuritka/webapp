package controller

import (
	"html/template"
	"net/http"
)

type ictrl interface {
	registerRoutes()
	handle(writer http.ResponseWriter, request *http.Request)
	setTemplate(*template.Template)
}

var (
	homeController home
	shopController shop
)

func Startup(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()

	//needs to be refactored, controller can have multiple templates
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["category.html"]
	shopController.registerRoutes()

	//these two are handled automatically by fileserver
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
