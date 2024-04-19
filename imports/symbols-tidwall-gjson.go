// yaegi extract github.com/tidwall/gjson

package imports

import (
	"reflect"

	"github.com/tidwall/gjson"
)

func init() {
	Symbols["github.com/tidwall/gjson/gjson"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddModifier":      reflect.ValueOf(gjson.AddModifier),
		"AppendJSONString": reflect.ValueOf(gjson.AppendJSONString),
		"DisableModifiers": reflect.ValueOf(&gjson.DisableModifiers).Elem(),
		"Escape":           reflect.ValueOf(gjson.Escape),
		"False":            reflect.ValueOf(gjson.False),
		"ForEachLine":      reflect.ValueOf(gjson.ForEachLine),
		"Get":              reflect.ValueOf(gjson.Get),
		"GetBytes":         reflect.ValueOf(gjson.GetBytes),
		"GetMany":          reflect.ValueOf(gjson.GetMany),
		"GetManyBytes":     reflect.ValueOf(gjson.GetManyBytes),
		"JSON":             reflect.ValueOf(gjson.JSON),
		"ModifierExists":   reflect.ValueOf(gjson.ModifierExists),
		"Null":             reflect.ValueOf(gjson.Null),
		"Number":           reflect.ValueOf(gjson.Number),
		"Parse":            reflect.ValueOf(gjson.Parse),
		"ParseBytes":       reflect.ValueOf(gjson.ParseBytes),
		"String":           reflect.ValueOf(gjson.String),
		"True":             reflect.ValueOf(gjson.True),
		"Valid":            reflect.ValueOf(gjson.Valid),
		"ValidBytes":       reflect.ValueOf(gjson.ValidBytes),

		// type definitions
		"Result": reflect.ValueOf((*gjson.Result)(nil)),
		"Type":   reflect.ValueOf((*gjson.Type)(nil)),
	}
}
