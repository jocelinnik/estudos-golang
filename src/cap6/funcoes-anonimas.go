package main

import (
	"fmt"
	"regexp"
	"strings"
)

func regexp1(){
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")

	fmt.Println(expr.ReplaceAllString(texto, "3"))
}

func regexp2(){
	texto := "antonio carlos jobim"
	expr := regexp.MustCompile("\\b\\w")

	processado := expr.ReplaceAllStringFunc(texto, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println(processado)
}

func regexp3(){
	expr := regexp.MustCompile("\\b\\w")

	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}

	texto := "antonio carlos jobim"

	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}

func main() {
	//regexp1()
	//regexp2()
	regexp3()
}
