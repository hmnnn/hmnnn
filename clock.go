package main

import (
	"fmt"
	"time"
)

type Shizhong struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

func (a Shizhong) show() {
	fmt.Println(a)
}

func main() {
	var a Shizhong
	a.year = 2002
	a.month = 2
	a.day = 12

	for true {
		a.hour = time.Now().Hour()
		a.minute = time.Now().Minute()
		a.second = time.Now().Second()
		a.show()
		time.Sleep(time.Second)
	}
}
