package result

type Result[T any] struct {
	OK bool `json:"ok"`
	Value T `json:"value,omitempty"`
	Code string `json:"code,omitempty"`
	Msg string `json:"message,omitempty"`
}

func Ok[T any](v T) Result[T] { return Result[T]{OK: true, Value: v} }
func Err[T any](code, msg string) Result[T] { return Result[T]{OK: false, Code: code, Msg: msg} }
