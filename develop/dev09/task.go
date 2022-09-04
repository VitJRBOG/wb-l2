package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

Флаги:
-l, --level=ЧИСЛО глубина рекурсии
-o, --output-file=ФАЙЛ записывать сообщения (логи) в ФАЙЛ
*/

func main() {
	flags := parseFlags()

	f := initLogging(flags.LogFileName)
	defer f.Close()

	u, err := parseArgs()
	if err != nil {
		log.Fatalln(err)
	}

	err = Execute(u, flags)
	if err != nil {
		log.Fatalln(err)
	}
}

// Execute запускает основные функции программы.
func Execute(u string, flags Flags) error {
	err := downloading(u, composeDirName(u), flags.RecursionLevel)
	if err != nil {
		return err
	}
	return nil
}

// downloading вызывает функцию загрузки файлов.
// Если параметр recursionLevel не равен 0, то функция вызывается рекурсивно.
// Если параметр recursionLevel равен -1 - рекурсия выполняется до тех пор,
// пока внутри документов не закончался все ссылки.
// Если параметр recursionLevel больше 0 - то с каждой рекурсией параметр декрементируется.
func downloading(u, dirName string, recursionLevel int) error {
	filename, err := fetchFile(u, dirName)
	if err != nil {
		return err
	}

	switch {
	case recursionLevel == 0:
		return nil
	case recursionLevel > 0:
		recursionLevel--
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	globalURLs := selectGlobalURLs(string(data))

	hostname, err := selectHostname(u)
	if err != nil {
		return err
	}

	localURLs := selectLocalURLs(hostname, string(data))

	for _, u := range localURLs {
		err = downloading(u, path.Join(dirName, composeDirName(u)), recursionLevel)
		if err != nil {
			return err
		}
	}

	for _, u := range globalURLs {
		err = downloading(u, path.Join(dirName, composeDirName(u)), recursionLevel)
		if err != nil {
			return err
		}
	}

	return nil
}

// fetchFile скачивает файл, находящийся по ссылке u,
// и сохраняет его локально в папке dirName. Имя файла - последний элемент u.
func fetchFile(u, dirName string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = createDir(dirName)
	if err != nil {
		return "", err
	}

	filename := path.Join(dirName, path.Base(u))

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = writeFile(filename, data)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// createDir создает директорию с именем dirName.
// Если директория существует, то функция не делает ничего.
func createDir(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0777)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

// writeFile выполняет запись data в файл filename.
func writeFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// selectGlobalURLs находит все внешние URL в html-коде и возвращает строковый слайс с ними.
func selectGlobalURLs(pageCode string) []string {
	re := `(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
	urls := regexp.MustCompile(re).FindAllString(pageCode, -1)

	return urls
}

// selectHostname извлекает имя хоста из переданного u.
func selectHostname(u string) (string, error) {
	url, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	hostname := fmt.Sprintf("%s://%s", url.Scheme, url.Hostname())

	return hostname, nil
}

// selectLocalURLs находит все локальные URL в html-коде и возвращает строковый слайс с ними.
func selectLocalURLs(hostname, pageCode string) []string {
	re := `\/[\w\-_]+\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
	urls := regexp.MustCompile(re).FindAllString(pageCode, -1)

	for i := range urls {
		urls[i] = hostname + urls[i]
	}

	return urls
}

// composeDirName создает имя для директории из переданного u.
func composeDirName(u string) string {
	items := strings.Split(u, "/")

	dir := items[len(items)-2]
	if dir == "" {
		dir = items[len(items)-1]
	}

	return dir
}

// Flags хранит параметры (флаги) запуска программы.
type Flags struct {
	RecursionLevel int
	LogFileName    string
}

// parseFlags парсит флаги запуска программы.
func parseFlags() Flags {
	l := flag.Int("l", 0, "")
	o := flag.String("o", "", "")

	flag.Parse()

	return Flags{
		RecursionLevel: *l,
		LogFileName:    *o,
	}
}

// parseArgs парсит агрументы запуска программы и возвращает переданный url.
func parseArgs() (string, error) {
	if !strings.Contains(os.Args[len(os.Args)-1], "http") {
		return "", fmt.Errorf("URL has no gived")
	}

	return os.Args[len(os.Args)-1], nil
}

// initLogging инициализирует логгер.
func initLogging(filename string) *os.File {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	if filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err)
		}

		log.SetOutput(f)

		return f
	}

	return nil
}
