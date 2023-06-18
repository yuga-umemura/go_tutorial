package main

import "fmt"

func main() {
	fmt.Println("counting")
	//deferへ渡した関数が複数ある場合は、その呼び出しはスタックされる。
	//呼び出し元の関数がreturnするとき、deferへ渡した関数はLIFOの順番で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}