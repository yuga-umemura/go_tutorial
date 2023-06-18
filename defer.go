package main

import "fmt"

func main() {
	// deferへ渡した関数の実行を、呼び出し元の関数の終わり（return）まで遅延させる
	// deferへ渡した関数の引数は、すぐに評価される
	defer fmt.Println("world")

	fmt.Println("Hello")
}