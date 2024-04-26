package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nxcc/go-script-runner/imports"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {

	flagImports := flag.Bool("imports", false, "list available imports")
	flag.Parse()

	if *flagImports {
		for i := range imports.Symbols {
			fmt.Println(i)
		}
		os.Exit(0)
	}

	args := []string{"-"}
	if flag.NArg() > 0 {
		args = flag.Args()
	}

	i := interp.New(interp.Options{Args: args})
	i.Use(imports.Symbols)
	i.Use(stdlib.Symbols)

	switch {
	case flag.NArg() >= 1 && flag.Arg(0) == "-":
		evalFromStdin(i)
	case flag.NArg() == 0:
		evalFromStdin(i)
	case flag.NArg() == 1:
		evalFromFile(flag.Arg(0), i)
	case flag.NArg() > 1 && flag.Arg(0) == "-":
		evalFromStdin(i)
	case flag.NArg() > 1 && flag.Arg(0) != "-":
		evalFromFile(flag.Arg(0), i)
	default:
		panic("case not handled")
	}
}

func evalFromStdin(i *interp.Interpreter) {
	input, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Printf("read stdin: %s\n", err)
		os.Exit(1)
	}
	_, err = i.Eval(string(input))
	if err != nil {
		fmt.Printf("eval: %s\n", err)
		os.Exit(1)
	}
}

func evalFromFile(f string, i *interp.Interpreter) {
	_, err := i.EvalPath(f)
	if err != nil {
		fmt.Printf("eval path: %s\n", err)
		os.Exit(1)
	}
}
