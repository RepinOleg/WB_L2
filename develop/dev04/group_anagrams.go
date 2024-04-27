package main

import (
	"fmt"
	"sort"
	"strings"
)

// GroupAnagrams принимает на вход слайс строк, возвращает map со значением в виде слайса анаграмм
func GroupAnagrams(input []string) map[string][]string {

	input = uniqueLowerStrings(input)
	anagramMap := make(map[string][]string)
	for _, str := range input {
		s := sortString(str)
		anagramMap[s] = append(anagramMap[s], str)
	}

	res := make(map[string][]string)
	for _, val := range anagramMap {
		if len(val) > 1 {
			firstWord := val[0]
			sort.Strings(val)
			res[firstWord] = val
		}
	}
	return res
}

// uniqueLowerStrings возвращает слайс не повторяющихся строк в нижнем регистре
func uniqueLowerStrings(lines []string) []string {
	mp := make(map[string]bool)
	result := make([]string, 0, len(lines))
	for _, str := range lines {
		if !mp[str] {
			mp[str] = true
			result = append(result, strings.ToLower(str))
		}
	}
	return result
}

// sortStrings сортирует строку по возрастанию в кодировке utf8
func sortString(s string) string {
	str := []rune(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	return string(str)
}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(GroupAnagrams(input))
}
