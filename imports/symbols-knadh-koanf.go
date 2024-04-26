// yaegi extract github.com/knadh/koanf/v2

package imports

import (
	"reflect"

	"github.com/knadh/koanf/v2"
)

func init() {
	Symbols["github.com/knadh/koanf/v2/koanf"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New":           reflect.ValueOf(koanf.New),
		"NewWithConf":   reflect.ValueOf(koanf.NewWithConf),
		"WithMergeFunc": reflect.ValueOf(koanf.WithMergeFunc),

		// type definitions
		"Conf":          reflect.ValueOf((*koanf.Conf)(nil)),
		"KeyMap":        reflect.ValueOf((*koanf.KeyMap)(nil)),
		"Koanf":         reflect.ValueOf((*koanf.Koanf)(nil)),
		"Option":        reflect.ValueOf((*koanf.Option)(nil)),
		"Parser":        reflect.ValueOf((*koanf.Parser)(nil)),
		"Provider":      reflect.ValueOf((*koanf.Provider)(nil)),
		"UnmarshalConf": reflect.ValueOf((*koanf.UnmarshalConf)(nil)),

		// interface wrapper definitions
		"_Parser":   reflect.ValueOf((*_github_com_knadh_koanf_v2_Parser)(nil)),
		"_Provider": reflect.ValueOf((*_github_com_knadh_koanf_v2_Provider)(nil)),
	}
}

// _github_com_knadh_koanf_v2_Parser is an interface wrapper for Parser type
type _github_com_knadh_koanf_v2_Parser struct {
	IValue     interface{}
	WMarshal   func(a0 map[string]interface{}) ([]byte, error)
	WUnmarshal func(a0 []byte) (map[string]interface{}, error)
}

func (W _github_com_knadh_koanf_v2_Parser) Marshal(a0 map[string]interface{}) ([]byte, error) {
	return W.WMarshal(a0)
}
func (W _github_com_knadh_koanf_v2_Parser) Unmarshal(a0 []byte) (map[string]interface{}, error) {
	return W.WUnmarshal(a0)
}

// _github_com_knadh_koanf_v2_Provider is an interface wrapper for Provider type
type _github_com_knadh_koanf_v2_Provider struct {
	IValue     interface{}
	WRead      func() (map[string]interface{}, error)
	WReadBytes func() ([]byte, error)
}

func (W _github_com_knadh_koanf_v2_Provider) Read() (map[string]interface{}, error) {
	return W.WRead()
}
func (W _github_com_knadh_koanf_v2_Provider) ReadBytes() ([]byte, error) {
	return W.WReadBytes()
}
