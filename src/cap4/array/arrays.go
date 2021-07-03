package main

import "fmt"

func arrays1(){
	var a [3]int
	var nomes [2]string

	numeros := [5]int{1, 2, 3, 4, 5}
	primos := [...]int{2, 3, 5, 7,11, 13}

	fmt.Println(a, numeros, primos, nomes)
	fmt.Println(len(a), len(numeros), len(primos), len(nomes))
}

func arrays2(){
	var multiA [2][2]int

	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("Multi A: ", multiA)
	fmt.Println("Multi B: ", multiB)
}

func main() {
	arrays1()
	arrays2()
}
