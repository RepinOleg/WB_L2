Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Выведется
2
1

В функции test() создается именованный возвращаемый параметр x, особенность defer в том что он может изменять такие параметры
В функции anotherTest() возвращается обычная переменная, на которую defer никак не вляет

defer работает по принципу стэка, он выполняется после вызова return или panic, но до реального возврата из функции

```