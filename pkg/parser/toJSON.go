package parser

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type omitThisField struct {
}

func ParseAsJSON(jsonObject interface{}, args InterfaceArgs) string {
	var (
		res []byte
		err error
	)

	jsonObject = forceParseNumericAndNull(jsonObject, args)

	if args.ShouldMinify {
		res, err = json.Marshal(jsonObject)
	} else {
		res, err = json.MarshalIndent(jsonObject, "", "    ")
	}
	if err != nil {
		panic(err)
	}
	return string(res)
}

func forceParseNumericAndNull(object interface{}, args InterfaceArgs) interface{} {
	if object == nil {
		panic("object to be parsed is actually nil")
	}
	s := reflect.ValueOf(object)
	if s.Kind() == reflect.Ptr || s.Kind() == reflect.Interface {
		s = s.Elem()
	}
	return recursivelyParseNumericAndNull(s, args)
}

func recursivelyParseNumericAndNull(s reflect.Value, args InterfaceArgs) interface{} {
	if s.Kind() == reflect.Ptr || s.Kind() == reflect.Interface {
		s = s.Elem()
	}
	switch s.Kind() {
	case reflect.Map:
		mapp := make(map[string]interface{})
		for _, value := range s.MapKeys() {
			v := recursivelyParseNumericAndNull(s.MapIndex(value), args)
			if _, ok := v.(omitThisField); !ok {
				mapp[value.String()] = v
			}
		}
		return mapp
	case reflect.Slice, reflect.Array:
		slice := make([]interface{}, 0)
		for i := 0; i < s.Len(); i++ {
			v := recursivelyParseNumericAndNull(s.Index(i), args)
			if _, ok := v.(omitThisField); !ok {
				slice = append(slice, v)
			}
		}
		return slice
	case reflect.String:
		if s.String() == "null" {
			if args.ShouldExcludeNulls {
				return omitThisField{}
			}
			return nil
		} else if v, ok := isNumeric(s.String()); ok {
			return v
		} else {
			return s.String()
		}
	default:
		panic("not implemented for this type")
	}
}

func isNumeric(s string) (value float64, ok bool) {
	value, err := strconv.ParseFloat(s, 64)
	return value, err == nil
}
