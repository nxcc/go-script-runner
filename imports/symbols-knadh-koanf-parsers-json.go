// yaegi extract github.com/knadh/koanf/parsers/json

package imports

import (
	"reflect"

	"github.com/knadh/koanf/parsers/json"
)

func init() {
	Symbols["github.com/knadh/koanf/parsers/json/json"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Parser": reflect.ValueOf(json.Parser),

		// type definitions
		"JSON": reflect.ValueOf((*json.JSON)(nil)),
	}
}
