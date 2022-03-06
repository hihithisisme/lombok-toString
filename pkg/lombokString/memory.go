package lombokString

import "strings"

type Memory struct {
	isPrevCharAString bool
	continuedString   string
	fieldName         string
}

func newMemory() *Memory {
	return &Memory{
		isPrevCharAString: false,
		continuedString:   "",
		fieldName:         "",
	}
}

func (m Memory) continuedStringTrimmed() string {
	return strings.TrimSpace(m.continuedString)
}

type InfoFromPrevRecursionLayer struct {
	openBracket string
	className   string
}
