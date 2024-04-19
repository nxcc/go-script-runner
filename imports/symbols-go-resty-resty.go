// yaegi extract github.com/go-resty/resty/v2

package imports

import (
	"go/constant"
	"go/token"
	"net/http"
	"reflect"

	"github.com/go-resty/resty/v2"
)

func init() {
	Symbols["github.com/go-resty/resty/v2/resty"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Backoff":                   reflect.ValueOf(resty.Backoff),
		"DetectContentType":         reflect.ValueOf(resty.DetectContentType),
		"DomainCheckRedirectPolicy": reflect.ValueOf(resty.DomainCheckRedirectPolicy),
		"ErrAutoRedirectDisabled":   reflect.ValueOf(&resty.ErrAutoRedirectDisabled).Elem(),
		"ErrDigestAlgNotSupported":  reflect.ValueOf(&resty.ErrDigestAlgNotSupported).Elem(),
		"ErrDigestBadChallenge":     reflect.ValueOf(&resty.ErrDigestBadChallenge).Elem(),
		"ErrDigestCharset":          reflect.ValueOf(&resty.ErrDigestCharset).Elem(),
		"ErrDigestNoQop":            reflect.ValueOf(&resty.ErrDigestNoQop).Elem(),
		"ErrDigestQopNotSupported":  reflect.ValueOf(&resty.ErrDigestQopNotSupported).Elem(),
		"ErrRateLimitExceeded":      reflect.ValueOf(&resty.ErrRateLimitExceeded).Elem(),
		"FlexibleRedirectPolicy":    reflect.ValueOf(resty.FlexibleRedirectPolicy),
		"IsJSONType":                reflect.ValueOf(resty.IsJSONType),
		"IsStringEmpty":             reflect.ValueOf(resty.IsStringEmpty),
		"IsXMLType":                 reflect.ValueOf(resty.IsXMLType),
		"MaxWaitTime":               reflect.ValueOf(resty.MaxWaitTime),
		"MethodDelete":              reflect.ValueOf(constant.MakeFromLiteral("\"DELETE\"", token.STRING, 0)),
		"MethodGet":                 reflect.ValueOf(constant.MakeFromLiteral("\"GET\"", token.STRING, 0)),
		"MethodHead":                reflect.ValueOf(constant.MakeFromLiteral("\"HEAD\"", token.STRING, 0)),
		"MethodOptions":             reflect.ValueOf(constant.MakeFromLiteral("\"OPTIONS\"", token.STRING, 0)),
		"MethodPatch":               reflect.ValueOf(constant.MakeFromLiteral("\"PATCH\"", token.STRING, 0)),
		"MethodPost":                reflect.ValueOf(constant.MakeFromLiteral("\"POST\"", token.STRING, 0)),
		"MethodPut":                 reflect.ValueOf(constant.MakeFromLiteral("\"PUT\"", token.STRING, 0)),
		"New":                       reflect.ValueOf(resty.New),
		"NewWithClient":             reflect.ValueOf(resty.NewWithClient),
		"NewWithLocalAddr":          reflect.ValueOf(resty.NewWithLocalAddr),
		"NoRedirectPolicy":          reflect.ValueOf(resty.NoRedirectPolicy),
		"ResetMultipartReaders":     reflect.ValueOf(resty.ResetMultipartReaders),
		"Retries":                   reflect.ValueOf(resty.Retries),
		"RetryConditions":           reflect.ValueOf(resty.RetryConditions),
		"RetryHooks":                reflect.ValueOf(resty.RetryHooks),
		"Unmarshalc":                reflect.ValueOf(resty.Unmarshalc),
		"Version":                   reflect.ValueOf(constant.MakeFromLiteral("\"2.12.0\"", token.STRING, 0)),
		"WaitTime":                  reflect.ValueOf(resty.WaitTime),

		// type definitions
		"Client":              reflect.ValueOf((*resty.Client)(nil)),
		"ErrorHook":           reflect.ValueOf((*resty.ErrorHook)(nil)),
		"File":                reflect.ValueOf((*resty.File)(nil)),
		"Logger":              reflect.ValueOf((*resty.Logger)(nil)),
		"MultipartField":      reflect.ValueOf((*resty.MultipartField)(nil)),
		"OnRetryFunc":         reflect.ValueOf((*resty.OnRetryFunc)(nil)),
		"Option":              reflect.ValueOf((*resty.Option)(nil)),
		"Options":             reflect.ValueOf((*resty.Options)(nil)),
		"PreRequestHook":      reflect.ValueOf((*resty.PreRequestHook)(nil)),
		"RateLimiter":         reflect.ValueOf((*resty.RateLimiter)(nil)),
		"RedirectPolicy":      reflect.ValueOf((*resty.RedirectPolicy)(nil)),
		"RedirectPolicyFunc":  reflect.ValueOf((*resty.RedirectPolicyFunc)(nil)),
		"Request":             reflect.ValueOf((*resty.Request)(nil)),
		"RequestLog":          reflect.ValueOf((*resty.RequestLog)(nil)),
		"RequestLogCallback":  reflect.ValueOf((*resty.RequestLogCallback)(nil)),
		"RequestMiddleware":   reflect.ValueOf((*resty.RequestMiddleware)(nil)),
		"Response":            reflect.ValueOf((*resty.Response)(nil)),
		"ResponseError":       reflect.ValueOf((*resty.ResponseError)(nil)),
		"ResponseLog":         reflect.ValueOf((*resty.ResponseLog)(nil)),
		"ResponseLogCallback": reflect.ValueOf((*resty.ResponseLogCallback)(nil)),
		"ResponseMiddleware":  reflect.ValueOf((*resty.ResponseMiddleware)(nil)),
		"RetryAfterFunc":      reflect.ValueOf((*resty.RetryAfterFunc)(nil)),
		"RetryConditionFunc":  reflect.ValueOf((*resty.RetryConditionFunc)(nil)),
		"SRVRecord":           reflect.ValueOf((*resty.SRVRecord)(nil)),
		"SuccessHook":         reflect.ValueOf((*resty.SuccessHook)(nil)),
		"TraceInfo":           reflect.ValueOf((*resty.TraceInfo)(nil)),
		"User":                reflect.ValueOf((*resty.User)(nil)),

		// interface wrapper definitions
		"_Logger":         reflect.ValueOf((*_github_com_go_resty_resty_v2_Logger)(nil)),
		"_RateLimiter":    reflect.ValueOf((*_github_com_go_resty_resty_v2_RateLimiter)(nil)),
		"_RedirectPolicy": reflect.ValueOf((*_github_com_go_resty_resty_v2_RedirectPolicy)(nil)),
	}
}

// _github_com_go_resty_resty_v2_Logger is an interface wrapper for Logger type
type _github_com_go_resty_resty_v2_Logger struct {
	IValue  interface{}
	WDebugf func(format string, v ...interface{})
	WErrorf func(format string, v ...interface{})
	WWarnf  func(format string, v ...interface{})
}

func (W _github_com_go_resty_resty_v2_Logger) Debugf(format string, v ...interface{}) {
	W.WDebugf(format, v...)
}
func (W _github_com_go_resty_resty_v2_Logger) Errorf(format string, v ...interface{}) {
	W.WErrorf(format, v...)
}
func (W _github_com_go_resty_resty_v2_Logger) Warnf(format string, v ...interface{}) {
	W.WWarnf(format, v...)
}

// _github_com_go_resty_resty_v2_RateLimiter is an interface wrapper for RateLimiter type
type _github_com_go_resty_resty_v2_RateLimiter struct {
	IValue interface{}
	WAllow func() bool
}

func (W _github_com_go_resty_resty_v2_RateLimiter) Allow() bool {
	return W.WAllow()
}

// _github_com_go_resty_resty_v2_RedirectPolicy is an interface wrapper for RedirectPolicy type
type _github_com_go_resty_resty_v2_RedirectPolicy struct {
	IValue interface{}
	WApply func(req *http.Request, via []*http.Request) error
}

func (W _github_com_go_resty_resty_v2_RedirectPolicy) Apply(req *http.Request, via []*http.Request) error {
	return W.WApply(req, via)
}
