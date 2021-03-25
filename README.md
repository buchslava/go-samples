# GC
https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit#

# Interface

Interfaces in Go do not enforce a type to implement methods but interfaces are very powerful tools. A type can choose to implement methods of an interface. Using interfaces, a value can be represented in multiple types, AKA, polymorphism.

An interface is a collection of method signatures that a Type can implement (using methods). Hence interface defines (not declares) the behavior of the object (of the type Type).

# channels
https://habr.com/ru/post/308070/

qcount — количество элементов в буфере
dataqsiz — размерность буфера
buf — указатель на буфер для элементов канала
closed — флаг, указывающий, закрыт канал или нет
recvq — указатель на связанный список горутин, ожидающих чтения из канала
sendq -указатель на связанный список горутин, ожидающих запись в канал
lock — мьютекс для безопасного доступа к каналу

# mutex theory
https://habr.com/ru/post/271789/

# performance

https://tutorialedge.net/golang/benchmarking-your-go-programs/

# test
```go
package main

import (
	"fmt"
)

func main() {
  var s = []int{1,2,2,3,3,4,5}
  var ss = make(map[int]int)
  for _, el := range s {
    if ss[el] == 0 {
      fmt.Print(el)
    }
    ss[el]++;
  }

  var m = map[string]string{"john":"name", "bob":"name","Winston Ono":"middlename","Lennon":"surname","Dylan":"surname"}
  var r = make(map[string][]string)
  for key, value := range m {
    r[value] = append(r[value], key)
  }
  fmt.Println(m, r, s)
}
```
