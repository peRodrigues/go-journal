package tiposdedados

import "fmt"

func Tiposdedados() {
	fmt.Println("Tipos de dados:")
	// Declaração completa
	var inteiro int = 13
	var mininteiro int8 = 127
	fmt.Printf("Valor: %v\tTipo: %T\nValor: %v\tTipo: %T\n", inteiro, inteiro, mininteiro, mininteiro)

	// Declaração da variável com tipo inferido como string
	texto := `"Hello, World"`
	fmt.Printf("Valor: %v\tTipo: %T\n", texto, texto)
}