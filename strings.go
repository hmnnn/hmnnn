package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "acbdw,1269547,AASIDX,AIUydjs,12sjaA,3819247,ausSHSzio,IUFISsi" //case 1
	s := strings.Split(str, ",")
	for i := 0; i < 8; i++ {
		a := sz(s[i])
		b := dx(s[i])
		c := xx(s[i])
		if a == 1 {
			fmt.Println(s[i] + "全是数字")
		} else if b == 1 {
			fmt.Println(s[i] + "全是大写字母")
		} else if c == 1 {
			fmt.Println(s[i] + "全是小写字母")
		} else {
			fmt.Println(s[i] + "既不是全是大写字母也不是全是小写字母也不是全是数字")
		}
	}

}
func sz(t string) int {
	for i := 0; i < len(t); i++ {
		if t[i] >= 48 && t[i] <= 57 {
			continue
		} else {
			return 0
		}

	}
	return 1
}

func dx(t string) int {
	for i := 0; i < len(t); i++ {
		if t[i] >= 65 && t[i] <= 90 {
			continue
		} else {
			return 0
		}

	}
	return 1
}

func xx(t string) int {
	for i := 0; i < len(t); i++ {
		if t[i] >= 97 && t[i] <= 122 {
			continue
		} else {
			return 0
		}

	}
	return 1
}
