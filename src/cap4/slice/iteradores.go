package main

import "fmt"

func iteradores1(){
	a, b := 0, 10

	for a < b {
		a += 1
	}

	fmt.Println(a)
}

func iteradores2(){
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func iteradores3(){
	var i int

	for i=0;i<10;i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(i)
}

func iteradores4(){
	i := 0
	for i < 10 {
		i += 1
	}

	for j := 0; j < 10; j++ {

	}

	var k int
	for k=0;k<10; {
		k += 1
	}

	l := 0
	for ; l < 10; l++ {

	}
}

func iteradores5(){
	numeros := []int{1, 2, 3, 4, 5}

	for i := range numeros {
		numeros[i] *= 2
	}

	fmt.Println(numeros)
}

func main() {
	//iteradores1()
	//iteradores2()
	//iteradores3()
	//iteradores4()
	iteradores5()
}
