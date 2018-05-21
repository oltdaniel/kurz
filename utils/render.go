package utils

// Load libraries
import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/gin/render"

import "html/template"
import "net/http"

// Define the structs we need
type Render struct {
  templates map[string]*template.Template
}

type RenderTemplate struct {
  Template *template.Template
  Name string
  Data interface{}
}

// Initialize a new Render
func RenderEngine() Render {
  return Render{
    make(map[string]*template.Template),
  }
}

func (r Render) Instance(name string, data interface{}) render.Render {
  // load template from render engine
  t := r.templates[name]

  // Build new template object
  return RenderTemplate{
    t,
    name,
    data,
  }
}

func (r *Render) Add(name string, layout string) {
  // Compile template
  t, err := template.ParseFiles("views/" + name, "views/" + layout)

  // Check for error
  if err != nil {
    panic(err)
  }

  // Add template
  r.templates[name] = t
}

func (r *Render) HTML(c *gin.Context, code int, name string, data interface{}) {
  // Build new template instance
  template := r.Instance(name, data)

  // Render with gin context
  c.Render(code, template)
}

func (t RenderTemplate) Render(w http.ResponseWriter) error {
  // Execute the template
  return t.Template.Execute(w, t.Data)
}

// Placeholder
func (t RenderTemplate) WriteContentType(w http.ResponseWriter) {}
