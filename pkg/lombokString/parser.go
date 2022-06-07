package lombokString

import "strings"

func SegmentByBrackets(in string) []string {
	fields := []string{}
	pointer := 0
	for i, char := range in {
		if char == '[' || char == '(' {
			fields = append(fields, in[pointer:i])
			pointer = i
		} else if char == ']' || char == ')' {
			fields = append(fields, in[pointer:i+1])
			pointer = i + 1
		}
	}

	if fields[0] == "" {
		fields = fields[1:]
	}

	for i, field := range fields {
		field = strings.TrimSpace(field)
		if field[0] == ',' {
			fields[i-1] = fields[i-1] + ","
			field = strings.TrimSpace(field[1:])
		}
		fields[i] = field
	}

	return fields
}
