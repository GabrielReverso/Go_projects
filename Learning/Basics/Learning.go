package Learning

// go get ---> pega o pacote da internet e baixa
// go mod tidy ----> remove dependencias não usadas
// go mod vendor

import (
	"fmt"
)

/*
int ---> int8, int16, int32, int64 (Numero de bits, se for "int" ele vai usar o da arquitetura do pc)
uint (unsygned int, positivo apenas)
float (não usado explicitamente) ----> float32, float64
string
"char" (mostra o numero da tabela ASCII que representa o caractere. Não existe variavel char, para usar é só colocar '')
bool
error ---->  para criar erro, usar pacote interno, ex: var erro error = errors.New("Erro interno")
*/

/* Alias
INT32 = rune
UINT8 = byte
*/

// Funções de multiplos retornos
/*
	func Calculos(n1, n2 int8) (int8, int8) {
		soma := n1 + n2
		subtracao := n1 - n2
		return soma, subtracao
	}

	// Pra isso, precisa de 2 variaveis pra armazenar o resultado
	// Exemplo: resultado1, resultado2 := Calculos(1, 2)
	// Exemplo2: resultado1, _ := Calculos(1, 2) -----> _ serve pra ignorar um return respectivamente
*/

func Learning_go() {

	var a string // Declaração da variavel "a" do tipo string, se não for utilizada impedirá a compilação

	a = "Gabriel" // Atribuição da variável

	b := 1 // Declaração implicita da variavel (inferencia de tipos), o tipo será "adivinhado" pelo compilador, mas não poderá ser mudado depois

	c := false

	d := 3.1

	e := `Como vai
				você`

	/*
		var(
			f string = "oi"
			j string = "oi"
		)
	*/

	// Variavel1, Variavel2 := "oi", "oi"

	// const constante int = 2

	// Função como um tipo
	/*
		var f = func(txt string) {
			fmt.Println(txt)
		}
	*/

	fmt.Printf("%v \n", a) // %v mostrar variavel
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	fmt.Printf("%T \n", a) // %T mostrar o tipo
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

}
