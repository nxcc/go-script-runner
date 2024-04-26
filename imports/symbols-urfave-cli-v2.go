// yaegi extract github.com/urfave/cli/v2

package imports

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/urfave/cli/v2"
)

func init() {
	Symbols["github.com/urfave/cli/v2/cli"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AppHelpTemplate":           reflect.ValueOf(&cli.AppHelpTemplate).Elem(),
		"BashCompletionFlag":        reflect.ValueOf(&cli.BashCompletionFlag).Elem(),
		"CommandHelpTemplate":       reflect.ValueOf(&cli.CommandHelpTemplate).Elem(),
		"DefaultAppComplete":        reflect.ValueOf(cli.DefaultAppComplete),
		"DefaultCompleteWithFlags":  reflect.ValueOf(cli.DefaultCompleteWithFlags),
		"ErrWriter":                 reflect.ValueOf(&cli.ErrWriter).Elem(),
		"Exit":                      reflect.ValueOf(cli.Exit),
		"FishCompletionTemplate":    reflect.ValueOf(&cli.FishCompletionTemplate).Elem(),
		"FlagEnvHinter":             reflect.ValueOf(&cli.FlagEnvHinter).Elem(),
		"FlagFileHinter":            reflect.ValueOf(&cli.FlagFileHinter).Elem(),
		"FlagNamePrefixer":          reflect.ValueOf(&cli.FlagNamePrefixer).Elem(),
		"FlagNames":                 reflect.ValueOf(cli.FlagNames),
		"FlagStringer":              reflect.ValueOf(&cli.FlagStringer).Elem(),
		"HandleAction":              reflect.ValueOf(cli.HandleAction),
		"HandleExitCoder":           reflect.ValueOf(cli.HandleExitCoder),
		"HelpFlag":                  reflect.ValueOf(&cli.HelpFlag).Elem(),
		"HelpPrinter":               reflect.ValueOf(&cli.HelpPrinter).Elem(),
		"HelpPrinterCustom":         reflect.ValueOf(&cli.HelpPrinterCustom).Elem(),
		"MarkdownDocTemplate":       reflect.ValueOf(&cli.MarkdownDocTemplate).Elem(),
		"NewApp":                    reflect.ValueOf(cli.NewApp),
		"NewContext":                reflect.ValueOf(cli.NewContext),
		"NewExitError":              reflect.ValueOf(cli.NewExitError),
		"NewFloat64Slice":           reflect.ValueOf(cli.NewFloat64Slice),
		"NewInt64Slice":             reflect.ValueOf(cli.NewInt64Slice),
		"NewIntSlice":               reflect.ValueOf(cli.NewIntSlice),
		"NewStringSlice":            reflect.ValueOf(cli.NewStringSlice),
		"NewTimestamp":              reflect.ValueOf(cli.NewTimestamp),
		"NewUint64Slice":            reflect.ValueOf(cli.NewUint64Slice),
		"NewUintSlice":              reflect.ValueOf(cli.NewUintSlice),
		"OsExiter":                  reflect.ValueOf(&cli.OsExiter).Elem(),
		"ShowAppHelp":               reflect.ValueOf(cli.ShowAppHelp),
		"ShowAppHelpAndExit":        reflect.ValueOf(cli.ShowAppHelpAndExit),
		"ShowCommandCompletions":    reflect.ValueOf(cli.ShowCommandCompletions),
		"ShowCommandHelp":           reflect.ValueOf(cli.ShowCommandHelp),
		"ShowCommandHelpAndExit":    reflect.ValueOf(cli.ShowCommandHelpAndExit),
		"ShowCompletions":           reflect.ValueOf(cli.ShowCompletions),
		"ShowSubcommandHelp":        reflect.ValueOf(cli.ShowSubcommandHelp),
		"ShowSubcommandHelpAndExit": reflect.ValueOf(cli.ShowSubcommandHelpAndExit),
		"ShowVersion":               reflect.ValueOf(cli.ShowVersion),
		"SubcommandHelpTemplate":    reflect.ValueOf(&cli.SubcommandHelpTemplate).Elem(),
		"SuggestCommand":            reflect.ValueOf(&cli.SuggestCommand).Elem(),
		"SuggestDidYouMeanTemplate": reflect.ValueOf(&cli.SuggestDidYouMeanTemplate).Elem(),
		"SuggestFlag":               reflect.ValueOf(&cli.SuggestFlag).Elem(),
		"VersionFlag":               reflect.ValueOf(&cli.VersionFlag).Elem(),
		"VersionPrinter":            reflect.ValueOf(&cli.VersionPrinter).Elem(),

		// type definitions
		"ActionFunc":             reflect.ValueOf((*cli.ActionFunc)(nil)),
		"ActionableFlag":         reflect.ValueOf((*cli.ActionableFlag)(nil)),
		"AfterFunc":              reflect.ValueOf((*cli.AfterFunc)(nil)),
		"App":                    reflect.ValueOf((*cli.App)(nil)),
		"Args":                   reflect.ValueOf((*cli.Args)(nil)),
		"Author":                 reflect.ValueOf((*cli.Author)(nil)),
		"BashCompleteFunc":       reflect.ValueOf((*cli.BashCompleteFunc)(nil)),
		"BeforeFunc":             reflect.ValueOf((*cli.BeforeFunc)(nil)),
		"BoolFlag":               reflect.ValueOf((*cli.BoolFlag)(nil)),
		"CategorizableFlag":      reflect.ValueOf((*cli.CategorizableFlag)(nil)),
		"Command":                reflect.ValueOf((*cli.Command)(nil)),
		"CommandCategories":      reflect.ValueOf((*cli.CommandCategories)(nil)),
		"CommandCategory":        reflect.ValueOf((*cli.CommandCategory)(nil)),
		"CommandNotFoundFunc":    reflect.ValueOf((*cli.CommandNotFoundFunc)(nil)),
		"Commands":               reflect.ValueOf((*cli.Commands)(nil)),
		"CommandsByName":         reflect.ValueOf((*cli.CommandsByName)(nil)),
		"Context":                reflect.ValueOf((*cli.Context)(nil)),
		"Countable":              reflect.ValueOf((*cli.Countable)(nil)),
		"DocGenerationFlag":      reflect.ValueOf((*cli.DocGenerationFlag)(nil)),
		"DocGenerationSliceFlag": reflect.ValueOf((*cli.DocGenerationSliceFlag)(nil)),
		"DurationFlag":           reflect.ValueOf((*cli.DurationFlag)(nil)),
		"ErrorFormatter":         reflect.ValueOf((*cli.ErrorFormatter)(nil)),
		"ExitCoder":              reflect.ValueOf((*cli.ExitCoder)(nil)),
		"ExitErrHandlerFunc":     reflect.ValueOf((*cli.ExitErrHandlerFunc)(nil)),
		"Flag":                   reflect.ValueOf((*cli.Flag)(nil)),
		"FlagCategories":         reflect.ValueOf((*cli.FlagCategories)(nil)),
		"FlagEnvHintFunc":        reflect.ValueOf((*cli.FlagEnvHintFunc)(nil)),
		"FlagFileHintFunc":       reflect.ValueOf((*cli.FlagFileHintFunc)(nil)),
		"FlagNamePrefixFunc":     reflect.ValueOf((*cli.FlagNamePrefixFunc)(nil)),
		"FlagStringFunc":         reflect.ValueOf((*cli.FlagStringFunc)(nil)),
		"FlagsByName":            reflect.ValueOf((*cli.FlagsByName)(nil)),
		"Float64Flag":            reflect.ValueOf((*cli.Float64Flag)(nil)),
		"Float64Slice":           reflect.ValueOf((*cli.Float64Slice)(nil)),
		"Float64SliceFlag":       reflect.ValueOf((*cli.Float64SliceFlag)(nil)),
		"Generic":                reflect.ValueOf((*cli.Generic)(nil)),
		"GenericFlag":            reflect.ValueOf((*cli.GenericFlag)(nil)),
		"Int64Flag":              reflect.ValueOf((*cli.Int64Flag)(nil)),
		"Int64Slice":             reflect.ValueOf((*cli.Int64Slice)(nil)),
		"Int64SliceFlag":         reflect.ValueOf((*cli.Int64SliceFlag)(nil)),
		"IntFlag":                reflect.ValueOf((*cli.IntFlag)(nil)),
		"IntSlice":               reflect.ValueOf((*cli.IntSlice)(nil)),
		"IntSliceFlag":           reflect.ValueOf((*cli.IntSliceFlag)(nil)),
		"InvalidFlagAccessFunc":  reflect.ValueOf((*cli.InvalidFlagAccessFunc)(nil)),
		"MultiError":             reflect.ValueOf((*cli.MultiError)(nil)),
		"OnUsageErrorFunc":       reflect.ValueOf((*cli.OnUsageErrorFunc)(nil)),
		"Path":                   reflect.ValueOf((*cli.Path)(nil)),
		"PathFlag":               reflect.ValueOf((*cli.PathFlag)(nil)),
		"RequiredFlag":           reflect.ValueOf((*cli.RequiredFlag)(nil)),
		"Serializer":             reflect.ValueOf((*cli.Serializer)(nil)),
		"StringFlag":             reflect.ValueOf((*cli.StringFlag)(nil)),
		"StringSlice":            reflect.ValueOf((*cli.StringSlice)(nil)),
		"StringSliceFlag":        reflect.ValueOf((*cli.StringSliceFlag)(nil)),
		"SuggestCommandFunc":     reflect.ValueOf((*cli.SuggestCommandFunc)(nil)),
		"SuggestFlagFunc":        reflect.ValueOf((*cli.SuggestFlagFunc)(nil)),
		"Timestamp":              reflect.ValueOf((*cli.Timestamp)(nil)),
		"TimestampFlag":          reflect.ValueOf((*cli.TimestampFlag)(nil)),
		"Uint64Flag":             reflect.ValueOf((*cli.Uint64Flag)(nil)),
		"Uint64Slice":            reflect.ValueOf((*cli.Uint64Slice)(nil)),
		"Uint64SliceFlag":        reflect.ValueOf((*cli.Uint64SliceFlag)(nil)),
		"UintFlag":               reflect.ValueOf((*cli.UintFlag)(nil)),
		"UintSlice":              reflect.ValueOf((*cli.UintSlice)(nil)),
		"UintSliceFlag":          reflect.ValueOf((*cli.UintSliceFlag)(nil)),
		"VisibleFlag":            reflect.ValueOf((*cli.VisibleFlag)(nil)),
		"VisibleFlagCategory":    reflect.ValueOf((*cli.VisibleFlagCategory)(nil)),

		// interface wrapper definitions
		"_ActionableFlag":         reflect.ValueOf((*_github_com_urfave_cli_v2_ActionableFlag)(nil)),
		"_Args":                   reflect.ValueOf((*_github_com_urfave_cli_v2_Args)(nil)),
		"_CategorizableFlag":      reflect.ValueOf((*_github_com_urfave_cli_v2_CategorizableFlag)(nil)),
		"_CommandCategories":      reflect.ValueOf((*_github_com_urfave_cli_v2_CommandCategories)(nil)),
		"_CommandCategory":        reflect.ValueOf((*_github_com_urfave_cli_v2_CommandCategory)(nil)),
		"_Countable":              reflect.ValueOf((*_github_com_urfave_cli_v2_Countable)(nil)),
		"_DocGenerationFlag":      reflect.ValueOf((*_github_com_urfave_cli_v2_DocGenerationFlag)(nil)),
		"_DocGenerationSliceFlag": reflect.ValueOf((*_github_com_urfave_cli_v2_DocGenerationSliceFlag)(nil)),
		"_ErrorFormatter":         reflect.ValueOf((*_github_com_urfave_cli_v2_ErrorFormatter)(nil)),
		"_ExitCoder":              reflect.ValueOf((*_github_com_urfave_cli_v2_ExitCoder)(nil)),
		"_Flag":                   reflect.ValueOf((*_github_com_urfave_cli_v2_Flag)(nil)),
		"_FlagCategories":         reflect.ValueOf((*_github_com_urfave_cli_v2_FlagCategories)(nil)),
		"_Generic":                reflect.ValueOf((*_github_com_urfave_cli_v2_Generic)(nil)),
		"_MultiError":             reflect.ValueOf((*_github_com_urfave_cli_v2_MultiError)(nil)),
		"_RequiredFlag":           reflect.ValueOf((*_github_com_urfave_cli_v2_RequiredFlag)(nil)),
		"_Serializer":             reflect.ValueOf((*_github_com_urfave_cli_v2_Serializer)(nil)),
		"_VisibleFlag":            reflect.ValueOf((*_github_com_urfave_cli_v2_VisibleFlag)(nil)),
		"_VisibleFlagCategory":    reflect.ValueOf((*_github_com_urfave_cli_v2_VisibleFlagCategory)(nil)),
	}
}

// _github_com_urfave_cli_v2_ActionableFlag is an interface wrapper for ActionableFlag type
type _github_com_urfave_cli_v2_ActionableFlag struct {
	IValue     interface{}
	WApply     func(a0 *flag.FlagSet) error
	WIsSet     func() bool
	WNames     func() []string
	WRunAction func(a0 *cli.Context) error
	WString    func() string
}

func (W _github_com_urfave_cli_v2_ActionableFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_ActionableFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_ActionableFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_ActionableFlag) RunAction(a0 *cli.Context) error {
	return W.WRunAction(a0)
}
func (W _github_com_urfave_cli_v2_ActionableFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_Args is an interface wrapper for Args type
type _github_com_urfave_cli_v2_Args struct {
	IValue   interface{}
	WFirst   func() string
	WGet     func(n int) string
	WLen     func() int
	WPresent func() bool
	WSlice   func() []string
	WTail    func() []string
}

func (W _github_com_urfave_cli_v2_Args) First() string {
	return W.WFirst()
}
func (W _github_com_urfave_cli_v2_Args) Get(n int) string {
	return W.WGet(n)
}
func (W _github_com_urfave_cli_v2_Args) Len() int {
	return W.WLen()
}
func (W _github_com_urfave_cli_v2_Args) Present() bool {
	return W.WPresent()
}
func (W _github_com_urfave_cli_v2_Args) Slice() []string {
	return W.WSlice()
}
func (W _github_com_urfave_cli_v2_Args) Tail() []string {
	return W.WTail()
}

// _github_com_urfave_cli_v2_CategorizableFlag is an interface wrapper for CategorizableFlag type
type _github_com_urfave_cli_v2_CategorizableFlag struct {
	IValue       interface{}
	WApply       func(a0 *flag.FlagSet) error
	WGetCategory func() string
	WIsSet       func() bool
	WIsVisible   func() bool
	WNames       func() []string
	WString      func() string
}

func (W _github_com_urfave_cli_v2_CategorizableFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_CategorizableFlag) GetCategory() string {
	return W.WGetCategory()
}
func (W _github_com_urfave_cli_v2_CategorizableFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_CategorizableFlag) IsVisible() bool {
	return W.WIsVisible()
}
func (W _github_com_urfave_cli_v2_CategorizableFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_CategorizableFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_CommandCategories is an interface wrapper for CommandCategories type
type _github_com_urfave_cli_v2_CommandCategories struct {
	IValue      interface{}
	WAddCommand func(category string, command *cli.Command)
	WCategories func() []cli.CommandCategory
}

func (W _github_com_urfave_cli_v2_CommandCategories) AddCommand(category string, command *cli.Command) {
	W.WAddCommand(category, command)
}
func (W _github_com_urfave_cli_v2_CommandCategories) Categories() []cli.CommandCategory {
	return W.WCategories()
}

// _github_com_urfave_cli_v2_CommandCategory is an interface wrapper for CommandCategory type
type _github_com_urfave_cli_v2_CommandCategory struct {
	IValue           interface{}
	WName            func() string
	WVisibleCommands func() []*cli.Command
}

func (W _github_com_urfave_cli_v2_CommandCategory) Name() string {
	return W.WName()
}
func (W _github_com_urfave_cli_v2_CommandCategory) VisibleCommands() []*cli.Command {
	return W.WVisibleCommands()
}

// _github_com_urfave_cli_v2_Countable is an interface wrapper for Countable type
type _github_com_urfave_cli_v2_Countable struct {
	IValue interface{}
	WCount func() int
}

func (W _github_com_urfave_cli_v2_Countable) Count() int {
	return W.WCount()
}

// _github_com_urfave_cli_v2_DocGenerationFlag is an interface wrapper for DocGenerationFlag type
type _github_com_urfave_cli_v2_DocGenerationFlag struct {
	IValue          interface{}
	WApply          func(a0 *flag.FlagSet) error
	WGetDefaultText func() string
	WGetEnvVars     func() []string
	WGetUsage       func() string
	WGetValue       func() string
	WIsSet          func() bool
	WNames          func() []string
	WString         func() string
	WTakesValue     func() bool
}

func (W _github_com_urfave_cli_v2_DocGenerationFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) GetDefaultText() string {
	return W.WGetDefaultText()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) GetEnvVars() []string {
	return W.WGetEnvVars()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) GetUsage() string {
	return W.WGetUsage()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) GetValue() string {
	return W.WGetValue()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _github_com_urfave_cli_v2_DocGenerationFlag) TakesValue() bool {
	return W.WTakesValue()
}

// _github_com_urfave_cli_v2_DocGenerationSliceFlag is an interface wrapper for DocGenerationSliceFlag type
type _github_com_urfave_cli_v2_DocGenerationSliceFlag struct {
	IValue          interface{}
	WApply          func(a0 *flag.FlagSet) error
	WGetDefaultText func() string
	WGetEnvVars     func() []string
	WGetUsage       func() string
	WGetValue       func() string
	WIsSet          func() bool
	WIsSliceFlag    func() bool
	WNames          func() []string
	WString         func() string
	WTakesValue     func() bool
}

func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) GetDefaultText() string {
	return W.WGetDefaultText()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) GetEnvVars() []string {
	return W.WGetEnvVars()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) GetUsage() string {
	return W.WGetUsage()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) GetValue() string {
	return W.WGetValue()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) IsSliceFlag() bool {
	return W.WIsSliceFlag()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _github_com_urfave_cli_v2_DocGenerationSliceFlag) TakesValue() bool {
	return W.WTakesValue()
}

// _github_com_urfave_cli_v2_ErrorFormatter is an interface wrapper for ErrorFormatter type
type _github_com_urfave_cli_v2_ErrorFormatter struct {
	IValue  interface{}
	WFormat func(s fmt.State, verb rune)
}

func (W _github_com_urfave_cli_v2_ErrorFormatter) Format(s fmt.State, verb rune) {
	W.WFormat(s, verb)
}

// _github_com_urfave_cli_v2_ExitCoder is an interface wrapper for ExitCoder type
type _github_com_urfave_cli_v2_ExitCoder struct {
	IValue    interface{}
	WError    func() string
	WExitCode func() int
}

func (W _github_com_urfave_cli_v2_ExitCoder) Error() string {
	return W.WError()
}
func (W _github_com_urfave_cli_v2_ExitCoder) ExitCode() int {
	return W.WExitCode()
}

// _github_com_urfave_cli_v2_Flag is an interface wrapper for Flag type
type _github_com_urfave_cli_v2_Flag struct {
	IValue  interface{}
	WApply  func(a0 *flag.FlagSet) error
	WIsSet  func() bool
	WNames  func() []string
	WString func() string
}

func (W _github_com_urfave_cli_v2_Flag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_Flag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_Flag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_Flag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_FlagCategories is an interface wrapper for FlagCategories type
type _github_com_urfave_cli_v2_FlagCategories struct {
	IValue             interface{}
	WAddFlag           func(category string, fl cli.Flag)
	WVisibleCategories func() []cli.VisibleFlagCategory
}

func (W _github_com_urfave_cli_v2_FlagCategories) AddFlag(category string, fl cli.Flag) {
	W.WAddFlag(category, fl)
}
func (W _github_com_urfave_cli_v2_FlagCategories) VisibleCategories() []cli.VisibleFlagCategory {
	return W.WVisibleCategories()
}

// _github_com_urfave_cli_v2_Generic is an interface wrapper for Generic type
type _github_com_urfave_cli_v2_Generic struct {
	IValue  interface{}
	WSet    func(value string) error
	WString func() string
}

func (W _github_com_urfave_cli_v2_Generic) Set(value string) error {
	return W.WSet(value)
}
func (W _github_com_urfave_cli_v2_Generic) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_MultiError is an interface wrapper for MultiError type
type _github_com_urfave_cli_v2_MultiError struct {
	IValue  interface{}
	WError  func() string
	WErrors func() []error
}

func (W _github_com_urfave_cli_v2_MultiError) Error() string {
	return W.WError()
}
func (W _github_com_urfave_cli_v2_MultiError) Errors() []error {
	return W.WErrors()
}

// _github_com_urfave_cli_v2_RequiredFlag is an interface wrapper for RequiredFlag type
type _github_com_urfave_cli_v2_RequiredFlag struct {
	IValue      interface{}
	WApply      func(a0 *flag.FlagSet) error
	WIsRequired func() bool
	WIsSet      func() bool
	WNames      func() []string
	WString     func() string
}

func (W _github_com_urfave_cli_v2_RequiredFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_RequiredFlag) IsRequired() bool {
	return W.WIsRequired()
}
func (W _github_com_urfave_cli_v2_RequiredFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_RequiredFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_RequiredFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_Serializer is an interface wrapper for Serializer type
type _github_com_urfave_cli_v2_Serializer struct {
	IValue     interface{}
	WSerialize func() string
}

func (W _github_com_urfave_cli_v2_Serializer) Serialize() string {
	return W.WSerialize()
}

// _github_com_urfave_cli_v2_VisibleFlag is an interface wrapper for VisibleFlag type
type _github_com_urfave_cli_v2_VisibleFlag struct {
	IValue     interface{}
	WApply     func(a0 *flag.FlagSet) error
	WIsSet     func() bool
	WIsVisible func() bool
	WNames     func() []string
	WString    func() string
}

func (W _github_com_urfave_cli_v2_VisibleFlag) Apply(a0 *flag.FlagSet) error {
	return W.WApply(a0)
}
func (W _github_com_urfave_cli_v2_VisibleFlag) IsSet() bool {
	return W.WIsSet()
}
func (W _github_com_urfave_cli_v2_VisibleFlag) IsVisible() bool {
	return W.WIsVisible()
}
func (W _github_com_urfave_cli_v2_VisibleFlag) Names() []string {
	return W.WNames()
}
func (W _github_com_urfave_cli_v2_VisibleFlag) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}

// _github_com_urfave_cli_v2_VisibleFlagCategory is an interface wrapper for VisibleFlagCategory type
type _github_com_urfave_cli_v2_VisibleFlagCategory struct {
	IValue interface{}
	WFlags func() []cli.VisibleFlag
	WName  func() string
}

func (W _github_com_urfave_cli_v2_VisibleFlagCategory) Flags() []cli.VisibleFlag {
	return W.WFlags()
}
func (W _github_com_urfave_cli_v2_VisibleFlagCategory) Name() string {
	return W.WName()
}
