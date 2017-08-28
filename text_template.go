package main

import (
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func example_0() {
	tmpStr := "Hello {{.}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, "world"))

	type student struct {
		Name string
		Age  int
	}
	student_1 := student{"Alice", 21}
	tmpStr = "Name: {{.Name}} and Age: {{.Age}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, student_1))
}

func example_1() {
	tmpStr := "{{23 -}} < {{- 45}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, nil))

	tmpStr = `***
	{{- "Hello" -}}
	***`
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, nil))
}

func main() {
	example_0()
}
