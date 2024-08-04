package exercicios

import (
	"fmt"
)

type y int
var x y

func Ex004() {
	fmt.Println("\nex004:")
	// Cria um tipo baseado em int, declara uma variável desse tipo, demonstra seu valor e tipo, atribui 42 a ela e mostra o valor atualizado.
	fmt.Println("Valor de x:", x)
	fmt.Printf("Tipo de x: %T\n", x)
	x = 13
	fmt.Println("Novo valor de x:", x)


}