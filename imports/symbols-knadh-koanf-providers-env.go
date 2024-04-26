// yaegi extract github.com/knadh/koanf/providers/env

package imports

import (
	"reflect"

	"github.com/knadh/koanf/providers/env"
)

func init() {
	Symbols["github.com/knadh/koanf/providers/env/env"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Provider":          reflect.ValueOf(env.Provider),
		"ProviderWithValue": reflect.ValueOf(env.ProviderWithValue),

		// type definitions
		"Env": reflect.ValueOf((*env.Env)(nil)),
	}
}
