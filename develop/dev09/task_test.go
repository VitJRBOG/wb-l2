package main

import (
	"os"
	"testing"
)

func TestCreateDir(t *testing.T) {
	input := "DirName"

	err := createDir(input)
	if err != nil {
		t.Fatalf("\nan error occured: %s\n", err.Error())
	}

	if _, err := os.Stat(input); err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("\nbad result: dir %s not found\n", input)
		}
	}

	err = os.Remove(input)
	if err != nil {
		t.Fatalf("\nan error occured: %s\n", err.Error())
	}
}

func TestWriteFile(t *testing.T) {
	input1 := "test.html"

	input2 := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>WB_L2</title>
		</head>
		<body>Hello world</body>
	</html>
	`

	err := writeFile(input1, []byte(input2))
	if err != nil {
		t.Fatalf("\nan error occured: %s\n", err.Error())
	}

	os.Remove(input1)
}

func TestSelectGlobalURLs(t *testing.T) {
	input := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>WB_L2</title>
		</head>
		<body>
		<ol>
			<li>https://example1.com</li>
			<li>https://example2.com</li>
			<li>https://example3.com/helloworld</li>
		</ol>
		</body>
	</html>
	`

	expectations := []string{
		"https://example1.com",
		"https://example2.com",
		"https://example3.com/helloworld",
	}

	result := selectGlobalURLs(input)

	if len(result) != len(expectations) {
		t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
			input, expectations, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expectations[i] {
			t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
				input, expectations, result)
		}
	}
}

func TestSelectHostname(t *testing.T) {
	input := "https://example.com/foo/bar/url/blablabla"

	expectations := "https://example.com"

	result, err := selectHostname(input)
	if err != nil {
		t.Fatalf("\nan error occured: %s\n", err.Error())
	}

	if expectations != result {
		t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
			input, expectations, result)
	}
}

func TestSelectLocalURLs(t *testing.T) {
	input := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>WB_L2</title>
		</head>
		<body>
		<ol>
			<li>/js/script.js</li>
			<li>/img/image.png</li>
			<li>/css/style.css</li>
		</ol>
		</body>
	</html>
	`

	expectations := []string{
		"https://example.com/js/script.js",
		"https://example.com/img/image.png",
		"https://example.com/css/style.css",
	}

	result := selectLocalURLs("https://example.com", input)

	if len(result) != len(expectations) {
		t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
			input, expectations, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expectations[i] {
			t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
				input, expectations, result)
		}
	}
}

func TestComposeDirName(t *testing.T) {
	input := "https://example.com/foo/bar/url/blablabla"

	expectations := "url"

	result := composeDirName(input)

	if expectations != result {
		t.Fatalf("\nbas result for %s\nexpected %s\ngot %s\n",
			input, expectations, result)
	}
}
