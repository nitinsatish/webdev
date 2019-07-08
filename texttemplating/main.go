package main

import (
	"fmt"
	"os"
	"text/template"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
	tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "09")

	fd, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Cannot create file", err)
	}
	tmpl.ExecuteTemplate(fd, "tpl.gohtml", "Enlightment")

}
