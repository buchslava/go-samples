package main

import (
  "fmt"
  "sync"
  // "time"
)

func main() {
  var wg sync.WaitGroup
  sayHello := func() {
    defer wg.Done()
    fmt.Println("hello")
  }
  wg.Add(1)
  go sayHello()
  wg.Wait()
}
