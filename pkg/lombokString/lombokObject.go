package lombokString

const openSquare = "["
const closeSquare = "]"

const openRound = "("
const closeRound = ")"

func isOpenBrackets(s string) bool {
	return s == openSquare || s == openRound
}

func isCloseBrackets(s string) bool {
	return s == closeSquare || s == closeRound
}

// TODO: handle scenario whereby LombokObject is only string -- should convert LombokObject to string
type LombokObject struct {
	objType   string
	tempMap   *map[string]interface{}
	tempArr   *[]interface{}
	className string
}

func newLombokObject(history InfoFromPrevRecursionLayer) *LombokObject {
	tempMap := make(map[string]interface{})
	tempArr := make([]interface{}, 0)
	return &LombokObject{
		objType:   getType(history.openBracket),
		tempMap:   &tempMap,
		tempArr:   &tempArr,
		className: history.className,
	}
}

func getType(openBrackets string) string {
	if openBrackets == "" {
		return "none"
	}
	switch openBrackets {
	case openRound:
		return "object"
	case openSquare:
		return "array"
	default:
		panic("invalid brackets")
	}
}
