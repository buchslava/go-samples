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
        close(c1)
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
        close(c2)
    }()

    for i := 0; i < 12; i++ {
        select {
        case msg1, more := <-c1:
            fmt.Println("received from 1 is ", msg1, more)
        case msg2, more := <-c2:
            fmt.Println("received from 2 is ", msg2, more)
        default: fmt.Print(".")
            time.Sleep(1 * time.Second)
       //case <-time.After(time.Second * 1):
       //fmt.Print(".")
      }
    }
}
