package Array

import (
	"fmt"
	"reflect"
)

// Array é fixo, inflexivel e tipado

func Array_learning() {

	fmt.Println("Array")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6} //Os 3 pontos servem pra fixar o tamanho de acordo com a quantia de itens
	fmt.Println(array3)
	fmt.Println(reflect.TypeOf(array3))

}

// Slice possui flexibilidade no tamanho, porem ainda tipado
// "aponta" para um array

func Slice_learning() {

	fmt.Println("Slice")

	slice := []int{10, 11, 12}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18) // Pega o slice criado anteriormente e cria um novo com uma nova posição

	array4 := [5]int{1, 2, 3, 4, 5}
	slice2 := array4[1:3] // Cria um slice de um array de um indice a outro (inclusivo), no caso do 1 até o 3. Funciona como um ponteiro
	fmt.Println(slice2)

	// Array interno
	fmt.Println("------")
	slice3 := make([]float32, 10, 11) // Tipo, Tamanho e Capacidade maxima. No caso, esse slice pega as 10 primeiras casas das 11
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Mostra tamanho
	fmt.Println(cap(slice3)) // Mostra capacidade maxima do array

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Mostra tamanho
	fmt.Println(cap(slice3)) // Mostra capacidade maxima do array

	// Ao estourar a capacidade maxima do array interno, o Go cria um novo array com o dobro da capacidade
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Mostra tamanho
	fmt.Println(cap(slice3)) // Mostra capacidade maxima do array

	// O ultimo parametro do make é opcional. Ao ser ocultado, a capacidade será a mesma que a quantidade de casas
}
