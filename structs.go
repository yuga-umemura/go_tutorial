package main

import "fmt"

//structs(構造体)はフィールドの集まり
type Vertex struct {
	X int //これがフィールド
	Y int
}

func main() {
	fmt.Println(Vertex{1,  2})
}