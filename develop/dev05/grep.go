package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/
type grepFlags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNumber bool
}

func main() {
	var f grepFlags
	parseFlags(&f)
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: grep [OPTION]... PATTERNS [FILE]...")
		return
	}

	// Формируем регулярное выражение
	pattern := flag.Arg(0)
	var re *regexp.Regexp
	if f.fixed {
		re = regexp.MustCompile(`^` + regexp.QuoteMeta(pattern) + `$`)
	} else {
		re = regexp.MustCompile(pattern)
	}

	if f.ignoreCase {
		re = regexp.MustCompile("(?i)" + pattern)
	}

	if f.context > 0 {
		f.after = f.context
		f.before = f.context
	}

	// проходим по всем переданным файлам
	for _, filename := range args[1:] {
		lines, err := readFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// ищем регулярное выражение в файле
		search(lines, re, f)
	}
}

func parseFlags(f *grepFlags) {
	flag.IntVar(&f.after, "A", 0, "Print  NUM  lines  of trailing context after matching lines.")
	flag.IntVar(&f.before, "B", 0, "Print NUM lines of leading context before matching lines.")
	flag.IntVar(&f.context, "C", 0, "Print  NUM  lines of output context.")
	flag.BoolVar(&f.count, "c", false, "Suppress normal output; instead print a count of matching lines for each input file.")
	flag.BoolVar(&f.ignoreCase, "i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other.")
	flag.BoolVar(&f.invert, "v", false, "Invert the sense of matching, to select non-matching lines.")
	flag.BoolVar(&f.fixed, "F", false, "Interpret PATTERNS as fixed strings, not regular expressions.")
	flag.BoolVar(&f.lineNumber, "n", false, "Prefix each line of output with the 1-based line number within its input file.")
	flag.Parse()
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
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

func search(lines []string, re *regexp.Regexp, f grepFlags) {
	matchCount := 0

	mp := make(map[int]bool)
	for idx, line := range lines {
		match := re.MatchString(line)
		if (match && !f.invert) || (!match && f.invert) {
			matchCount++

			if !f.count {
				printBefore(f, idx, lines, &mp)

				if !mp[idx] {
					printLine(f, idx, line)
				}

				printAfter(f, idx, lines, &mp)
			}
			mp[idx] = true
		}

	}
	if f.count {
		fmt.Println(matchCount)
	}
}

func printBefore(f grepFlags, idx int, lines []string, mp *map[int]bool) {
	if f.before > 0 {
		start := idx - f.before
		if start < 0 {
			start = 0
		}
		for i := start; i < idx && i < len(lines); i++ {
			if !(*mp)[i] {
				printLine(f, i, lines[i])
				(*mp)[i] = true
			}
		}
	}
}

func printAfter(f grepFlags, idx int, lines []string, mp *map[int]bool) {
	if f.after > 0 {
		start := idx + 1
		for i := start; i < len(lines) && i <= idx+f.after; i++ {
			if !(*mp)[i] {
				printLine(f, i, lines[i])
				(*mp)[i] = true
			}
		}
	}
}

func printLine(f grepFlags, idx int, line string) {
	if f.lineNumber {
		fmt.Printf("%d:", idx+1)
	}
	fmt.Println(line)
}
