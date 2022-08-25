package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.
Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием
этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки:
распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	time := GetTime()

	fmt.Println(time.Format("2006-01-02 15:04:05"))
}

/*
GetTime возвращает текущее время, полученное с сервера 0.beevik-ntp.pool.ntp.org.
При ошибке выводит текст ошибки в stderr и завершает работу программы с кодом 1.
*/
func GetTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1)
	}

	return time
}
