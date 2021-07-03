package main

import "fmt"

func produzir2(c chan int){
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func buffers1(){
	c := make(chan int, 3)

	go produzir2(c)

	for valor := range c {
		fmt.Println(valor)
	}
}

func main() {
	buffers1()
}
