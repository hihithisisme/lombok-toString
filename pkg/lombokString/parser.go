package lombokString

type LombokString struct {
	LString string
}

func New(lombokString string) *LombokString {
	return &LombokString{LString: lombokString}
}

func (ls LombokString) recursiveParse(index int, history InfoFromPrevRecursionLayer) (result *LombokObject, toContinueIndex int) {
	result = newLombokObject(history)
	memory := newMemory(result.objType)

	for i := index; i < len(ls.LString); i++ {
		var tmpMem Memory
		tmpMem = *memory
		tempMemory := &tmpMem

		currentChar := string(ls.LString[i])

		if isOpenBrackets(currentChar) {
			if contIndex, o := ls.handleOpenBracket(i, result, memory, tempMemory); o != nil {
				return o, contIndex
			} else {
				i = contIndex
			}
		} else if isCloseBrackets(currentChar) {
			if history.openBracket == "" || !isMatchingBrackets(history.openBracket, currentChar) {
				panic("The inputted string has an invalid structure. Please check that you have copied the entire string.")
			}

			ls.handleCloseBracket(result, memory, tempMemory)
			return result, i
		}

		if isEndOfElement(memory, currentChar) {
			ls.commitElementIntoResult(result, memory, tempMemory)
		} else if isEndOfFieldName(memory, currentChar) {
			tempMemory.fieldName = memory.continuedStringTrimmed()
		}

		if isContinuationOfElement(memory, currentChar) {
			tempMemory.continuedString = memory.continuedString + currentChar
		} else if isStartOfElement(memory, currentChar) {
			handleStartOfElement(tempMemory, currentChar)
		}

		if isSpecialCharacter(currentChar, memory) {
			handleEndOfText(tempMemory)
		}

		memory = tempMemory
	}
	return result, len(ls.LString)
}

func isMatchingBrackets(openBracket string, currentChar string) bool {
	return (openBracket == openRound && currentChar == closeRound) || (openBracket == openSquare && currentChar == closeSquare)
}

func (ls LombokString) handleOpenBracket(index int, result *LombokObject, memory *Memory, tempMemory *Memory) (toContinueIndex int, lombokObject *LombokObject) {
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

	tempMemory.fieldName = ""
	return toContinueIndex, nil
}

func (ls LombokString) handleCloseBracket(result *LombokObject, memory *Memory, tempMemory *Memory) {
	ls.commitElementIntoResult(result, memory, tempMemory)
}

func (ls LombokString) commitElementIntoResult(result *LombokObject, memory *Memory, tempMemory *Memory) {
	if result.objType == "array" && memory.continuedString != "" {
		*result.tempArr = append(*result.tempArr, memory.continuedStringTrimmed())
	} else if result.objType != "array" && memory.fieldName != "" {
		(*result.tempMap)[memory.fieldName] = memory.continuedStringTrimmed()
		tempMemory.fieldName = ""
	}
}

// TODO: handle scenario whereby we might want to register = and , as non-special characters
func isSpecialCharacter(char string, memory *Memory) bool {
	specialCharacters := `[](),`
	for _, c := range specialCharacters {
		if char == string(c) {
			return true
		}
	}

	if char == "=" {
		return shouldEqualSignBeSpecial(memory)
	}
	return false
}

func shouldEqualSignBeSpecial(memory *Memory) bool {
	if memory.objType == "array" {
		return false
	} else if memory.objType == "object" && memory.fieldName != "" {
		return false
	} else {
		return true
	}
}

func isEndOfElement(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() &&
		(isCloseBrackets(currentChar) || currentChar == ",")
}

func isEndOfFieldName(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() && "=" == currentChar && shouldEqualSignBeSpecial(memory)
}

func isContinuationOfElement(memory *Memory, currentChar string) bool {
	return memory.isPrevCharAString() && !isSpecialCharacter(currentChar, memory)
}

func isStartOfElement(memory *Memory, currentChar string) bool {
	return !memory.isPrevCharAString() && !isSpecialCharacter(currentChar, memory)
}

func handleStartOfElement(memory *Memory, currentChar string) {
	memory.continuedString = currentChar
}

func handleEndOfText(memory *Memory) {
	memory.continuedString = ""
}
