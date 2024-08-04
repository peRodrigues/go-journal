package variaveisconstantes

import (
    "fmt"
    "math"
)

func Variaveisconstantes() {
	fmt.Println("Variáveis e Constantes:")
    // Declaração de variáveis
    var idade int32 = 17
    altura := 1.75
    programador := true

    // Declaração de constantes
    const pi float64 = 3.14159
    const saudacao = "Olá, mundo!"
    const maxInt32 int32 = math.MaxInt32
    const minInt32 int32 = math.MinInt32

    // Imprimindo valores
    fmt.Println(saudacao)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("É programador?", programador)
    fmt.Println("Valor de Pi:", pi)
    fmt.Println("Valor máximo de int32:", maxInt32)
    fmt.Println("Valor mínimo de int32:", minInt32)
}