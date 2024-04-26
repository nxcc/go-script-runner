package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitfield/script"
	"github.com/spf13/cast"
)

func main() {
	date, _ := script.Exec("date").String()
	fmt.Println("date:", strings.TrimSpace(date))
	boolean := cast.ToBool("0")
	fmt.Println("bool:", boolean)
	fmt.Printf("args: %+v\n", os.Args)
}

func Main() { main() }
