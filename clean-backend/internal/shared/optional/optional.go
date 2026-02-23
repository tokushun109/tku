package optional

import "strings"

// ParseOptionalString は任意文字列入力をパースする。
// 空文字または空白のみの入力は nil として扱う。
func ParseOptionalString[T any](v string, parser func(string) (T, error)) (*T, error) {
	if strings.TrimSpace(v) == "" {
		return nil, nil
	}

	parsed, err := parser(v)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

// ToStringPtr は ~string 型ポインタを *string に変換する。
// nil が渡された場合は nil を返す。
func ToStringPtr[T ~string](value *T) *string {
	if value == nil {
		return nil
	}

	s := string(*value)
	return &s
}

// ToTrimmedStringOrEmpty は ~string 型ポインタを文字列に変換する。
// nil が渡された場合は空文字を返し、値がある場合は trim した文字列を返す。
func ToTrimmedStringOrEmpty[T ~string](value *T) string {
	if value == nil {
		return ""
	}

	return strings.TrimSpace(string(*value))
}
