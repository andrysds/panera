package template

import (
	"html/template"
	"io"
	"log"

	"github.com/andrysds/clarity"
)

// Template represents application template
var Template *template.Template

// Init initializes application template
func Init() {
	var err error
	Template, err = template.ParseGlob("template/*")
	clarity.PanicIfError(err, "error on parsing templates")
	log.Println("* Templates initialized")
}

// Execute applies the template
func Execute(w io.Writer, name string, data interface{}) {
	if err := Template.ExecuteTemplate(w, name, data); err != nil {
		log.Println(err)
	}
}
