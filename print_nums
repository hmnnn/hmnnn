package main

import (
	"fmt"
	"time"
)

var POOL = 100

func groutine1(p chan int) {

	for i := 1; i <= POOL; i++ {
		p <- i
		if i%10 == 1 {
			fmt.Println("groutine-1:", i, i+2, i+4, i+6, i+8)
		}
	}
}

func groutine2(p chan int) {

	for i := 1; i <= POOL; i++ {
		<-p
		if i%10 == 2 {
			fmt.Println("groutine-2:", i, i+2, i+4, i+6, i+8)
		}
	}
}

func main() {
	msg := make(chan int) //只能在类型位map或chan的场景
	//管道缓冲区依据缓冲区容量被初始化。如果容量为 0 或者忽略容量，管道是没有缓冲区的

	go groutine1(msg)
	go groutine2(msg)

	time.Sleep(time.Second * 1)

}
