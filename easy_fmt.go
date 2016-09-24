package easy_fmt

import (
	"fmt"
	"reflect"
)

func FormatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.IsNil() {
			return "nil"
		} else {
			return formatArrayValue(v)
		}
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		} else {
			return "&" + FormatValue(reflect.Indirect(v))
		}
	case reflect.Struct:
		return formatStructValue(v)
	default:
		return fmt.Sprintf("%#v", v.Interface())
	}
}

func EasyFormat(v interface{}) string {
	return FormatValue(reflect.ValueOf(v))
}

func formatArrayValue(vt reflect.Value) string {
	rt := vt.Type()
	valueType := rt.Elem()
	if valueType.Kind() != reflect.Ptr && valueType.Kind() != reflect.Struct {
		return fmt.Sprintf("%#v", vt.Interface())
	}
	res := fmt.Sprintf("[]%s{", valueType.String())
	for i := 0; i < vt.Len(); i++ {
		f := vt.Index(i)
		res += FormatValue(f) + ", "
	}
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	res += "}"
	return res
}

func formatStructValue(vt reflect.Value) string {
	rt := vt.Type()
	res := rt.Name() + "("
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" {
			continue
		}
		fv := vt.Field(i)
		fName := f.Name
		res = res + fmt.Sprintf("%s:%s, ", fName, FormatValue(fv))
	}
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	res += ")"
	return res
}
