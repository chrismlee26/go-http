package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	Occupation string
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("template/layout.html"))

		data := []User{
			{Name: "John", Occupation: "Student"},
			{Name: "Jane", Occupation: "Teacher"},
		}
		tmpl.Execute(w, data)

	})

	http.ListenAndServe(":8080", nil)
}
