// yaegi extract github.com/knadh/koanf/parsers/yaml

package imports

import (
	"reflect"

	"github.com/knadh/koanf/parsers/yaml"
)

func init() {
	Symbols["github.com/knadh/koanf/parsers/yaml/yaml"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Parser": reflect.ValueOf(yaml.Parser),

		// type definitions
		"YAML": reflect.ValueOf((*yaml.YAML)(nil)),
	}
}
