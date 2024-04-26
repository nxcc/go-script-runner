// yaegi extract gopkg.in/yaml.v2

package imports

import (
	"reflect"

	"gopkg.in/yaml.v2"
)

func init() {
	Symbols["gopkg.in/yaml.v2/yaml"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FutureLineWrap":  reflect.ValueOf(yaml.FutureLineWrap),
		"Marshal":         reflect.ValueOf(yaml.Marshal),
		"NewDecoder":      reflect.ValueOf(yaml.NewDecoder),
		"NewEncoder":      reflect.ValueOf(yaml.NewEncoder),
		"Unmarshal":       reflect.ValueOf(yaml.Unmarshal),
		"UnmarshalStrict": reflect.ValueOf(yaml.UnmarshalStrict),

		// type definitions
		"Decoder":     reflect.ValueOf((*yaml.Decoder)(nil)),
		"Encoder":     reflect.ValueOf((*yaml.Encoder)(nil)),
		"IsZeroer":    reflect.ValueOf((*yaml.IsZeroer)(nil)),
		"MapItem":     reflect.ValueOf((*yaml.MapItem)(nil)),
		"MapSlice":    reflect.ValueOf((*yaml.MapSlice)(nil)),
		"Marshaler":   reflect.ValueOf((*yaml.Marshaler)(nil)),
		"TypeError":   reflect.ValueOf((*yaml.TypeError)(nil)),
		"Unmarshaler": reflect.ValueOf((*yaml.Unmarshaler)(nil)),

		// interface wrapper definitions
		"_IsZeroer":    reflect.ValueOf((*_gopkg_in_yaml_v2_IsZeroer)(nil)),
		"_Marshaler":   reflect.ValueOf((*_gopkg_in_yaml_v2_Marshaler)(nil)),
		"_Unmarshaler": reflect.ValueOf((*_gopkg_in_yaml_v2_Unmarshaler)(nil)),
	}
}

// _gopkg_in_yaml_v2_IsZeroer is an interface wrapper for IsZeroer type
type _gopkg_in_yaml_v2_IsZeroer struct {
	IValue  interface{}
	WIsZero func() bool
}

func (W _gopkg_in_yaml_v2_IsZeroer) IsZero() bool {
	return W.WIsZero()
}

// _gopkg_in_yaml_v2_Marshaler is an interface wrapper for Marshaler type
type _gopkg_in_yaml_v2_Marshaler struct {
	IValue       interface{}
	WMarshalYAML func() (interface{}, error)
}

func (W _gopkg_in_yaml_v2_Marshaler) MarshalYAML() (interface{}, error) {
	return W.WMarshalYAML()
}

// _gopkg_in_yaml_v2_Unmarshaler is an interface wrapper for Unmarshaler type
type _gopkg_in_yaml_v2_Unmarshaler struct {
	IValue         interface{}
	WUnmarshalYAML func(unmarshal func(interface{}) error) error
}

func (W _gopkg_in_yaml_v2_Unmarshaler) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return W.WUnmarshalYAML(unmarshal)
}
