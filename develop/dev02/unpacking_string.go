package dev02

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

*/

// UnpackingString осуществляет примитивную распаковку строки, содержащую повторяющиеся символы
func UnpackingString(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	in := []rune(input)
	var escape bool
	var builder strings.Builder

	for i := 0; i < len(in); i++ {
		if unicode.IsDigit(in[i]) {
			count, _ := strconv.Atoi(string(in[i]))
			if i > 0 && unicode.IsLetter(in[i-1]) {
				writeToString(count-1, &builder, in[i-1])
			} else {
				return "", fmt.Errorf("inorrect input string %s", input)
			}
		} else if unicode.IsLetter(in[i]) {
			builder.WriteRune(in[i])
		} else if string(in[i]) == `\` {
			escape = true
			i++
			builder.WriteRune(in[i])
			if i+1 < len(in) && unicode.IsDigit(in[i+1]) {
				count, _ := strconv.Atoi(string(in[i+1]))
				writeToString(count-1, &builder, in[i])
				i++
			}
		} else {
			return "", fmt.Errorf("inorrect input string %s", input)
		}
	}

	if escape {
		builder.WriteString(" (*)")
	}

	return builder.String(), nil
}

func writeToString(count int, builder *strings.Builder, ch rune) {
	for j := 0; j < count; j++ {
		builder.WriteRune(ch)
	}
}
