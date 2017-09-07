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

func example_2() {
	arr_1 := []int{11, 22, 33}
	map_1 := map[int]string{5: "AA", 6: "BB", 7: "CC"}
	map_2 := map[string]string{"a": "A", "b": "B", "c": "C"}

	tmpStr := "Value: {{index . 0}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, arr_1))

	tmpStr = "Value: {{index . 5}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, map_1))

	tmpStr = "Value: {{.a}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, map_2))
}

func example_3() {
	tmpStr := "Iterate: {{range .}} Value = {{.}} {{end}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{11, 22, 33}))

	tmpStr = "Iterate: {{range $i, $v := .}} {{$i}}:{{$v}} {{end}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{11, 22, 33}))

	tmpStr = "Iterate: {{range .}} Value = {{.}} {{else}}Data is empty{{end}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{}))
}

func main() {
	example_0()
}
