package main

import (
	"fmt"

	sa "github.com/grayson/structuredannotations/src"
)

func main() {
	a := sa.NewAnnotation(sa.WithStyle(sa.Info))

	a.Add(
		&sa.Heading{
			Text:  sa.Text("Hello, world!"),
			Level: sa.H1,
		},
		sa.CreateParagraph(
			sa.Text("This is a"),
			sa.StrongEmphasis{Text: sa.Text("demonstration")},
			sa.Text("of using Annotations from Go code by leveraging"),
			&sa.Link{
				Text:        sa.Text("bitrise-plugins-annotations"),
				Destination: "https://github.com/bitrise-io/bitrise-plugins-annotations",
			},
			sa.Text("."),
		),
	)

	fmt.Println("Sending the following to the annotation server:")
	fmt.Println(a.Markdown())
	err := a.Send()
	fmt.Printf("error: %v\n", err)
}
