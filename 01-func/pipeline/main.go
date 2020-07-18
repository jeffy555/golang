package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

//Variable tpl is a pointer to a template

var tpl *template.Template

//func New
//func New(name string) *Template
//New allocates a new, undefined template with the given name.

//func (*Template) Funcs
//func (t *Template) Funcs(funcMap FuncMap) *Template
//Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed.
//It panics if a value in the map is not a function with appropriate return type or
//if the name cannot be used syntactically as a function in a template.It is legal to overwrite elements of the map.
//The return value is the template, so calls can be chained.

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("char.html"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

var fm = template.FuncMap{
	"DBLE": double,
	"SQ":   square,
}

func main() {
	//Executetemplate
	// ExecuteTemplate applies the template associated with t that has the given name
	// to the specified data object and writes the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel, although if parallel
	// executions share a Writer the output may be interleaved.
	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

	//Passing the data inside here. Passing the value of number 5
	err := tpl.ExecuteTemplate(os.Stdout, "char.html", 5)
	if err != nil {
		log.Fatalln(err)
	}
}
