Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

err не будет равен nil, так как функция test возвращает не интерфейс error, а указатель на структуру customError. Следовательно, у err появляется указатель на тип, после чего он уже не будет равен nil.
