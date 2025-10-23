package main

import (
  "html/template"
  "log"
  "net/http"
)

// Loading all the htmx templates at once
var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

func main() {
  // Serve the static files (CSS and images)
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Routes
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/project", projectHandler)

  log.Println("Server running at http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := []map[string]string{
		{"ID": "jobmatch", "Title": "Job Matcher", "Desc": "Full-stack job matching app using TF-IDF + Cosine similarity."},
		{"ID": "dsg", "Title": "DSG Home Finance", "Desc": "Mortgage website built for a family business."},
		{"ID": "thera", "Title": "Therapeutic LLM Attempt", "Desc": "AI model for emotional tone detection."},
	}

  tmpl.ExecuteTemplate(w, "index.html", data)
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")

  data := map[string]string{
    "Title": "Project" + id,
    "Desc": "This is a dynamic htmx-rendered partial for " + id,
  }

  tmpl.ExecuteTemplate(w, "project_card.html", data)
}
