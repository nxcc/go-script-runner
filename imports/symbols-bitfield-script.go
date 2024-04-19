// yaegi extract github.com/bitfield/script

package imports

import (
	"reflect"

	"github.com/bitfield/script"
)

func init() {
	Symbols["github.com/bitfield/script/script"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Args":              reflect.ValueOf(script.Args),
		"Do":                reflect.ValueOf(script.Do),
		"Echo":              reflect.ValueOf(script.Echo),
		"Exec":              reflect.ValueOf(script.Exec),
		"File":              reflect.ValueOf(script.File),
		"FindFiles":         reflect.ValueOf(script.FindFiles),
		"Get":               reflect.ValueOf(script.Get),
		"IfExists":          reflect.ValueOf(script.IfExists),
		"ListFiles":         reflect.ValueOf(script.ListFiles),
		"NewPipe":           reflect.ValueOf(script.NewPipe),
		"NewReadAutoCloser": reflect.ValueOf(script.NewReadAutoCloser),
		"Post":              reflect.ValueOf(script.Post),
		"Slice":             reflect.ValueOf(script.Slice),
		"Stdin":             reflect.ValueOf(script.Stdin),
		// type definitions
		"Pipe":           reflect.ValueOf((*script.Pipe)(nil)),
		"ReadAutoCloser": reflect.ValueOf((*script.ReadAutoCloser)(nil)),
	}
}
