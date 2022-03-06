package lombokString

import (
	"encoding/json"
	"strconv"
)

func (ls *LombokString) ParseAsJSON() string {
	lombok, _ := ls.recursiveParse(0, InfoFromPrevRecursionLayer{})
	res, err := json.MarshalIndent(extractOnlyJSONObject(lombok), "", "    ")
	if err != nil {
		panic(err)
	}
	return string(res)
}

func extractOnlyJSONObject(o interface{}) interface{} {
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
		m := make(map[string]interface{})
		for k, v := range *lo.tempMap {
			m[k] = extractOnlyJSONObject(v)
		}
		return m
	} else if lo.objType == "array" {
		a := make([]interface{}, len(*lo.tempArr))
		for i, ele := range *lo.tempArr {
			a[i] = extractOnlyJSONObject(ele)
		}
		return a
	}
	return nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
