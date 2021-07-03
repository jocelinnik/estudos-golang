package main

import "fmt"

func ImprimirVersao(){
	fmt.Println("1.15")
}

func ImprimirSaudacao(nome string){
	fmt.Printf("Olá, %s\n", nome)
}

func ImprimirDados(nome string, idade int){
	fmt.Printf("%s, %d anos.\n", nome, idade)
}

func ImprimirSoma(a, b int){
	fmt.Println(a + b)
}

func ImprimirSoma2(a, b int, texto string){
	fmt.Printf("%s: %d\n", texto, a + b)
}

func Somar(a, b int) int {
	return a + b
}

func main() {
	ImprimirVersao()
}
