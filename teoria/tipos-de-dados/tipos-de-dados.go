package tiposdedados

import "fmt"

func Tiposdedados() {
	fmt.Println("Tipos de dados:")
	var inteiro int = 13
	var mininteiro int8 = 127
	fmt.Printf("Valor: %v\tTipo: %T\nValor: %v\tTipo: %T\n", inteiro, inteiro, mininteiro, mininteiro)

	texto := `"Hello, World"`
	fmt.Printf("Valor: %v\tTipo: %T\n", texto, texto)
}