package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя)
и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а
данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу
(через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	initLog()

	arguments, err := parseArgs()
	if err != nil {
		log.Fatalln(err.Error())
	}

	timeout := parseFlags()

	Execute(arguments, timeout)
}

// Execute запускает основные функции программы.
func Execute(arguments Arguments, timeout int) {
	conn, err := connect(arguments.Host, arguments.Port, timeout)
	if err != nil {
		log.Fatalln(err.Error())
	}

	go copyTo(os.Stdout, conn)
	copyTo(conn, os.Stdin)
}

// connect выполняет подключение к указанному адресу и порту.
// Если подключение не происходит в течение timeout секунд,
// то возвращает ошибку.
func connect(host, port string, timeout int) (net.Conn, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", address,
		time.Second*time.Duration(timeout))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// copyTo выполняет обмен данными между указанными стримами.
func copyTo(writer io.Writer, reader io.Reader) {
	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatalln(err.Error())
	}
}

// parseFlags парсит флаги запуска программы.
func parseFlags() int {
	t := flag.Int("timeout", 10, "")

	flag.Parse()

	return *t
}

// Arguments хранит аргументы запуска программы.
type Arguments struct {
	Host string
	Port string
}

// parseArgs парсит аргументы запуска программы.
func parseArgs() (Arguments, error) {
	args := Arguments{}

	switch {
	case len(os.Args) < 3:
		return args, fmt.Errorf("not enought arguments")
	case len(os.Args) == 3:
		args.Host = os.Args[1]
		args.Port = os.Args[2]
		return args, nil
	case len(os.Args) == 4:
		args.Host = os.Args[2]
		args.Port = os.Args[3]
		return args, nil
	default:
		return args, fmt.Errorf("redundant arguments passed")
	}
}

// initLog инициализирует логгер.
func initLog() {
	log.SetFlags(log.Ldate | log.Llongfile)
}
