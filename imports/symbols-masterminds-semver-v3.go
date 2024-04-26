// yaegi extract github.com/Masterminds/semver/v3

package imports

import (
	"reflect"

	"github.com/Masterminds/semver/v3"
)

func init() {
	Symbols["github.com/Masterminds/semver/v3/semver"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrEmptyString":       reflect.ValueOf(&semver.ErrEmptyString).Elem(),
		"ErrInvalidCharacters": reflect.ValueOf(&semver.ErrInvalidCharacters).Elem(),
		"ErrInvalidMetadata":   reflect.ValueOf(&semver.ErrInvalidMetadata).Elem(),
		"ErrInvalidPrerelease": reflect.ValueOf(&semver.ErrInvalidPrerelease).Elem(),
		"ErrInvalidSemVer":     reflect.ValueOf(&semver.ErrInvalidSemVer).Elem(),
		"ErrSegmentStartsZero": reflect.ValueOf(&semver.ErrSegmentStartsZero).Elem(),
		"MustParse":            reflect.ValueOf(semver.MustParse),
		"New":                  reflect.ValueOf(semver.New),
		"NewConstraint":        reflect.ValueOf(semver.NewConstraint),
		"NewVersion":           reflect.ValueOf(semver.NewVersion),
		"StrictNewVersion":     reflect.ValueOf(semver.StrictNewVersion),

		// type definitions
		"Collection":  reflect.ValueOf((*semver.Collection)(nil)),
		"Constraints": reflect.ValueOf((*semver.Constraints)(nil)),
		"Version":     reflect.ValueOf((*semver.Version)(nil)),
	}
}
