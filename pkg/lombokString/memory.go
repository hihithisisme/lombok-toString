package lombokString

import "strings"

type Memory struct {
	// continuedString refers to the accumulated uncommitted string
	continuedString string
	// fieldName holds the key for maps. Must be stored in the recursion-level Memory because value might be another LombokObject
	fieldName string
	// objType of LombokObject
	objType string
}

func newMemory(objType string) *Memory {
	return &Memory{
		continuedString: "",
		fieldName:       "",
		objType:         objType,
	}
}

func (m Memory) isPrevCharAString() bool {
	if m.continuedString == "" {
		return false
	}

	return !isSpecialCharacter(string(m.continuedString[len(m.continuedString)-1]), &m)
}

func (m Memory) continuedStringTrimmed() string {
	return strings.TrimSpace(m.continuedString)
}

type InfoFromPrevRecursionLayer struct {
	openBracket string
	className   string
}
