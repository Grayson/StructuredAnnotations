package main

import (
	"fmt"
	"time"

	sa "github.com/grayson/structuredannotations/src"
)

func basicStructuredCode() {
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

func templateWithSwappable() {
	const delay = 5 * time.Second

	a := sa.NewAnnotation(sa.WithStyle(sa.Warning), sa.WithUniqueContext())

	t, _ := sa.NewTemplate(`
# This is a templated annotation!

We're going to populate the following list:

{{.list}}

It should update dynamically as you watch!
`)
	a.Add(t)

	swap := sa.CreateSwappable(sa.Text("This will be swapped!"))

	list := sa.CreateOrderedList(
		sa.Text("First item"),
		swap,
		sa.Text("Last item"),
	)
	t.Add("list", list)
	fmt.Printf("Initial:\n%v\n\n***\n\n", a.Markdown())

	update := func(s string) {
		<-time.NewTimer(delay).C
		if len(s) == 0 {
			swap.Delete()
		} else {
			swap.Swap(sa.Text(s))
		}
		err := a.Send()
		fmt.Printf("New markdown:\n***\n%v\n", a.Markdown())
		fmt.Printf("error: %v\n", err)
	}
	update("Swapped text!")
	update("")
	update("This was deleted and restored!")

	<-time.NewTimer(delay).C
	list.Add(&sa.Link{
		Text:        sa.Text("Added a link!"),
		Destination: "https://github.com/Grayson/StructuredAnnotations",
	})
	err := a.Send()
	fmt.Printf("Final markdown:\n***\n%v\n", a.Markdown())
	fmt.Printf("error: %v\n", err)
}

func main() {
	basicStructuredCode()
	templateWithSwappable()
}
