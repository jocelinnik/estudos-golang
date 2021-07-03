package main

import "fmt"

func imprimirDados(nome string, idade int){
	fmt.Printf("%s tem %d anos.\n", nome, idade)
}

func soma(m, n int) int {
	return n + m
}

func main(){
	// imprimirDados("Linnik", 23)
	s := soma(3, 5)
	fmt.Println("A soma é", s)
}
