package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Keys хранит параметры ключей, с которыми была запущена программа
type Keys struct {
	ColumnSort  int
	NumbersSort bool
	Reverse     bool
	SelectUnic  bool
}

func main() {
	text := "some\ntest\nstring\nwith\nfew\nlines"

	text = Execute(text)

	fmt.Println(text)
}

// Execute запускает первичные функции программы
func Execute(text string) string {
	keys := parseKeys()

	switch {
	case keys.ColumnSort > -1:
		text = sortByColumn(text, 0)
	case keys.NumbersSort:
		text = sortByNumbers(text)
	case keys.Reverse:
		text = reverse(text)
	case keys.SelectUnic:
		text = selectUnic(text)
	}

	return text
}

func parseKeys() Keys {
	k := flag.Int("k", -1, "")
	n := flag.Bool("n", false, "")
	r := flag.Bool("r", false, "")
	u := flag.Bool("u", false, "")

	flag.Parse()

	return Keys{
		ColumnSort:  *k,
		NumbersSort: *n,
		Reverse:     *r,
		SelectUnic:  *u,
	}
}

func sortByColumn(text string, columnNumber int) string {
	lines := strings.Split(text, "\n")

	if len(lines) == 0 {
		return text
	}

	array := [][]string{}

	for _, line := range lines {
		array = append(array, strings.Split(line, " "))
	}

	newLines := []string{}

	for i := 0; i < len(array); i++ {
		newLines = append(newLines, array[i][columnNumber])
	}

	sort.Sort(sort.StringSlice(newLines))

	for i := 0; i < len(array); i++ {
		array[i][columnNumber] = newLines[i]
	}

	for i, line := range array {
		lines[i] = strings.Join(line, " ")
	}

	return strings.Join(lines, "\n")
}

func sortByNumbers(text string) string {
	newLines := []string{}

	lines := strings.Split(text, "\n")

	numbers := []int{}

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			newLines = append(newLines, line)
			continue
		}

		numbers = append(numbers, n)
	}

	if len(numbers) > 0 {
		sort.Ints(numbers)
	}

	last := -1
	for i := 0; i < len(lines); i++ {
		if i < len(numbers) {
			lines[i] = strconv.Itoa(numbers[i])
			continue
		}

		if last == -1 {
			last = i
		}

		lines[i] = newLines[i-last]
	}

	return strings.Join(lines, "\n")
}

func reverse(text string) string {
	lines := strings.Split(text, "\n")

	sort.Sort(sort.Reverse(sort.StringSlice(lines)))

	return strings.Join(lines, "\n")
}

func selectUnic(text string) string {
	unicLines := []string{}

	lines := strings.Split(text, "\n")

	items := map[string]struct{}{}

	for _, line := range lines {
		if _, ok := items[line]; ok {
			continue
		}

		items[line] = struct{}{}
		unicLines = append(unicLines, line)
	}

	return strings.Join(unicLines, "\n")
}

// TODO: сделать дополнительное задание
