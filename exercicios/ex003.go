package exercicios

import "fmt"

var nome string = "Pedro"
var idade uint8 = 17
var programador bool = true

func Ex003() {
	fmt.Println("\nex003:")
	// Em package-level scope, atribua valores às variáveis, use fmt.Sprintf na função main para combinar esses valores em uma string, e demonstre a variável resultante "s".
	s := fmt.Sprintf("Meu nome é %v e tenho %v anos. Quero me tornar programador: %v.", nome, idade, programador) // fmt.Sprintf formata e retorna uma string
	fmt.Println(s)
}