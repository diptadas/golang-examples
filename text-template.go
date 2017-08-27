package main

import (
	"os"
	"text/template"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func example_1() {
	// define a template
	const letter = `Dear {{.Name}},
{{- if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{- with .Gift}}
Thank you for the lovely {{.}}.
{{- end}}
Best wishes,
Josie
`
	// Prepare some data to insert into the template
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	recipients := []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient
	for _, r := range recipients {
		check(t.Execute(os.Stdout, r))
		fmt.Println("==============")
	}
}

func main() {
	example_1()
}
