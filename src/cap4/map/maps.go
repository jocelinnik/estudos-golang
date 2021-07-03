package main

import "fmt"

func maps1(){
	vazio1 := map[int]string{}
	vazio2 := make(map[int]string)
	mapaGrande := make(map[int]string, 4096)

	fmt.Println(vazio1)
	fmt.Println(vazio2)
	fmt.Println(mapaGrande)
}

func maps2(){
	capitais := map[string]string{
		"GO": "Goiânia",
		"PB": "João Pessoa",
		"PR": "Curitiba",
	}

	capitais["RN"] = "Natal"
	capitais["AM"] = "Manaus"
	capitais["SE"] = "Aracajú"

	populacao := make(map[string]int, 6)
	populacao["GO"] = 6434052
	populacao["PB"] = 3914418
	populacao["PR"] = 10997462
	populacao["RN"] = 3373960
	populacao["AM"] = 3807923
	populacao["SE"] = 2228489

	fmt.Println(capitais)
	fmt.Println(populacao)
}

func maps3(){
	idades := map[string]int{
		"João": 37,
		"Ricardo": 26,
		"Joaquim": 41,
	}

	idades["Joaquim"] = 42

	fmt.Println(idades["Joaquim"])
}

func maps4(){
	quadrados := make(map[int]int, 15)

	for i := 1; i <= 15; i++ {
		quadrados[i] = i * i
	}

	for n, quadrado := range quadrados {
		fmt.Printf("%d^2 = %d\n", n, quadrado)
	}
}

func main() {
	//maps1()
	//maps2()
	//maps3()
	maps4()
}
