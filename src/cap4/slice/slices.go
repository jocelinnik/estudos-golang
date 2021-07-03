package main

import "fmt"

func slices1(){
	var a []int

	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}

	fmt.Println(a, primos, nomes)
}

func slices2(){
	b := make([]int, 10)
	c := make([]int, 10, 20)

	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}

func main(){
	slices1()
	slices2()
	//numeros := []int{1, 2, 3, 4, 5}
	//
	//for i := range numeros {
	//	numeros[i] *= 2
	//}
	//
	//fmt.Println(numeros)
}
