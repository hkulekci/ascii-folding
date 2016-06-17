package asciiFolding

import (
	"bytes"
	"strconv"
	"strings"
)

func RuneToAscii(r rune) string {
	if r < 128 {
		return string(r)
	} else {
		return strings.ToUpper(strconv.FormatInt(int64(r), 16))
	}
}

func Fold(str string) string {
	var buffer bytes.Buffer
	for _, ch := range []rune(str) {
		buffer.WriteString(FoldMap(RuneToAscii(ch)))
	}
	return buffer.String()
}
