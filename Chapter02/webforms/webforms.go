// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
	//    "fmt"
	"net/http"
	"text/template"

	//    "errors"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Home(writer http.ResponseWriter, reader *http.Request) {
	// var template_html *template.Template
	var template_html = template.Must(template.ParseFiles("main.html"))
	err := template_html.Execute(writer, nil)
	checkError(err)

}

func main() {
	log.Println("Server started on: http://localhost:8000")
	//  var template_html *template.Template
	//template_html = template.Must(template.ParseFiles("main.html"))
	http.HandleFunc("/", Home)
	//  http.HandleFunc("/show", Show)
	//  http.HandleFunc("/new", New)
	//  http.HandleFunc("/edit", Edit)
	//  http.HandleFunc("/insert", Insert)
	//  http.HandleFunc("/update", Update)
	//    http.HandleFunc("/delete", Delete)
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}
