package Ponteiros

import "fmt"

//Ponteiro é uma referencia de memoria

func Ponteiros_learning() {
	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)

	fmt.Println(variavel, *ponteiro)

	fmt.Println(&variavel, *ponteiro)
}
