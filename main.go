package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Principle struct {
  Title string
  Desc string
  Icon string
} 

// parse templates on init
func init() {
	t := template.Must(template.ParseGlob("templates/*.html"))
	t = template.Must(t.ParseGlob("templates/partials/*.html"))
	tmpl = t
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	principlesData := []Principle{
		{
			Title: "Clean Code",
			Desc:  "Writing maintainable, scalable code with modern best practices and attention to detail.",
			Icon:  "/static/icons/code.svg",
		},
		{
			Title: "Thoughtful Design",
			Desc:  "Creating intuitive interfaces that balance aesthetics with usability and accessibility.",
			Icon:  "/static/icons/palette.svg",
		},
		{
			Title: "Performance",
			Desc:  "Optimizing every interaction for speed, efficiency, and delightful user experiences.",
			Icon:  "/static/icons/zap.svg",
		},
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", principlesData); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		log.Println("template error:", err)
		return
	}
	log.Println("Serving index page")
}


