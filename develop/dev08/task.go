package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	Run()
}

func Run() {
	reader := bufio.NewReader(os.Stdin)

	fullPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for {
		fmt.Fprintf(os.Stdin, "\ngosh:%s $ ", fullPath)

		args := []string{fullPath}

		args, err = receiveConsoleInput(reader, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		output, err := executeCommand(&args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		fullPath = args[0]

		fmt.Fprint(os.Stdin, output)
	}
}

func receiveConsoleInput(reader *bufio.Reader, args []string) ([]string, error) {
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	userInput = strings.TrimSuffix(userInput, "\n")

	args = append(args, strings.Split(userInput, " ")...)

	return args, nil
}

func executeCommand(args *[]string) (string, error) {
	switch (*args)[1] {
	case "cd":
		return cdCmd(args)
	case "pwd":
		return pwdCmd(args)
	case "echo":
		return echoCmd(args)
	case "kill":
		return killCmd(args)
	case "ps":
		return psCmd(args)
	case "fork":
		return forkCmd(args)
	case "exec":
		return execCmd(args)
	case "exit":
		return exitCmd()
	case "":
		return "", nil
	default:
		return "", fmt.Errorf("gosh: command not found: %s", (*args)[1])
	}
}

func cdCmd(args *[]string) (string, error) {
	err := os.Chdir((*args)[2])
	if err != nil {
		return "", err
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	(*args)[0] = dir

	return "", nil
}

func pwdCmd(args *[]string) (string, error) {
	return (*args)[0], nil
}

func echoCmd(args *[]string) (string, error) {
	return strings.Join((*args)[2:], " "), nil
}

func killCmd(args *[]string) (string, error) {
	pid, err := strconv.Atoi((*args)[2])
	if err != nil {
		return "", err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return "", err
	}

	err = process.Kill()
	if err != nil {
		return "", err
	}

	return "", nil
}

func psCmd(args *[]string) (string, error) {
	output, err := exec.Command((*args)[1]).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func forkCmd(args *[]string) (string, error) {
	cmd := exec.Command((*args)[2], (*args)[3:]...)
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(cmd.Process.Pid), nil
}

func execCmd(args *[]string) (string, error) {
	cmd := exec.Command((*args)[2], (*args)[3:]...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func exitCmd() (string, error) {
	os.Exit(0)

	return "", nil
}

// TODO: сделать дополнительное задание
