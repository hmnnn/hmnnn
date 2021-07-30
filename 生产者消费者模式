package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c int = 0

func Product(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- rand.Intn(10)
	}
}
func Consumer(ch <-chan int) {
	for i := 0; i < 100; i++ {
		a := <-ch
		fmt.Println("Consmuer:", a)
	}
}
func reg() {
	c++
}

func main() {
	reg()
	reg()
	for i := 0; i < c; i++ {
		var ch = [5]chan int{}
		ch[i] = make(chan int, 1)
		go Product(ch[i])
		go Consumer(ch[i])
	}

	time.Sleep(time.Second)
}
