package a

import "fmt"

func f() {
	// The pattern can be written in regular expression.
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}

func hoge(){
	fugafuga := 1
	hogehoge := 2
	sum := fugafuga + hogehoge
	fmt.Println(sum)
	sum2 := fugafuga +7
	fmt.Println(sum2)
}
