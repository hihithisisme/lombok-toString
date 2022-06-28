package parser

import (
	"fmt"
	"strings"
)

const SPECIAL_CHARS = "()[]=,"

func SegmentByBrackets(in string) (elements []string) {
	pointer := 0
	// segment by brackets
	for i, char := range in {
		if char == '[' || char == '(' {
			elements = append(elements, in[pointer:i])
			pointer = i
		} else if char == ']' || char == ')' {
			elements = append(elements, in[pointer:i+1])
			pointer = i + 1
		}
	}

	// handle case of being wrapped by brackets
	if elements[0] == "" {
		elements = elements[1:]
	}

	elements = pushLeadingCommasToPreviousElement(elements)
	elements = pullClassNamesFromPreviousElement(elements)

	return elements
}

func Process(elements []string) interface{} {
	res, _ := process(elements, 0)
	return res
}

func process(elements []string, index int) (res interface{}, toContinueIndex int) {
	str := elements[index]
	if isCompleteClosure(str) {
		return processComplete(elements, index), index + 1
	}

	if isClass(elements, index) {
		return processClass(elements, index)
	} else if isArray(elements, index) {
		return processArray(elements, index)
	}

	return nil, index + 1
}

func processClass(elements []string, index int) (res interface{}, toContinueIndex int) {
	m := make(map[string]interface{})

	for i := index; i < len(elements); {
		toContinueIndex = -1
		str := elements[i]
		strippedString := str
		if (endsWithRoundBracket(str) || beginsWithRoundBracket(str)) && isClass(elements, i) {
			strippedString = stripTrailingBrackets(str)
		}
		parts := strings.Split(strippedString, ",")
		for _, s := range parts {
			pair := strings.SplitN(s, "=", 2)
			if len(pair) == 1 {
				if pair[0] != "" {
					log("not expected!!! class field can't be processed", pair)
				}
				continue
			}
			if strings.TrimSpace(pair[1]) != "" {
				m[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
			} else {
				m[strings.TrimSpace(pair[0])], toContinueIndex = process(elements, i+1)
			}

		}
		if endsWithRoundBracket(str) && isClass(elements, i) {
			return m, i + 1
		}

		if toContinueIndex != -1 {
			i = toContinueIndex
		} else {
			i++
		}
	}
	panic("no matching close round bracket found")
}

func processArray(elements []string, index int) (res interface{}, toContinueIndex int) {
	a := make([]interface{}, 0)

	for i := index; i < len(elements); {
		if i == index && elements[i] == "[" {
			i++
			continue
		}
		str := elements[i]
		if endsWithSquareBracket(str) && isArray(elements, i) {
			return a, i + 1
		}
		r, ii := process(elements, i)
		i = ii // for some reason, golang doesn't allow me to reassign a value to i
		a = append(a, r)
	}
	panic("no matching close square bracket found")
}

func processComplete(elements []string, index int) interface{} {
	str := stripTrailingBrackets(elements[index])
	parts := strings.Split(str, ",")
	if isClass(elements, index) {
		m := make(map[string]interface{})
		for _, s := range parts {
			pair := strings.SplitN(s, "=", 2)
			m[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
		}
		return m
	} else if isArray(elements, index) {
		a := make([]string, 0)
		for _, s := range parts {
			a = append(a, strings.TrimSpace(s))
		}
		return a
	} else {
		log("processed complete none", elements[index])
		return elements[index]
	}
}

func stripTrailingBrackets(s string) string {
	a := strings.Split(s, "")
	prefixIndex := 0
	suffixIndex := len(a)
	for i := 0; i < len(a); i++ {
		if strings.Contains("[(", a[i]) {
			prefixIndex = i + 1
			break
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if strings.Contains(")]", a[i]) {
			suffixIndex = i
			break
		}
	}

	return strings.TrimSpace(s[prefixIndex:suffixIndex])
}

func isClass(elements []string, index int) bool {
	return isWhatObject(elements, index) == "class"
}

func isArray(elements []string, index int) bool {
	return isWhatObject(elements, index) == "array"
}

// isWhatObject denotes whether the string at index can be classified as a class object without other context
func isWhatObject(elements []string, index int) string {
	str := elements[index]
	if beginsWithRoundBracket(str) || endsWithRoundBracket(str) {
		if len(strings.Split(str, "=")) <= 1 &&
			!(str == ")" || str == "),") {
			return "none"
		}
		return "class"
	} else if beginsWithSquareBracket(str) || endsWithSquareBracket(str) {
		return "array"
	} else {
		//return isInWhatObject(elements, index)
		return "none"
	}
}

func pullClassNamesFromPreviousElement(elements []string) (res []string) {
	tmp := copyStringSlice(elements)
	for i, element := range tmp {
		if element[0] == '(' {
			var className string
			idx := strings.LastIndexAny(tmp[i-1], SPECIAL_CHARS)
			if idx == -1 {
				className = tmp[i-1]
				tmp[i-1] = ""
			} else {
				className = tmp[i-1][idx+1:]
				tmp[i-1] = tmp[i-1][:idx+1]
			}
			tmp[i] = className + tmp[i]
		}
	}

	for _, s := range tmp {
		if s != "" {
			res = append(res, s)
		}
	}

	return
}

//
//func isInClass(elements []string, index int) bool {
//	return isInWhatObject(elements, index) == "class"
//}
//
//func isInArray(elements []string, index int) bool {
//	return isInWhatObject(elements, index) == "array"
//}

//func isInWhatObject(elements []string, index int) string {
//	if index < 0 || index >= len(elements) {
//		return "none"
//	} else if !containsCompleteClosure(elements, index) &&
//		(endsWithRoundBracket(elements[index]) || endsWithSquareBracket(elements[index])) {
//		return isInWhatObject_ForwardsRecursion(elements, index)
//	} else {
//		return isInWhatObject_BackwardsRecursion(elements, index)
//	}
//}
//
//func isInWhatObject_ForwardsRecursion(elements []string, index int) string {
//	if index+1 < 0 || index+1 >= len(elements) {
//		return "none"
//	}
//	str := elements[index+1]
//	if beginsWithRoundBracket(str) && !endsWithRoundBracket(str) {
//		return "class"
//	} else if beginsWithSquareBracket(str) && !endsWithSquareBracket(str) {
//		return "array"
//	} else {
//		return isInWhatObject_ForwardsRecursion(elements, index+1)
//	}
//}
//
//func isInWhatObject_BackwardsRecursion(elements []string, index int) string {
//	if index-1 < 0 || index-1 >= len(elements) {
//		return "none"
//	}
//	str := elements[index-1]
//	if beginsWithRoundBracket(str) && !endsWithRoundBracket(str) {
//		return "class"
//	} else if beginsWithSquareBracket(str) && !endsWithSquareBracket(str) {
//		return "array"
//	} else {
//		return isInWhatObject_BackwardsRecursion(elements, index-1)
//	}
//}
//
func isCompleteClosure(str string) bool {
	return (beginsWithRoundBracket(str) && endsWithRoundBracket(str)) ||
		(beginsWithSquareBracket(str) && endsWithSquareBracket(str))
}

func pushLeadingCommasToPreviousElement(elements []string) (res []string) {
	res = copyStringSlice(elements)
	for i, element := range res {
		element = strings.TrimSpace(element)
		if element[0] == ',' {
			res[i-1] = res[i-1] + ","
			element = strings.TrimSpace(element[1:])
		}
		res[i] = element
	}

	return
}

func copyStringSlice(source []string) []string {
	return append(make([]string, 0, len(source)), source...)
}

func log(str string, obj interface{}) {
	fmt.Printf("%s \t {%#v}\n", str, obj)
}
