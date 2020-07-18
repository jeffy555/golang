package main

import (
	"log"
	"os"
	"text/template"
	"time"
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
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("date.gohtml"))
}

func MonthDayYear(t time.Time) string {
	return t.Format("01-01-2006")
}

//https://golang.org/pkg/text/template/#FuncMap
//type FuncMap map[string]interface{}

var fm = template.FuncMap{
	"fdateMDY": MonthDayYear,
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

	err := tpl.ExecuteTemplate(os.Stdout, "date.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
