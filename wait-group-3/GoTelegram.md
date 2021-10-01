```go
package main

import (
	"fmt"
)

func add(arr []int, v int) {
  arr = append(arr, v)
}

func main() {
	arr := make([]int, 0, 100000)
	fmt.Printf("%v %p\n", arr, &arr)
	add(arr, 10)
	fmt.Printf("%v %p\n", arr, &arr)
}
```

```
[] 0xc0000b4018
[] 0xc0000b4018
```

---

```go
package main

func f(...interface{}) {}

func main() {
  f(nil...)
  f([]int{1,2,3}...)
  f([]interface{}{1,2,3}...)
}
```
```
./prog.go:7:10: cannot use []int{...} (type []int) as type []interface {} in argument to f
```

---

What clause should we put instead ABC if we want to print all Caps letters firstly

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 26

func main() {
  const GOMAXPROCS = 1
  runtime.GOMAXPROCS(GOMAXPROCS)

  var wg sync.WaitGroup
  wg.Add(2 * N)
  for i := 0; i < N; i++ {
    go func(i int) {
      defer wg.Done()
      // ABC
      fmt.Printf("%c", 'a' + i)
    }(i)
    go func(i int) {
      defer wg.Done()
      fmt.Printf("%c", 'A' + i)
    }(i)
  }
  wg.Wait()
}
```

Options are:
* wg.Wait()
* runtime.Lock()
* runtime.Gosched (right answer)
* sleep(100)
* it's ok with the example

```
ZABCDEFGHIJKLMNOPQRSTUVWXYabcdefghijklmnopqrstuvwxyz
```

---

```go
package main

import (
	"fmt"
)

var o = fmt.Print

type A int
func (A) g() {}
func (A) m() int {return 1}

type B int
func (B) g() {}
func (B) f() {}

type C struct {A; B}
func (C) m() int {return 9}

func main() {
  var c interface{} = C{}
  _, bf := c.(interface{f()})
  _, bg := c.(interface{g()})
  i := c.(interface{m() int})

  fmt.Println(bf, bg, i.m())
}
```
```
true false 9
```

---

```go
package main

import (
	"fmt"
)

func main() {
  x := []int{123}
  x, x[0] = []int{432}, 456
  fmt.Println(x)
}
```
```
432
```

---

```go
package main

import (
	"fmt"
)

func main() {
  x := []int{123}
  x, x[0] = nil, 456
  fmt.Println(x)
}
```
```
[]
```

---

```go
package main

import (
	"fmt"
)

type T struct {
  _ int
  _ bool
}

func main() {
  var t1 = T{123, true}
  var t2 = T{789, false}
  fmt.Println(t1 == t2)
}
```
```
true
```

---

```go
package main

import (
	"fmt"
	"flag"
)

var port int

func init() {
  flag.IntVar(&port, "port", 8000, "port number")
}

func main() {
  flag.Parse()
  fmt.Printf("%d", port)
}
```

---

```go
package main

import (
	"fmt"
)

type Test interface {
  string() string
}

type TestImpl struct {
  sentence string
}

func (stat *TestImpl) string() string {
  return stat.sentence
}

func main() {
  var t Test = &TestImpl{"hello proghub"}
  fmt.Println(t)
}
```
```
&{hello proghub}
```

---

```go
package main

import (
	"fmt"
)

func main() {
  var a []int = nil
  a, a[0] = []int{1, 2}, 9
  fmt.Println(a)
}
```
```
panic: runtime error: index out of range [0] with length 0
```

---

```go
package main

import (
	"fmt"
)

var (
  _ = f("w", x)
  x = f("x", z)
  y = f("y", x)
  z = f("z")
)

func f(s string, deps ...int) int {
  fmt.Print(s)
  return 0
}

func main() {
   f("finish\n")
}
```
```
zxwyfinish
```

---

```go
package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4,5,6,7}
	// https://yourbasic.org/golang/three-dots-ellipsis/
	a2 := append(a, a[5:]...)
	fmt.Println(a2)
}
```

---

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
  var f func(int)
  {
    var f = func() {
      fmt.Printf("%T, ", f)
    }
    f()
  }

  type A = int
  {
    type A struct {*A}
    fmt.Println(reflect.TypeOf(A{}.A).Elem().Kind())
  }
}
```
```
func(int), struct
```

---

```go
package main

import (
	"fmt"
)

func main() {
   i := 0
   defer fmt.Println(i)
   i++
   return
}
```
```
0
```

---

```go
package main

import (
	"fmt"
)

type S struct {
  name string
}

func main() {
  m := map[string]S{"x": S{"one"}}
  m["x"].name = "two"
  fmt.Println(m["x"].name)
}
```
```
./prog.go:13:15: cannot assign to struct field m["x"].name in map
```

---

```go
package main

import (
	"fmt"
)

func main() {
  s := "qwe"
  ps := &s
  b := []byte(*ps)
  pb := &b

  s += "r"
  *ps += "t"
  *pb = append(*pb, []byte("y")[0])

  fmt.Println(*ps)
}
```
```
qwert
```

---

```go
package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
  m := make(map[int]int)

  wg := &sync.WaitGroup{}
  mu := &sync.Mutex{}
  wg.Add(N)

  for i := 0; i < N; i++ {
    go func() {
      defer wg.Done()
      mu.Lock()
      m[i] = i
      mu.Unlock()
    }()
  }

  wg.Wait()
  fmt.Println(len(m))
}
```
```
1
```
