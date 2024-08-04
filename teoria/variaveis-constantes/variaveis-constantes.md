# Variáveis e Constantes em Golang

Em Go, variáveis e constantes são fundamentais para armazenar e manipular dados. Vamos explorar cada uma delas em detalhes.

## Variáveis

As variáveis em Go são usadas para armazenar valores que podem mudar durante a execução do programa. Para declarar uma variável, você pode usar a palavra-chave `var` ou o operador curto `:=`.

### Sintaxe de declaração com `var`

```go
var nome string
var idade int
```

### Sintaxe de inicialização com `var`

```go
var nome string = "Pedro"
var idade int = 17
```

### Operador curto `:=`

Essa forma é muito utilizada em Go para declarar e inicializar variáveis de uma só vez, mas só pode ser usada dentro de funções.

```go
nome := "Pedro"
idade := 17
```

### Tipos de Variáveis

Go é uma linguagem fortemente tipada, o que significa que cada variável deve ter um tipo específico. Alguns tipos básicos são:
- `int`: para números inteiros
- `float64`: para números de ponto flutuante
- `string`: para cadeias de caracteres
- `bool`: para valores booleanos (true ou false)

## Constantes

As constantes são semelhantes às variáveis, mas, uma vez definidas, não podem ser alteradas. Elas são usadas para armazenar valores que você sabe que não vão mudar durante a execução do programa.

### Declaração de constantes

Para declarar uma constante, usa-se a palavra-chave `const`.

```go
const pi float64 = 3.14159
const nome = "Rodrigues" // O tipo é inferido como string
```

### Inferência de tipo

Go também permite a inferência de tipo em constantes.

```go
const saudacao = "Olá, mundo!"
```

## Exemplos práticos

Vamos ver um exemplo prático que usa variáveis e constantes em um pequeno programa:

```go
package main

import "fmt"

func main() {
    // Declaração e inicialização de variáveis
    var nome string = "Rodrigues"
    var idade int = 17

    // Uso do operador curto para declarar e inicializar
    altura := 1.75
    programador := true

    // Declaração de constantes
    const pi float64 = 3.14159
    const saudacao = "Olá, mundo!"

    // Imprimindo valores
    fmt.Println(saudacao)
    fmt.Println("Nome:", nome)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("É programador?", programador)
    fmt.Println("Valor de Pi:", pi)
}
```

## Resumo

- **Variáveis**: Usadas para armazenar valores que podem mudar. Declaradas com `var` ou `:=`.
- **Constantes**: Usadas para armazenar valores imutáveis. Declaradas com `const`.