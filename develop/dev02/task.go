package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую
повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку.
Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("Hello world")
}

func unpack(line string) (string, error) {
	unpacked := []string{}

	escape := false

	s := strings.Split(line, "")

	for i, r := range s {
		if r == "\\" {
			if escape {
				unpacked = append(unpacked, r)
				escape = false
				continue
			}
			escape = true
			continue
		}

		isDigit, err := regexp.MatchString("[1-9]", r)
		if err != nil {
			return "", err
		}

		if i == 0 && isDigit {
			return "", fmt.Errorf("error: invalid line")
		}

		if isDigit {
			if !escape {
				n, err := strconv.Atoi(r)
				if err != nil {
					return "", err
				}

				unpacked = append(unpacked, strings.Repeat(s[i-1], n-1))
				continue
			}

			unpacked = append(unpacked, r)
			escape = false
		}

		isLetter, err := regexp.MatchString("[a-zA-Z]", r)
		if err != nil {
			return "", err
		}

		if isLetter {
			unpacked = append(unpacked, r)
			continue
		}
	}

	return strings.Join(unpacked, ""), nil
}
