package main

import "fmt"

//初期化子が与えられている場合、型を省略できる
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}