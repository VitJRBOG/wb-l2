package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

func ExecuteStrategyExample() {
	a := 5.12
	b := 3.14

	actions := []string{"addition", "substracting", "multiplication"}

	for _, action := range actions {
		switch action {
		case "addition":
			CalculateNumbers(&AdditionStrategy{}, a, b)
		case "substracting":
			CalculateNumbers(&SubstractingStrategy{}, a, b)
		case "multiplication":
			CalculateNumbers(&MultiplicationStrategy{}, a, b)
		default:
			panic("couldn't find proper strategy")
		}
	}

}

func CalculateNumbers(strategy Strategy, a, b float64) {
	strategy.Execute(a, b)
}

type Strategy interface {
	Execute(a, b float64)
}

type AdditionStrategy struct{}

func (s *AdditionStrategy) Execute(a, b float64) {
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

type SubstractingStrategy struct{}

func (s *SubstractingStrategy) Execute(a, b float64) {
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
}

type MultiplicationStrategy struct{}

func (s *MultiplicationStrategy) Execute(a, b float64) {
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
}
