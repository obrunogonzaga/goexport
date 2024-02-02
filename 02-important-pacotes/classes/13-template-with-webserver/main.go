package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"toUpper": toUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			Curso{"Golang", 40},
			Curso{"Python", 35},
			Curso{"Java", 60},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
