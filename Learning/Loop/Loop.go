package Loop

import (
	"fmt"
	"time"
)

func Loop_learn() {

	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second) // Fazer esperar 1 segundo
	}

	/////////////////////////////////

	// J existe nesse escopo apenas
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second) // Fazer esperar 1 segundo
	}

	////////////////////////////////
	nomes := [3]string{"João", "Gabriel", "Felipe"}

	// Usado pra iterar em um maps, vetores e strings (struct não)
	// No caso do vetor, ele retorna 2 variaveis, indice e o valor. Para ocultar um valor, usar _
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Vai iterar palavra por palavra (numero na tabela ASCII caso não usar string(variavel))
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	/////////////////////////////
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	///////////////////////////
	for {
		fmt.Println("Infinito")
	}
}
