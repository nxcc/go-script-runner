package gsr

import (
	"fmt"
	"strings"

	"github.com/bitfield/script"
	"github.com/spf13/cast"
)

var ARGS []string

func Main() {
	date, _ := script.Exec("date").String()
	fmt.Println("date:", strings.TrimSpace(date))
	boolean := cast.ToBool("0")
	fmt.Println("bool:", boolean)
	UUID()
}
