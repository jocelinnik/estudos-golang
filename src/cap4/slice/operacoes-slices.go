package main

import "fmt"

func inserindovalores1(){
	s := make([]int, 0)

	s = append(s, 23)

	fmt.Println(s)
}

func inserindovalores2(){
	s := []int{23, 24, 25}

	s = append([]int{22}, s...)

	fmt.Println(s)
}

func inserindovalores3(){
	s := []int{11, 12, 13, 16, 17, 18}
	v := []int{14, 15}

	s = append(s[:3], append(v, s[3:]...)...)

	fmt.Println(s)
}

func removendovalores1(){
	s := []int{20, 30, 40, 50, 60}

	s = s[1:]

	fmt.Println(s)
}

func removendovalores2(){
	s := []int{20, 30, 40, 50, 60}

	s = s[:3]

	fmt.Println(s)
}

func removendovalores3(){
	s := []int{10, 20, 30, 40, 50, 60}

	s = append(s[:2], s[4:]...)

	fmt.Println(s)
}

func copiandovalores1(){
	numeros := []int{1, 2, 3, 4, 5}
	dobros := make([]int, len(numeros))

	copy(dobros, numeros)

	for i := range dobros {
		dobros[i] *= 2
	}

	fmt.Println(numeros)
	fmt.Println(dobros)
}

func main() {
	//inserindovalores1()
	//inserindovalores2()
	//inserindovalores3()
	//removendovalores1()
	//removendovalores2()
	//removendovalores3()
	copiandovalores1()
}
