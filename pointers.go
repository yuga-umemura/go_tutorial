package main

import "fmt"

func main() {
	i, j := 42, 2701

    //&オペレータはそのオペランドへのポインタを引き渡す
	p := &i         // point to i
    //*オペレータは、ポインタの指す先の変数を示す
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
