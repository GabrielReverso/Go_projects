package Map

import "fmt"

func Map_learning() {

	//map[tipo_chave]tipo_dado
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"ultimo":   "Reverso",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome") //Remover chave
	fmt.Println(usuario2)

	/* adicionando outra chave
	usuario2["signo"] = map[string]string {
		"nome": "gemeos"
	}
	*/
}
