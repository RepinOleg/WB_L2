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
Выведется error
Из функции test возвращается интерфейс с типом *customError и значением nil и записывается в err
Происходит сравнение err[nil, customError] и [nil, nil]
Структура customError реализует интерфейс error т.к. имплементирует метод Error()

```