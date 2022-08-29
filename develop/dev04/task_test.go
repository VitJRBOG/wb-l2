package main

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	input := [][]string{
		{"тяпка", "пятак", "лаванда", "листок", "столик", "пятак", "пятка", "слиток"},
		{"кино", "кони", "Пятак", "пятка", "слиток", "листок", "столик", "листок", "Порт", "рог", "тяпка"}}

	expectations := []map[string][]string{
		{"листок": {"листок", "слиток", "столик"}, "тяпка": {"пятак", "пятка", "тяпка"}},
		{"кино": {"кино", "кони"}, "пятак": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}},
	}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result := Execute(&input[i])
		if !(reflect.DeepEqual(*result, expectations[i])) {
			t.Fatalf("\nbad result for %v\nexpected %v\ngot %v\n",
				input[i], expectations[i], *result)
		}
	}
}
