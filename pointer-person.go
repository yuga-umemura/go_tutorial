package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	//ポインタ型変数宣言→ポインタを変数とするとき「ポインタ型変数」と呼ぶ？
	//*Personポインタ型→「*」をつける？
	//↓
	//&を使うことで、ポインタ型を生成することができる
	//Person型の変数pを&pとすると、Personへのポインタである*Person型を生み出すことができる
	//&pはpのアドレスという
	//https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
	var p *Person

	p = &Person {
		Name: "太郎",
		Age: 20,
	}
	fmt.Printf("変数pに格納されているアドレス :%p", p)
}
//&オペレータはそのオペランドへのポインタを引き渡す
//*オペレータは、ポインタの指す先の変数を示す