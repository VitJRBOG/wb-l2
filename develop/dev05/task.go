package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	flags := parseFlags()
	substring, path := parseArgs()

	text := readFile(path)

	result := Execute(flags, text, substring)

	fmt.Println(result)
}

/*
Execute запускает основные функции программы.
*/
func Execute(keys Flags, text, substring string) string {
	switch {
	case keys.SelectAfter > -1:
		return SelectLinesAfter(text, substring, keys.SelectAfter)
	case keys.SelectBefore > -1:
		return SelectLinesBefore(text, substring, keys.SelectBefore)
	case keys.SelectAround > -1:
		return SelectLinesAround(text, substring, keys.SelectAround)
	case keys.LineCount:
		linesCount := GetLinesCount(text, substring)
		return strconv.Itoa(linesCount)
	case keys.CaseIgnoring:
		return SelectLineIgnoringCase(text, substring)
	case keys.Exclude:
		return SelectLinesWithout(text, substring)
	case keys.Exact:
		return SelectExactLine(text, substring)
	case keys.LineNumber:
		lineNumber := GetLineNumber(text, substring)
		return strconv.Itoa(lineNumber)
	}

	return ""
}

/*
SelectLinesAfter ищет в тексте указанную подстроку
и возвращает lineCount строк перед ней.
*/
func SelectLinesAfter(text, substring string, lineCount int) string {
	lines := strings.Split(text, "\n")

	linesAfter := []string{}

	for i, line := range lines {
		if !strings.Contains(line, substring) {
			continue
		}

		for j, n := i+1, lineCount; j < len(lines) && n > 0; j, n = j+1, n-1 {
			linesAfter = append(linesAfter, lines[j])
		}
		break
	}

	return strings.Join(linesAfter, "\n")
}

/*
SelectLinesBefore ищет в тексте указанную подстроку
и возвращает lineCount строк после нее.
*/
func SelectLinesBefore(text, substring string, lineCount int) string {
	lines := strings.Split(text, "\n")

	linesBefore := []string{}

	for i, line := range lines {
		if !strings.Contains(line, substring) {
			continue
		}

		for j, n := i-1, lineCount; j >= 0 && n > 0; j, n = j-1, n-1 {
			linesBefore = append(linesBefore, lines[j])
		}
		break
	}

	if len(linesBefore) > 0 {
		reverseSlice(linesBefore)
	}

	return strings.Join(linesBefore, "\n")
}

/*
SelectLinesAround ищет в тексте указанную подстроку
и возвращает lineCount строк вокруг нее.
*/
func SelectLinesAround(text, substring string, lineCount int) string {
	lines := strings.Split(text, "\n")

	linesAround := []string{}

	for i, line := range lines {
		if !strings.Contains(line, substring) {
			continue
		}

		nAfter := lineCount / 2
		nBefore := lineCount - nAfter

		for j, n := i, nBefore; j >= 0 && n > 0; j, n = j-1, n-1 {
			linesAround = append(linesAround, lines[j])
		}
		reverseSlice(linesAround)

		for j, n := i+1, nAfter; j < len(lines) && n > 0; j, n = j+1, n-1 {
			linesAround = append(linesAround, lines[j])
		}
		break
	}

	return strings.Join(linesAround, "\n")
}

/*
GetLinesCount ищет в тексте строки с указанными подстроками
и возвращает их количество.
*/
func GetLinesCount(text, substring string) int {
	linesCount := 0
	for _, line := range strings.Split(text, "\n") {
		if strings.Contains(line, substring) {
			linesCount++
		}
	}

	return linesCount
}

/*
SelectLineIgnoringCase ищет в тексте строку с указанной подстрокой, игнорируя регистр,
и возвращает ее.
*/
func SelectLineIgnoringCase(text, substring string) string {
	for _, line := range strings.Split(text, "\n") {
		lowercasedLine := strings.ToLower(line)
		lowercasedSubstring := strings.ToLower(substring)
		if strings.Contains(lowercasedLine, lowercasedSubstring) {
			return line
		}
	}

	return ""
}

/*
SelectLinesWithout ищет в тексте строки без указанной подстроки и возвращает их.
*/
func SelectLinesWithout(text, substring string) string {
	lines := strings.Split(text, "\n")

	linesWithout := []string{}

	for _, line := range lines {
		if !strings.Contains(line, substring) {
			linesWithout = append(linesWithout, line)
		}
	}

	return strings.Join(linesWithout, "\n")
}

/*
SelectExactLine ищет в тексте строку, точно похожую на указанную, и возвращает ее.
*/
func SelectExactLine(text, exactLine string) string {
	for _, line := range strings.Split(text, "\n") {
		if line == exactLine {
			return line
		}
	}

	return ""
}

/*
GetLineNumber ищет в тексте подстроку и возвращает номер строки, где она была найдена.
Если ничего не найдено - возвращает -1.
*/
func GetLineNumber(text, substring string) int {
	for i, line := range strings.Split(text, "\n") {
		if strings.Contains(line, substring) {
			return i
		}
	}

	return -1
}

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
Flags хранит значения флагов запуска программы.
*/
type Flags struct {
	SelectAfter  int
	SelectBefore int
	SelectAround int
	LineCount    bool
	CaseIgnoring bool
	Exclude      bool
	Exact        bool
	LineNumber   bool
}

/*
parseKeys парсит ключи
*/
func parseFlags() Flags {
	A := flag.Int("A", -1, "")
	B := flag.Int("B", -1, "")
	C := flag.Int("C", -1, "")
	c := flag.Bool("c", false, "")
	i := flag.Bool("i", false, "")
	v := flag.Bool("v", false, "")
	F := flag.Bool("F", false, "")
	n := flag.Bool("n", false, "")

	flag.Parse()

	return Flags{
		SelectAfter:  *A,
		SelectBefore: *B,
		SelectAround: *C,
		LineCount:    *c,
		CaseIgnoring: *i,
		Exclude:      *v,
		Exact:        *F,
		LineNumber:   *n,
	}
}

/*
parseArgs парсит аргументы запуска программы.
*/
func parseArgs() (string, string) {
	switch {
	case len(os.Args) < 4:
		fmt.Fprintf(os.Stderr, "\nerror: not enought arguments\n")
		os.Exit(1)
	case len(os.Args) == 4:
		return os.Args[2], os.Args[3]
	}

	return "", ""
}

/*
readFile читает файл по указанному пути.
*/
func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	return string(data)
}
