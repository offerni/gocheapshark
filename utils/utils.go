package utils

import (
	"net/url"
	"reflect"
	"strconv"
)

func BuildQueryParams[T comparable](opts T) string {
	v := url.Values{}

	optsValue := reflect.ValueOf(opts)
	for i := 0; i < optsValue.NumField(); i++ {
		fieldValue := optsValue.Field(i)
		fieldType := optsValue.Type().Field(i)

		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			fieldName := fieldType.Tag.Get("json")
			var valueStr string
			switch fieldValue.Elem().Kind() {
			case reflect.Bool:
				valueStr = strconv.FormatBool(fieldValue.Elem().Bool())
			case reflect.Uint:
				valueStr = strconv.FormatUint(fieldValue.Elem().Uint(), 10)
			case reflect.String:
				valueStr = fieldValue.Elem().String()
			}
			v.Set(fieldName, valueStr)
		}
	}

	return v.Encode()
}

func ConvertNumericStringToUintPointer(s string) (*uint, error) {
	uint64Value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil, err
	}

	uintValue := uint(uint64Value)
	return &uintValue, nil
}
