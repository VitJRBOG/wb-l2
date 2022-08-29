package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого -
слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"тяпка", "пятак", "листок", "столик", "пятка", "слиток"}

	result := Execute(&words)

	fmt.Println(*result)

}

type Word struct {
	First bool
	Index int
}

func Execute(words *[]string) *map[string][]string {
	m := make(map[string]map[string]Word)

	for i, word := range *words {
		word := strings.ToLower(word)

		letters := sortLetters(word)

		first := false

		if _, ok := m[letters]; !ok {
			m[letters] = map[string]Word{}
			first = true
		}

		if _, ok := m[letters][word]; !ok {
			m[letters][word] = Word{
				First: first,
				Index: i,
			}
		}
	}

	result := map[string][]string{}

	for _, values := range m {
		k := []string{}
		for _, wordInfo := range values {
			k = append(k, strings.ToLower((*words)[wordInfo.Index]))
			if wordInfo.First && len(k) > 1 {
				k[0], k[len(k)-1] = k[len(k)-1], k[0]
				continue
			}
		}
		if len(k) > 1 {
			result[k[0]] = k
			sort.Sort(sort.StringSlice(k))
		}
	}

	return &result
}

type runeSlice []rune

func (s runeSlice) Len() int {
	return len(s)
}
func (s runeSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s runeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortLetters(word string) string {
	letters := []rune(word)

	sort.Sort(runeSlice(letters))

	return string(letters)
}
