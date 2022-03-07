package lombokString

import (
	"encoding/json"
	"strconv"
)

func (ls *LombokString) ParseAsJSON(args InterfaceArgs) string {
	lombok, _ := ls.recursiveParse(0, InfoFromPrevRecursionLayer{})
	jsonObject := extractOnlyJSONObject(lombok, args)

	var (
		res []byte
		err error
	)
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

func extractOnlyJSONObject(o interface{}, args InterfaceArgs) interface{} {
	if s, ok := o.(string); ok {
		if s == "null" {
			return nil
		}
		if isNumeric(s) {
			r, _ := strconv.ParseFloat(s, 64)
			return r
		}
		return s
	}

	lo, ok := o.(*LombokObject)
	if !ok {
		panic("there is an object that is neither string or LombokObject :/")
	}
	if lo.objType == "object" {
		m := handleObject(lo, args)
		return m
	} else if lo.objType == "array" {
		a := handleArray(lo, args)
		return a
	}
	return nil
}

func handleArray(lo *LombokObject, args InterfaceArgs) []interface{} {
	a := make([]interface{}, len(*lo.tempArr))
	for i, ele := range *lo.tempArr {
		extracted := extractOnlyJSONObject(ele, args)
		if args.ShouldExcludeNulls && extracted == nil {
			continue
		}
		a[i] = extracted
	}
	return a
}

func handleObject(lo *LombokObject, args InterfaceArgs) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range *lo.tempMap {
		extracted := extractOnlyJSONObject(v, args)
		if args.ShouldExcludeNulls && extracted == nil {
			continue
		}
		m[k] = extracted
	}
	return m
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
