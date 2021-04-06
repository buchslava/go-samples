# timeout
```go
package main

import (
	"fmt"
	"time"
)

func sleep(s int) {
  <-time.After(time.Second * time.Duration(s))
}

func main() {
  sleep(5)
  fmt.Println("Hello, playground")
}
```

# buffer
```go
package main

import "fmt"

func main() {
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"
    // messages <- "foo"
 
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    messages <- "foo"
    fmt.Println(<-messages)
}
```
```
buffered
channel
foo
```
# deadlock for buffered channel
```go
package main

import "fmt"

func main() {
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"
    messages <- "foo" // !!!
 
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox157669183/prog.go:10 +0x8d
```

# unbuffered channel issue
```go
package main

func main() {
    messages := make(chan string)
    messages <- "buffered"
}
```
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox653494253/prog.go:5 +0x50
```

# unbuffered channel without error
```go
package main

import (
  "fmt"
)

func main() {
  ch01 := make(chan string)

  go func() {
    fmt.Println(<-ch01)
  }()

  ch01 <- "Hello"
}
```

# swap
```go
package main

import "fmt"

func main()  {
	x := 1
	y := 2

	swap(&x, &y)

	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
```
# square
```go
package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 1.5
	square(&x)

	fmt.Println(x)
}
```

# rune
```go
package main

import "fmt"

func main() {
   i := 65
   fmt.Println(string(i))
}
```

# ???
```go
package main

import "fmt"

func main() {
   a := [5]int{1, 2, 3, 4, 5}
   t := a[3:4:4]
   fmt.Println(t[0])
}
```

# map iterate
```go
package main

import "fmt"

func main() {
    var employee = map[string]int{"Mark": 10, "Sandy": 20,
        "Rocky": 30, "Rajiv": 40, "Kate": 50}
    for key, element := range employee {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }
}
```

# slice zero value
```go
package main

import "fmt"

func main() {
    var s *struct{} // uninitialised pointer
    fmt.Println(s)  // <nil>

    s = &struct{}{} // pointer to an empty structure
    fmt.Println(s) 
}
```

# append - capacity
```go
package main

import "fmt"

type str3Bytes struct {
	a byte
	b byte
	c byte
}

func main() {

	arr := []struct{}{}
	oldCap := 0
	for i := 0; i < 100; i++ {
		arr = append(arr, struct{}{})
		if cap(arr) != oldCap {
			oldCap = cap(arr)
			fmt.Println("arr", len(arr), cap(arr))
		}
	}

	arrInt := []int{}
	oldCap = 0
	for i := 0; i < 100; i++ {
		arrInt = append(arrInt, i)
		if cap(arrInt) != oldCap {
			oldCap = cap(arrInt)
			fmt.Println("arrInt", len(arrInt), cap(arrInt))
		}
	}

	arrStr := []string{}
	oldCap = 0
	for i := 0; i < 100; i++ {
		arrStr = append(arrStr, fmt.Sprint(i))
		if cap(arrStr) != oldCap {
			oldCap = cap(arrStr)
			fmt.Println("arrStr", len(arrStr), cap(arrStr))
		}
	}

	arr3Bytes := []str3Bytes{}
	oldCap = 0
	for i := 0; i < 100; i++ {
		arr3Bytes = append(arr3Bytes, str3Bytes{1, 2, 3})
		if cap(arr3Bytes) != oldCap {
			oldCap = cap(arr3Bytes)
			fmt.Println("arr3Bytes", len(arr3Bytes), cap(arr3Bytes))
		}
	}
}
```

# close channel
```go
package main

import "fmt"

func main() {

    // Будем перебирать 2 значения в канале `queue`.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // Этот `range` проходит по каждому элементу, когда тот
    // получен из `queue`. Поскольку мы закрыли канал выше,
    // перебор завершается после получения 2 элементов.
    // Если бы мы не закрыли канал, получили бы блокировку
    // на 3 попытке приёма в цикле.

    for elem := range queue {
        fmt.Println(elem)
    }
}
```

# timer
```go
package main

import "time"
import "fmt"

func main() {

    // Таймер представляет из себя одиночное событие в будущем.
    // Вы говорите таймеру как долго нужно подождать и он
    // обеспечивает канал, который будет оповещён в указанное
    // время. Этот таймер будет ждать 2 секунды.
    timer1 := time.NewTimer(time.Second * 2)

    // `<-timer1.C` блокирует канал таймера `C`
    // до тех пор, пока он не отправит значение, означающее,
    // что вышел срок таймера.
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // Если нужно просто подождать, можно использовать
    // `time.Sleep`. Причина, по которой может быть полезен
    // таймер в том, что можно отменить таймер до окончания
    // его срока. Вот пример отмены.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
```
```
Timer 1 expired
Timer 2 stopped
```

# ticker
```go
package main

import "time"
import "fmt"

func main() {

    // Счетчик тиков использует механизм, похожий на
    // таймер: канал, который отправляет значения. Здесь мы
    // используем `range`, встроенный в канал для перебора
    // значений, поступающих каждые 500 мсек.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // Счетчики тиков могут быть остановлены подобно таймерам.
    // Как только счетчик будет остановлен, больше не будет
    // получать значения в канале. Остановим его после 1500 мсек.
    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
```
```
Tick at 2009-11-10 23:00:00.5 +0000 UTC m=+0.500000001
Tick at 2009-11-10 23:00:01 +0000 UTC m=+1.000000001
Ticker stopped
```

# workers
```go
package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    for a := 1; a <= 9; a++ {
        x := <-results
        fmt.Println(x)
    }
}
```

# collect functs
```go
package main

import "strings"
import "fmt"

// Возвращает первый индекс совпадения со строкой `t` или
// -1 если совпадение не найдено.
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Возвращает `true` если строка `t` находится в слайсе
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Возвращает `true` если одна из строк в слайсе
// удовлетворяет условие `f`
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// Возвращает `true` если все из строк в слайсе
// удовлетворяют условие `f`
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Возвращает новый слайс, содержащий все строки в
// слайсе, которые удовлетворяют условие `f`
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Возвращает новый слайс, содержащий результаты выполнения
// функции `f` с каждой строкой исходного слайса
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))
    fmt.Println(Map(strs, strings.ToUpper))
}
```

# strings
```go
package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
```

# sort struct
```go
package main

import "log"
import "sort"

// AxisSorter sorts planets by axis.
type AxisSorter []Planet

func (a AxisSorter) Len() int           { return len(a) }
func (a AxisSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AxisSorter) Less(i, j int) bool { return a[i].Axis < a[j].Axis }

// NameSorter sorts planets by name.
type NameSorter []Planet

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

type Planet struct {
    Name       string  `json:"name"`
    Aphelion   float64 `json:"aphelion"`   // in million km
    Perihelion float64 `json:"perihelion"` // in million km
    Axis       int64   `json:"Axis"`       // in km
    Radius     float64 `json:"radius"`
}

func main() {
    var mars Planet
    mars.Name = "Mars"
    mars.Aphelion = 249.2
    mars.Perihelion = 206.7
    mars.Axis = 227939100
    mars.Radius = 3389.5

    var earth Planet
    earth.Name = "Earth"
    earth.Aphelion = 151.930
    earth.Perihelion = 147.095
    earth.Axis = 149598261
    earth.Radius = 6371.0

    var venus Planet
    venus.Name = "Venus"
    venus.Aphelion = 108.939
    venus.Perihelion = 107.477
    venus.Axis = 108208000
    venus.Radius = 6051.8

    planets := []Planet{mars, venus, earth}
    log.Println("unsorted:", planets)

    sort.Sort(AxisSorter(planets))
    log.Println("by axis:", planets)

    sort.Sort(NameSorter(planets))
    log.Println("by name:", planets)
}
```

# reverse linked list
```go
package main

import "fmt"

type Item struct { 
    pNext *Item
    val rune
}

func createList() *Item {

    pHead := &Item{nil, 'a'}
    
    pCurr := pHead
    for i:= 'b'; i <= 'z'; i++ { 
        pItem := &Item{nil, i}       
        pCurr.pNext = pItem
        pCurr = pItem
    }
    
    return pHead
}

func printList(pList *Item) {

    pCurr := pList
    for {
        fmt.Printf("%c", pCurr.val)
        
        if pCurr.pNext != nil {
            pCurr = pCurr.pNext
        } else {
            break
        }
    }
    fmt.Println("")
}

func reverseList(pList *Item) *Item {
    
    pCurr := pList
    var pTop *Item = nil
    for {
        if pCurr == nil {
            break
        }
        pTemp := pCurr.pNext
        pCurr.pNext = pTop
        pTop = pCurr
        pCurr = pTemp        
    }
    
    return pTop
}

func main() {

    var pList = createList() 
    printList(pList)
    printList(reverseList(pList))
}
```

# inh & compos
```go
package main

import (
	"fmt"
)

type Walker struct {
	Legs int
}

func (w Walker) Walk() {
	fmt.Println("Walking started with ", w.Legs, " legs")
}

type Eater struct {}

func (e Eater) Eat() {
	fmt.Println("Started eating")
}

type Sleeper struct {}

func (s Sleeper) Sleep() {
	fmt.Println("Started sleeping")
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Walker
	Sleeper
}

type Cat struct {
	Animal
	Walker
	Eater
}

func main() {
	d := new(Dog)
	d.Name = "Bruce"
	d.Legs = 4
	d.Walk()
	d.Sleep()
	
	c := new(Cat)
	c.Name = "Cretin"
	c.Legs = 3
	c.Walk()
	c.Eat()
}
```

# efficiently string concat
```go
package main

import (
    "strings"
    "fmt"
)

func main() {
    // ZERO-VALUE:
    //
    // It's ready to use from the get-go.
    // You don't need to initialize it.
    var sb strings.Builder

    for i := 0; i < 1000; i++ {
        sb.WriteString("a")
    }

    fmt.Println(sb.String())
}
```

# efficiently string concat - old version
```go
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
    }

    fmt.Println(buffer.String())
}
```

# GOMAXPROCS
```go
package main

import (
  "runtime"
  "fmt"
)

func main() {
   runtime.GOMAXPROCS(1)

   done := false

   fmt.Println("1")

   go func() {
      fmt.Println("2")
      done = true
   }()

   fmt.Println("3")


   for !done {
     fmt.Println("4")
     // runtime.Gosched()
   }
   fmt.Println("finished")
}
```

# simple
```go
package main

import (
  "fmt"
)

func main() {
   c := [4]int{1,2,3,4}
   fmt.Println(c)
   fmt.Println(c[1:3])
   for n, data := range(c) {
      fmt.Println(n, data);
   }
}
```
# select-chan
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```
```
received one
received two
```

# limited parallel
```go
package main

import (
	"fmt"
	"sync"
	"time"
)


var (
    jobs    = 20                 // Run 20 jobs in total.
    running = make(chan bool, 3) // Limit concurrent jobs to 3.
    wg      sync.WaitGroup       // Keep track of which jobs are finished.
)

func main() {
  wg.Add(jobs)
  for i := 1; i <= jobs; i++ {
    running <- true // Fill running; this will block and wait if it's already full.

    go func(i int) {
        defer func() {
            <-running // Drain running so new jobs can be added.
            wg.Done() // Signal that this job is done.
        }()

        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }(i)
  }

  wg.Wait()
  fmt.Println("done")
}
```

# slice
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

# chan deadlock
```go
package main

import (
  "fmt"
)

func main() {
  ans := make(chan int)
  go iter(ans)
  for {
    select {
      case <-ans:
        fmt.Print(<-ans)
    }
  }
}

func iter(ans chan int) {
  for i := 0; i < 10; i++ {
    ans <- i
  }
}
```
```
13579fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox732557531/prog.go:12 +0x73
```
