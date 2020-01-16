package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player(table, "first")
	go player(table, "second")

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int, player string) {
	for {
		ball := <-table
		ball++
		fmt.Println(ball, player)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
