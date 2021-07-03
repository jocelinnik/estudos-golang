package main

import "fmt"

func fatiarSlice1(){
	fib := []int{1, 1, 2, 3, 5, 8, 13}

	fmt.Println(fib)
	fmt.Println(fib[:3])
	fmt.Println(fib[2:])
	fmt.Println(fib[:])
}

func fatiarSlice2(){
	original := []int{1, 2, 3, 4, 5}
	novo := original[1:3]

	fmt.Println("Original: ", original)
	fmt.Println("Novo: ", novo)

	original[2] = 13

	fmt.Println("Original pós modificação: ", original)
	fmt.Println("Novo pós modificação: ", novo)
}

func fatiarSlice3(){
	a := []string{"Paulo", "almoça", "em", "casa", "diariamente."}
	b := a[:len(a) - 1]
	c := b[:len(b) - 1]
	d := c[:len(c) - 1]
	e := d[:len(d) - 1]

	e[0] = "Tiago"

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)
}

func main() {
	//fatiarSlice1()
	//fatiarSlice2()
	fatiarSlice3()
}
