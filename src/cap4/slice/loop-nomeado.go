package main

import "fmt"

func loopnomeado1(){
	var i int

	externo:
	for {
		for i=0;i<10;i++ {
			if i == 5 {
				break externo
			}
		}
	}

	fmt.Println(i)
}

func loopnomeado2(){
	var i int

	loop:
	for i=0;i<10;i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d.\n", i)
			break
		case 5:
			fmt.Println("Quebrando loop, i == 5.")
			break loop
		}
	}

	fmt.Println("Fim")
}

func main() {
	//loopnomeado1()
	loopnomeado2()
}
