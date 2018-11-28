package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}

type Person struct {
	UserName string
}
