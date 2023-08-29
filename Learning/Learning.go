package Learning

import (
	"fmt" // Package padrão com várias funções, se não for usado será apagado do import
)

var a string // Declaração da variavel "a" do tipo string, se não for utilizada impedirá a compilação

/*
	int
	string
	float
	bool
*/

func Learning_go() {

	a = "Gabriel" //Atribuição da variável

	b := 1 //Declaração implicita da variavel, o tipo será "adivinhado" pelo compilador, mas não poderá ser mudado depois

	c := false

	d := 3.1

	e := `Como vai
			você`

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
