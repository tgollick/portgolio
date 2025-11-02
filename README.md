# Portgolio

A lightning-fast, server-driven portfolio built with Go, HTMX, and Tailwind CSS. This project showcases a return to simplicity—proving that modern web experiences don't require heavy JavaScript frameworks.

![Go](https://img.shields.io/badge/Go-1.23.3-00ADD8?style=flat&logo=go)
![HTMX](https://img.shields.io/badge/HTMX-1.9.10-3366CC?style=flat)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.0-38B2AC?style=flat&logo=tailwind-css)

## 🎯 Project Overview

Portgolio is my personal portfolio website that demonstrates full-stack capabilities with a focus on performance, simplicity, and developer control. After working extensively with frameworks like Next.js, I wanted to explore the other side—building from the ground up with complete control over the server, rendering pipeline, and user experience.

This project serves as both a portfolio showcase and a technical statement: modern web development doesn't always need complex build processes, massive JavaScript bundles, or client-side routing to deliver exceptional experiences.

## ✨ Key Features

- **Server-Side Rendering**: Pure HTML rendered on the server for instant page loads
- **Progressive Enhancement**: HTMX for dynamic interactions without heavy JavaScript
- **Blazing Fast**: Go's performance combined with minimal client-side overhead
- **Clean Architecture**: Simple, maintainable codebase with clear separation of concerns
- **Modern Design**: Tailwind CSS for a polished, responsive interface
- **Developer Experience**: Hot reload with Air for rapid development

## 🛠️ Tech Stack

### Backend
- **Go 1.23.3**: Chosen for its performance, simplicity, and excellent standard library
- **html/template**: Native Go templating for type-safe, server-rendered HTML
- **net/http**: Leveraging Go's robust built-in HTTP server

### Frontend
- **HTMX 1.9.10**: For seamless AJAX requests and dynamic updates without writing JavaScript
- **Tailwind CSS**: Utility-first CSS for rapid UI development
- **Vanilla HTML**: Semantic, accessible markup

### Development Tools
- **Air**: Live reload for Go applications during development
- **Git**: Version control

## 🧠 Why I Built This

After building multiple production applications with frameworks like Next.js, I wanted to understand what we gain—and what we sacrifice—with modern meta-frameworks. This project explores:

1. **Full Control**: Complete ownership of the web server, routing, and rendering pipeline
2. **Performance First**: Eliminating unnecessary JavaScript and leveraging server-side rendering
3. **Simplicity**: Reducing complexity without sacrificing user experience
4. **Learning**: Deepening my understanding of HTTP, templating, and web fundamentals

The result is a portfolio that loads instantly, requires minimal bandwidth, and demonstrates that thoughtful architecture can deliver excellent experiences without heavy frameworks.

## 📁 Project Structure

```
portgolio/
├── main.go                 # Application entry point and route handlers
├── templates/
│   ├── base.html          # Base layout template
│   ├── index.html         # Home page
│   └── projects.html      # Projects showcase
├── static/
│   ├── icons/             # SVG icons
│   └── Thomas_Gollick_Logo.png
├── go.mod                 # Go module definition
└── .air.toml             # Air configuration for hot reload
```

## 🚀 Getting Started

### Prerequisites

- Go 1.23.3 or higher
- Air (for development hot reload)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/tgollick/portgolio.git
   cd portgolio
   ```

2. **Install Air for development** (optional but recommended)
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

3. **Run the application**

   With Air (hot reload):
   ```bash
   air
   ```

   Or with standard Go:
   ```bash
   go run main.go
   ```

4. **Open your browser**
   ```
   http://localhost:8080
   ```

## 🏗️ Technical Decisions

### Why Go?
- **Performance**: Fast compilation and execution with minimal overhead
- **Standard Library**: Excellent built-in HTTP server and templating
- **Simplicity**: Clean syntax and straightforward concurrency model
- **Learning**: Wanted to deepen my understanding of lower-level web concepts

### Why HTMX?
- **Progressive Enhancement**: HTML-first approach with JavaScript as enhancement
- **Reduced Complexity**: No need for complex state management or virtual DOM
- **Small Bundle**: ~14KB compared to hundreds of KB for modern frameworks
- **Server-Driven**: Aligns with Go's strengths in server-side rendering

### Why Not Next.js/React?
This isn't a criticism of frameworks—I've built production apps with Next.js and appreciate its power. This project is about:
- Understanding what's possible without them
- Appreciating the tradeoffs we make when choosing frameworks
- Demonstrating versatility across different architectural approaches
- Proving that simpler solutions can work for specific use cases

## 📝 Code Highlights

### Type-Safe Templating
```go
type Project struct {
    Title string
    Desc  string
    Image string
    Tags  []string
}
```

Go's type system ensures compile-time safety for template data, catching errors before runtime.

### Clean Handler Pattern
```go
func projectsHandler(w http.ResponseWriter, r *http.Request) {
    data := Data{
        PageTitle:   "Projects | Thomas Portfolio",
        ProjectData: projectData,
    }
    
    if err := projectsTmpl.ExecuteTemplate(w, "base", data); err != nil {
        http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
        return
    }
}
```

Simple, readable handlers that are easy to test and maintain.

## 🎨 Design Philosophy

- **Content First**: Fast initial page loads with server-rendered HTML
- **Progressive Enhancement**: Enhanced with HTMX for smooth interactions
- **Accessibility**: Semantic HTML and keyboard navigation
- **Responsive**: Mobile-first design with Tailwind utilities
- **Performance**: Minimal JavaScript, optimized assets

## 🔮 Future Enhancements

While the core pages (Home, Projects, About, Contact) will be complete in the final version, potential technical enhancements include:

- Markdown-based blog with Go templates
- Dynamic project filtering with HTMX
- Contact form with Go backend validation
- Analytics integration
- Docker containerization for deployment

## 👨‍💻 About Me

I'm a full-stack developer passionate about clean code, thoughtful design, and continuous learning. This portfolio represents my journey in exploring different architectural approaches and my commitment to understanding web fundamentals.

**Connect with me:**
- GitHub: [@tgollick](https://github.com/tgollick)
- LinkedIn: [Thomas Gollick](https://linkedin.com/in/thomasgollick)

## 📄 License

This project is open source and available under the MIT License.
