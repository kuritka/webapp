package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"webapp/54_ViewModel/viewmodel"
	"webapp/61_Controllers/model"
	"webapp/common/guard"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s *shop) setTemplate(template *template.Template) {
	s.shopTemplate = template
}

func (s *shop) registerRoutes() {
	http.HandleFunc("/shop", s.handle)
	http.HandleFunc("/shop/", s.handle)
}

func (s *shop) handle(writer http.ResponseWriter, request *http.Request) {
	//with Gorilla Mux I don't need to do this! I can concentrate on header parameters, and other queries without parsing
	categoryPattern, err := regexp.Compile(`/shop/(\d+)`)
	if err != nil {
		fmt.Println(err)
	}
	matches := categoryPattern.FindStringSubmatch(request.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		//we know it is save protection so don't need to handle errors
		//at this point I can show completely different homeTemplate
		s.handleCategory(writer, categoryID)
		return
	}
	vm := viewmodel.NewBase()
	err = s.shopTemplate.Execute(writer, vm)
	guard.FailOnError(err, "shop template")
}

func (s *shop) handleCategory(writer http.ResponseWriter, categoryId int) {
	var colors = model.GetColors(categoryId)
	err := s.categoryTemplate.Execute(writer, viewmodel.NewColors(colors))
	guard.FailOnError(err, "category template")
}
