package main

import (
	"fmt"
	"html/template"
	"os"
)

type us struct {
	Text        string
	Names       []string
	ShouldPrint bool
	Num         int
}

var fn = template.FuncMap{
	"myfun2": func(s string, n int) string {
		r := s
		for i := 1; i <= n; i++ {
			r += s
		}
		return r
	},
}

func main() {
	tpl, err := template.New("").Funcs(fn).ParseGlob("D:\\Workspace\\Go\\snippets-go\\formats\\html-templates\\*.gohtml")
	fmt.Printf("err=%#v\n", err)

	v := us{
		Text: "text<alfa>",
		Names: []string{
			"Alfa", "Beta", "Omega",
		},
		ShouldPrint: true,
		Num:         20,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", v)
	if err != nil {
		fmt.Printf("err=%#v\n", err.Error())
	}

}
