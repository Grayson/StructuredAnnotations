package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	sa "github.com/grayson/structuredannotations/src"
)

func main() {
	templateBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	template, err := sa.NewTemplate(string(templateBytes))
	if err != nil {
		panic(err)
	}

	for _, env := range os.Environ() {
		tmp := strings.SplitN(env, "=", 2)
		template.Add(tmp[0], sa.Text(tmp[1]))
	}

	for _, args := range os.Args {
		tmp := strings.SplitN(args, "=", 2)
		if len(tmp) < 2 {
			continue
		}
		template.Add(tmp[0], sa.Text(tmp[1]))
	}

	a := sa.NewAnnotation(sa.WithStyle(sa.Info))
	a.Add(template)
	fmt.Println(a.Markdown())
	err = a.Send()
	fmt.Printf("error: %v\n", err)
}
