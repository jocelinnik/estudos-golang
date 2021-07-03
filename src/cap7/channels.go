package main

import "fmt"

func produzir(c chan int){
	c <- 33
}

func channels1(){
	c := make(chan int)

	go produzir(c)

	valor := <- c
	fmt.Println(valor)
}

func main() {
	channels1()
}
