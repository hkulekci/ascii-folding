package asciiFolding

import (
    "strconv"
    "bytes"
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
        _, ok := FoldMap[RuneToAscii(ch)];
        if ok {
            buffer.WriteString(FoldMap[RuneToAscii(ch)])
        } else {
            buffer.WriteRune(ch)
        }
    }
    return buffer.String()
}