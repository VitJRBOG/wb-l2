package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCdCmd(t *testing.T) {
	fullPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatalf("\nan error occured: %s", err.Error())
	}

	input1 := []string{
		"/private/var/folders/pv", "../", "/",
	}

	expectations := []string{
		"/private/var/folders/pv", "/private/var/folders", "/",
	}

	for i := 0; i < len(input1); i++ {
		input := []string{fullPath, "cd", input1[i]}
		_, err := cdCmd(&input)
		if err != nil {
			t.Fatalf("an error occured: %s", err.Error())
		}

		dir, err := os.Getwd()
		if err != nil {
			t.Fatalf("an error occured: %s", err.Error())
		}

		if dir != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\n got %s\n",
				input1[i], expectations[i], dir)
		}
	}

}

func TestPwdCmd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("an error occured: %s", err.Error())
	}

	input := []string{dir, "pwd"}

	result, err := pwdCmd(&input)
	if err != nil {
		t.Fatalf("an error occured: %s", err.Error())
	}

	if result != dir {
		t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
			input, dir, result)
	}
}

func TestEchoCmd(t *testing.T) {
	input1 := []string{
		"Hello world", "123", "aaaa",
	}

	expectations := []string{
		"Hello world", "123", "aaaa",
	}

	for i := 0; i < len(input1); i++ {
		input := []string{"", "echo", input1[i]}

		result, err := echoCmd(&input)
		if err != nil {
			t.Fatalf("an error occured: %s", err.Error())
		}
		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input, expectations[i], result)
		}
	}
}
