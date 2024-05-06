package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

func main() {
	fields := flag.String("f", "", "fields to select")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only lines with delimiter")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разделение строки на колонки
		columns := strings.Split(line, *delimiter)

		// Выбор запрошенных колонок
		selectedColumns := selectColumns(columns, *fields)

		// Вывод результата
		fmt.Println(strings.Join(selectedColumns, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// selectColumns выбирает запрошенные колонки из среза колонок.
func selectColumns(columns []string, fields string) []string {
	if fields == "" {
		return columns
	}

	var selectedColumns []string
	for _, field := range strings.Split(fields, ",") {
		if strings.Contains(field, "-") {
			rangeLimits := strings.Split(field, "-")
			start, end := parseRangeLimit(rangeLimits[0]), parseRangeLimit(rangeLimits[1])
			for i := start; i <= end; i++ {
				selectedColumns = append(selectedColumns, columns[i-1])
			}
		} else {
			index := parseRangeLimit(field)
			selectedColumns = append(selectedColumns, columns[index-1])
		}
	}

	return selectedColumns
}

// parseRangeLimit разбирает ограничение диапазона и обрабатывает ошибки.
func parseRangeLimit(limit string) int {
	index, err := strconv.Atoi(limit)
	if err != nil || index < 1 {
		fmt.Fprintln(os.Stderr, "invalid field:", limit)
		os.Exit(1)
	}
	return index
}
