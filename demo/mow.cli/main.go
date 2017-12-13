package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("cp", "Copy files around")

	app.Spec = "[-r] SRC... DST"

	var (
		recursive = app.BoolOpt("r recursive", false, "Copy files recursively")
		src       = app.StringsArg("SRC", nil, "Source files to copy")
		dst       = app.StringArg("DST", "", "Destination where to copy files to")
	)

	app.Action = func() {
		fmt.Printf("Copying %v to %s [recursively: %v]\n", *src, *dst, *recursive)
	}

	app.Run(os.Args)
}