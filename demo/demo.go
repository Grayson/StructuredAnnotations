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

	<-time.NewTimer(3 * time.Second).C
	swap.Swap(sa.Text("Swapped for text!"))
	err := a.Send()
	fmt.Printf("error: %v\n", err)

	<-time.NewTimer(3 * time.Second).C
	swap.Delete()
	err = a.Send()
	fmt.Printf("error: %v\n", err)

	<-time.NewTimer(3 * time.Second).C
	swap.Swap(sa.Text("This was deleted and restored!"))

	fmt.Println(a.Markdown())
	err = a.Send()
	fmt.Printf("error: %v\n", err)
}

func main() {
	basicStructuredCode()
	templateWithSwappable()
}
