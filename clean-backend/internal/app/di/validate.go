package di

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrNilDependency = errors.New("nil dependency")

// isReflectValueNil は reflect.Value が実質 nil かを判定する。
// interface の場合は中身を再帰的にたどることで、
// interface に入った typed nil（例: var x IFace = (*Impl)(nil)）も検知する。
func isReflectValueNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			return true
		}
		return isReflectValueNil(v.Elem())
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	return isReflectValueNil(reflect.ValueOf(v))
}

func requireNonNil(name string, v any) error {
	if isNil(v) {
		return fmt.Errorf("%w: %s", ErrNilDependency, name)
	}
	return nil
}

// requireStructFieldsNonNil は構造体（または構造体ポインタ）を受け取り、
// 本体と全フィールドが nil でないことを検証する。
// nil のフィールドを見つけたら "name.FieldName" 形式のエラーを返す。
func requireStructFieldsNonNil(name string, v any) error {
	if err := requireNonNil(name, v); err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("invalid dependency container: %s", name)
	}

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		if !isReflectValueNil(fieldValue) {
			continue
		}
		return fmt.Errorf("%w: %s.%s", ErrNilDependency, name, rt.Field(i).Name)
	}
	return nil
}
