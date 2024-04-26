// yaegi extract github.com/knadh/koanf/providers/file

package imports

import (
	"reflect"

	"github.com/knadh/koanf/providers/file"
)

func init() {
	Symbols["github.com/knadh/koanf/providers/file/file"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Provider": reflect.ValueOf(file.Provider),

		// type definitions
		"File": reflect.ValueOf((*file.File)(nil)),
	}
}
