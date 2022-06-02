package lombokString

import "strings"

type Memory struct {
	isPrevCharAString bool
	// continuedString refers to the accumulated uncommitted string
	continuedString string
	// TODO: can we remove fieldName and store it together with continuedString?
	// determine at commit-time whether it is meant to be a map element -- if yes, then parse it as such
	// fieldName holds the key name for maps since they need to be committed alongside the value
	fieldName string
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
