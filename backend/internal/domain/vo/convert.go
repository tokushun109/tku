package vo

import "strings"

// ToValuePtr は任意の ValueObject ポインタを、その基底値ポインタに変換する。
func ToValuePtr[T any, V ValueObject[T]](v *V) *T {
	if v == nil {
		return nil
	}

	value := (*v).Value()
	return &value
}

// ToValueOrEmpty は string を値にもつ ValueObject ポインタを string に変換する。
// nil が渡された場合は空文字を返す。
func ToValueOrEmpty[V ValueObject[string]](v *V) string {
	if v == nil {
		return ""
	}

	return (*v).Value()
}

// ParseOptionalValue は任意の生値を ValueObject ポインタに変換する。
// raw が nil の場合は nil として扱う。T が string のときは TrimSpace を適用し、
// 空白のみの入力は nil として扱う。
func ParseOptionalValue[T any, V ValueObject[T]](raw *T, parser func(T) (V, error)) (*V, error) {
	if raw == nil {
		return nil, nil
	}

	value := *raw
	if s, ok := any(value).(string); ok {
		trimmed := strings.TrimSpace(s)
		if trimmed == "" {
			return nil, nil
		}
		value = any(trimmed).(T)
	}

	parsed, err := parser(value)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
