package Switch

import "fmt"

// Não presecisa do break, ele ja faz sozinho

func DiaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Invalido"
	}
}

/* Usado pra comparar mais de uma variavel
func DiaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Invalido"
	}
}
*/

/*
func DiaDaSemana3(numero int) string {
	var dia string

	switch {
	case numero == 1:
		dia = "Domingo"

		fallthrough //usada pra continuar o proximo case, muito raro

	case numero == 2:
		dia = "Segunda-feira"
	case numero == 3:
		dia = "Terça-feira"
	case numero == 4:
		dia = "Quarta-feira"
	case numero == 5:
		dia = "Quinta-feira"
	case numero == 6:
		dia = "Sexta-feira"
	case numero == 7:
		dia = "Sabado"
	default:
		dia = "Invalido"
	}

	return dia
}
*/

func Switch_learn() {

	dia := DiaDaSemana(2)
	fmt.Println(dia)
}
