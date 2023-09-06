package Structs

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

// Para passar todos os tipos para outra struct, usar o nome da struct pra não ter que ficar dando varios pontos
/*
	type pessoa struct {
		nome string
		idade uint8
	}

	type universitario struct {
		pessoa                          // A struct recebe os campos nome e idade
		universidade string
	}
*/

func Struct_learning() {

	var u usuario

	u.nome = "Gabriel"
	u.idade = 21

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua x", 1000}

	u2 := usuario{"Thainá", 18, enderecoExemplo} //Precisa de todos os valores

	fmt.Println(u2)

	u3 := usuario{nome: "José"} //Caso não tenha todos os valores

	fmt.Println(u3)
}
