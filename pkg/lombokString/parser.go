package lombokString

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
