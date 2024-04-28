package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	column      int
	numeric     bool
	reverse     bool
	unique      bool
	sorted      bool
	ignoreSpace bool
}

var f flags

func init() {
	flag.IntVar(&f.column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&f.numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&f.reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&f.unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&f.sorted, "c", false, "отсортированы ли данные")
	flag.BoolVar(&f.ignoreSpace, "b", false, "игорировать хвостовые пробелы")
}

func main() {
	flag.Parse()
	lines, err := readFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}

	if f.sorted {
		res, ok := isSorted(lines)
		if !ok {
			fmt.Println("disorder: " + res)
		}
		return
	}

	sort.Slice(lines, compareLines(lines))

	if f.unique {
		lines = uniqueLines(lines)
	}
	printLines(lines)
}

func uniqueLines(lines []string) []string {
	uniqueLine := []string{lines[0]}

	mp := make(map[string]bool, len(lines))
	mp[lines[0]] = true

	for i := 1; i < len(lines); i++ {
		if !mp[lines[i]] {
			uniqueLine = append(uniqueLine, lines[i])
			mp[lines[i]] = true
		}
	}
	return uniqueLine
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func isSorted(lines []string) (string, bool) {
	for i := 1; i < len(lines); i++ {
		if compareLines(lines)(i, i-1) {
			return lines[i], false
		}
	}
	return "", true
}

func compareLines(lines []string) func(i, j int) bool {
	return func(i, j int) bool {
		a, b := lines[i], lines[j]

		if f.column > 0 {
			fieldsA := strings.Fields(a)
			fieldsB := strings.Fields(b)

			if f.column-1 < len(fieldsA) && f.column-1 < len(fieldsB) {
				a = fieldsA[f.column-1]
				b = fieldsB[f.column-1]
			}
		}

		if f.ignoreSpace {
			a = strings.TrimRight(a, " ")
			b = strings.TrimRight(b, " ")
		}

		if f.numeric {
			aNum, _ := strconv.ParseFloat(a, 64)
			bNum, _ := strconv.ParseFloat(b, 64)
			return aNum < bNum != f.reverse
		}

		return a < b != f.reverse
	}
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
