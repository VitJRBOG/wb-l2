package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	flags := parseFlags()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		}

		result := Execute(text, flags)

		fmt.Println(result)
	}
}

func Execute(text string, flags Flags) string {
	return SelectColumn(text, flags.Delimiter, flags.OnlySeparated, flags.Column)
}

/*
SelectColumn делит полученную строку по delimiter на столбцы, извлекает столбец
с номером columnNumber и возвращает его.
*/
func SelectColumn(text, delimiter string, onlySeparated bool, columnNumber int) string {
	selectedColumn := []string{}

	rows := strings.Split(text, "\n")

	for _, row := range rows {
		fields := strings.Split(row, delimiter)
		if onlySeparated && len(fields) <= 1 {
			continue
		}
		if len(fields) > columnNumber {
			selectedColumn = append(selectedColumn, fields[columnNumber])
		}
	}

	return strings.Join(selectedColumn, "\n")
}

/*
Flags хранит значения флагов запуска программы.
*/
type Flags struct {
	Column        int
	Delimiter     string
	OnlySeparated bool
}

/*
parseFlags парсит флаги запуска программы.
*/
func parseFlags() Flags {
	f := flag.Int("f", -1, "")
	d := flag.String("d", "	", "")
	s := flag.Bool("s", false, "")

	flag.Parse()

	return Flags{
		Column:        *f,
		Delimiter:     *d,
		OnlySeparated: *s,
	}
}
