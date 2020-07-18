package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type characters struct {
	Name  string
	Motto string
}

type Evil struct {
	Name  string
	Motto string
}
type Battle struct {
	God []characters
	Bad []Evil
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("char.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	char := characters{
		Name:  "Moses",
		Motto: "Free the Israel from Egypt",
	}
	char1 := characters{
		Name:  "Noah",
		Motto: "Ark to save the human from Satan",
	}
	char2 := characters{
		Name:  "Joshua",
		Motto: "War to capture Milk and honey",
	}
	char3 := characters{
		Name:  "Jesus",
		Motto: "Love",
	}
	ev1 := Evil{
		Name:  "Serpent",
		Motto: "Deceived Adam and made him to eat Apple",
	}
	ev2 := Evil{
		Name:  "Lucifer",
		Motto: "Deceived human beings",
	}

	evi := []Evil{ev1, ev2}
	charac := []characters{char, char1, char2, char3}

	battle_war := Battle{
		God: charac,
		Bad: evi,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "char.html", battle_war)
	if err != nil {
		log.Fatalln(err)
	}
}
