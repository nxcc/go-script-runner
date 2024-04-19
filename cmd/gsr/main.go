package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"github.com/nxcc/go-script-runner/imports"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {

	flagRun := flag.String("run", "Main", "function to run")
	flagImports := flag.Bool("imports", false, "list available imports")
	flag.Parse()

	var args []string

	if *flagImports {
		for i := range imports.Symbols {
			fmt.Println(i)
		}
		os.Exit(0)
	}

	i := interp.New(interp.Options{})
	i.Use(imports.Symbols)
	i.Use(stdlib.Symbols)

	switch {
	case flag.NArg() >= 1 && flag.Arg(0) == "-":
		evalFromStdin(i)
		args = flag.Args()
	case flag.NArg() == 0:
		evalFromStdin(i)
		args = []string{"-"}
	case flag.NArg() == 1:
		evalFromFile(flag.Arg(0), i)
		args = flag.Args()
	case flag.NArg() > 1 && flag.Arg(0) == "-":
		evalFromStdin(i)
		args = flag.Args()
	case flag.NArg() > 1 && flag.Arg(0) != "-":
		evalFromFile(flag.Arg(0), i)
		args = flag.Args()
	default:
		panic("case not handled")
	}

	_, err := i.Eval(fmt.Sprintf(`
		package gsr
		func init() { ARGS = %#v }
		`, args,
	))
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("gsr." + *flagRun)
	if err != nil {
		panic(err)
	}

	m := v.Interface().(func())
	m()
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
	fi, err := os.Stat(f)
	if err != nil {
		fmt.Printf("stat: %s\n", err)
		os.Exit(1)
	}
	if fi.Mode().IsRegular() {
		i.EvalPath(f)
	}
	if fi.IsDir() {
		fh, err := os.Open(f)
		if err != nil {
			fmt.Printf("open: %s\n", err)
			os.Exit(1)
		}
		fileNames, err := fh.Readdirnames(-1)
		if err != nil {
			fmt.Printf("readdirnames: %s\n", err)
			os.Exit(1)
		}
		for _, fileName := range fileNames {
			filePath := filepath.Join(fh.Name(), fileName)
			if filepath.Ext(filePath) == ".go" {
				if _, err := i.EvalPath(filePath); err != nil {
					fmt.Printf("evalPath: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}
}
