package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	//structのフィールドはドットでアクセス
	v.X = 4
	fmt.Println(v.X)
	//ポインタを通じてアクセスすることも可能
	//&オペレータでvのオペランドへのポインタを渡す
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}