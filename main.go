package main

import (
	"html/template"
	"log"
	"net/http"
)

type Principle struct {
	Title string
	Desc  string
	Icon  string
}

type Project struct {
	Title string
	Desc  string
	Image string
  Tags []string
}

type Data struct {
	PrincipleData []Principle
	ProjectData   []Project
	PageTitle     string
}

var (
	indexTmpl    *template.Template
	projectsTmpl *template.Template
)

func init() {
	indexTmpl = template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	projectsTmpl = template.Must(template.ParseFiles("templates/base.html", "templates/projects.html"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/projects", projectsHandler)

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

	projectData := []Project{
		{
			Title: "AI Job Matching Platform",
			Desc:  "Full-stack AI Job Matching platform using TF-IDF alongside Cosine similarity to compare CV's and jobs.",
			Image: "/static/placeholder.avif",
		},
		{
			Title: "DSG Home Finance",
			Desc:  "Production Full-stack web application for my father’s business DSG Home Finance.",
			Image: "/static/placeholder.avif",
		},
		{
			Title: "Therapeutic LLM",
			Desc:  "Dissertation project providing rich context to reasoning LLM's for improved therapeutic output.",
			Image: "/static/placeholder.avif",
		},
	}

	data := Data{
		PrincipleData: principlesData,
		ProjectData:   projectData,
		PageTitle:     "Home | Thomas Portfolio",
	}

	if err := indexTmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		log.Println("template error:", err)
		return
	}
	log.Println("Serving index page")
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
  projectData := []Project{
    {
      Title: "AI Job Matching Platform",
      Desc:  "Full-stack AI Job Matching platform using TF-IDF alongside Cosine similarity to compare CV's and jobs.",
      Image: "/static/placeholder.avif",
      Tags:  []string{"Next.js", "Web3", "Tailwind CSS", "Typescript", "Stripe"},
    },
    {
      Title: "DSG Home Finance",
      Desc:  "Production Full-stack web application for my father’s business DSG Home Finance.",
      Image: "/static/placeholder.avif",
      Tags:  []string{"Next.js", "Web3", "Tailwind CSS", "Typescript", "Stripe"},
    },
    {
      Title: "Therapeutic LLM",
      Desc:  "Dissertation project providing rich context to reasoning LLM's for improved therapeutic output.",
      Image: "/static/placeholder.avif",
      Tags:  []string{"Next.js", "Web3", "Tailwind CSS", "Typescript", "Stripe"},
    },
  }

	data := Data{
		PageTitle: "Projects | Thomas Portfolio",
    ProjectData: projectData,
	}

	if err := projectsTmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		log.Println("template error:", err)
		return
	}

	log.Println("Serving projects page")
}


