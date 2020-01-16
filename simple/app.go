package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()
    res := <-ch
    fmt.Println(res)
}
