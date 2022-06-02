package lombokString

type LombokString struct {
	LString string
}

func New(lombokString string) *LombokString {
	return &LombokString{LString: lombokString}
}

func (ls LombokString) recursiveParse(index int, history InfoFromPrevRecursionLayer) (result *LombokObject, toContinueIndex int) {
	memory := newMemory()
	result = newLombokObject(history)

	for i := index; i < len(ls.LString); i++ {
		currentChar := string(ls.LString[i])

		if isOpenBrackets(currentChar) {
			if contIndex, o := ls.handleOpenBracket(i, result, memory); o != nil {
				return o, contIndex
			} else {
				i = contIndex
			}
		} else if isCloseBrackets(currentChar) {
			if history.openBracket == "" || !isMatchingBrackets(history.openBracket, currentChar) {
				panic("The inputted string has an invalid structure. Please check that you have copied the entire string.")
			}

			ls.handleCloseBracket(result, memory)
			return result, i
		}

		if ls.isEndOfElement(memory, currentChar) {
			ls.commitElementIntoResult(result, memory)
		} else if ls.isEndOfFieldName(memory, currentChar) {
			memory.fieldName = memory.continuedStringTrimmed()
		}

		if ls.isContinuationOfElement(memory, currentChar) {
			memory.continuedString = memory.continuedString + currentChar
		} else if ls.isStartOfElement(memory, currentChar) {
			ls.handleStartOfElement(memory, currentChar)
		}

		if isSpecialCharacter(currentChar) {
			ls.handleEndOfText(memory)
		}
	}
	return result, len(ls.LString)
}

func isMatchingBrackets(openBracket string, currentChar string) bool {
	return (openBracket == openRound && currentChar == closeRound) || (openBracket == openSquare && currentChar == closeSquare)
}

func (ls LombokString) handleOpenBracket(index int, result *LombokObject, memory *Memory) (toContinueIndex int, lombokObject *LombokObject) {
	currentChar := string(ls.LString[index])
	className := ""
	if openRound == currentChar {
		className = memory.continuedStringTrimmed()
	}
	nextLayerHistory := InfoFromPrevRecursionLayer{
		openBracket: currentChar,
		className:   className,
	}

	lombokObject, toContinueIndex = ls.recursiveParse(index+1, nextLayerHistory)

	switch result.objType {
	case "none":
		return toContinueIndex, lombokObject
	case "object":
		(*result.tempMap)[memory.fieldName] = lombokObject
		break
	case "array":
		*result.tempArr = append(*result.tempArr, lombokObject)
		break
	}

	memory.fieldName = ""
	return toContinueIndex, nil
}

func (ls LombokString) handleCloseBracket(result *LombokObject, memory *Memory) {
	ls.commitElementIntoResult(result, memory)
}

func (ls LombokString) commitElementIntoResult(result *LombokObject, memory *Memory) {
	if result.objType == "array" && memory.continuedString != "" {
		*result.tempArr = append(*result.tempArr, memory.continuedStringTrimmed())
	} else if result.objType != "array" && memory.fieldName != "" {
		(*result.tempMap)[memory.fieldName] = memory.continuedStringTrimmed()
		memory.fieldName = ""
	}
}

// TODO: handle scenario whereby we might want to register = and , as non-special characters
func isSpecialCharacter(char string) bool {
	specialChars := `[](){}=,`
	for _, c := range specialChars {
		if char == string(c) {
			return true
		}
	}
	return false
}

func (ls LombokString) isEndOfElement(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() &&
		(isCloseBrackets(currentChar) || currentChar == ",")
}

func (ls LombokString) isEndOfFieldName(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() && "=" == currentChar
}

func (ls LombokString) isContinuationOfElement(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() && !isSpecialCharacter(currentChar)
}

func (ls LombokString) isStartOfElement(memory *Memory, currentChar string) bool {
	return !memory.isPrevCharAString() && !isSpecialCharacter(currentChar)
}

func (ls LombokString) handleStartOfElement(memory *Memory, currentChar string) {
	memory.continuedString = currentChar
}

func (ls LombokString) handleEndOfText(memory *Memory) {
	memory.continuedString = ""
}
