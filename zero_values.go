//https://pkg.go.dev/fmt

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	//%qだとダブルクォートで囲まれた文字列を出力できる
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Printf("%v %v %v %s\n", i, f, b, s)
}


