Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Программа выведет
<nil>
false

Интерфейс хранит в себе тип интерфейса и тип самого значения, значение интерфейса является nil когда и значение и тип являются nil
Функция Foo возвращает [nil, *os.PathError] и сравнивается с [nil, nil] поэтому выводится false
Пустой интерфейс реализовывает любой тип, потому что в нем нет ни одного метода
```