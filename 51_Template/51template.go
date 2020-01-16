package main

import (
	"fmt"
	"html/template"
	"os"
)

func main(){
	templateString := "Lemonade Stand Supply"
//template exists within html/temlate and  text/template. html/template automatically escapses string into literal strings to avoid security issues.
//text/template template doesnt have it
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	//os.Stdout = standard output
	err = t.Execute(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}


}
