package reflectutils

import (
	"fmt"
	"reflect"
)

const (
	TagEnv = "env"
)

func StructToEnv(v any) []string {
	metas := readStructMetadata(v)

	envs := make([]string, 0, len(metas))
	for _, meta := range metas {
		if meta.isZeroValue() {
			continue
		}

		envs = append(envs, fmt.Sprintf("%s=%s", meta.env, meta.fieldValueToString()))
	}

	return envs
}

type structMeta struct {
	env        string
	fieldName  string
	fieldValue reflect.Value
}

func (sm *structMeta) isZeroValue() bool {
	return sm.fieldValue.IsZero()
}

func (sm *structMeta) fieldValueToString() string {
	return fmt.Sprint(sm.fieldValue.Interface())
}

func readStructMetadata(structure any) []structMeta {
	type structNode struct {
		Val interface{}
	}

	structStack := []structNode{{Val: structure}}
	metas := make([]structMeta, 0)
	_ = metas

	for i := 0; i < len(structStack); i++ {

		s := reflect.ValueOf(structStack[i].Val)
		if s.Kind() == reflect.Pointer || s.Kind() == reflect.Interface {
			s = s.Elem()
		}

		if s.Kind() != reflect.Struct {
			continue
		}
		typeInfo := s.Type()

		for j := 0; j < s.NumField(); j++ {
			fType := typeInfo.Field(j)

			if fld := s.Field(j); fld.Kind() == reflect.Struct {
				if fld.CanInterface() {
					structStack = append(structStack, structNode{Val: fld.Interface()})
				}

				continue
			}

			if value, ok := fType.Tag.Lookup(TagEnv); ok {
				metas = append(metas, structMeta{
					env:        value,
					fieldName:  fType.Name,
					fieldValue: reflect.ValueOf(s.Field(j)),
				})
			}
		}
	}

	return metas
}
