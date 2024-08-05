
# Estruturas Condicionais em Go

Estruturas condicionais são usadas para executar diferentes blocos de código com base em condições específicas. Em Go, as principais estruturas condicionais são `if`, `else if`, `else` e `switch`.

## `if` e `else`

A estrutura `if` é usada para executar um bloco de código se uma condição for verdadeira. A estrutura `else` pode ser usada para executar um bloco de código alternativo se a condição `if` for falsa.

### Sintaxe

```go
if condição {
    // código a ser executado se a condição for verdadeira
} else {
    // código a ser executado se a condição for falsa
}
```

### Exemplo

```go
package main

import "fmt"

func main() {
    idade := 18

    if idade >= 18 {
        fmt.Println("Você é maior de idade.")
    } else {
        fmt.Println("Você é menor de idade.")
    }
}
```

## `else if`

A estrutura `else if` permite verificar múltiplas condições. Se a primeira condição `if` for falsa, a condição `else if` será verificada, e assim por diante.

### Sintaxe

```go
if condição1 {
    // código a ser executado se a condição1 for verdadeira
} else if condição2 {
    // código a ser executado se a condição1 for falsa e condição2 for verdadeira
} else {
    // código a ser executado se todas as condições anteriores forem falsas
}
```

### Exemplo

```go
package main

import "fmt"

func main() {
    nota := 85

    if nota >= 90 {
        fmt.Println("A+")
    } else if nota >= 80 {
        fmt.Println("A")
    } else if nota >= 70 {
        fmt.Println("B")
    } else {
        fmt.Println("C ou menor")
    }
}
```

## `switch`

A estrutura `switch` é uma maneira mais limpa de escrever múltiplas condições que são baseadas no valor de uma única variável.

### Sintaxe

```go
switch expressão {
case valor1:
    // código a ser executado se expressão == valor1
case valor2:
    // código a ser executado se expressão == valor2
default:
    // código a ser executado se nenhum dos casos anteriores for atendido
}
```

### Exemplo

```go
package main

import "fmt"

func main() {
    dia := "segunda"

    switch dia {
    case "segunda":
        fmt.Println("Início da semana.")
    case "sexta":
        fmt.Println("Quase fim de semana.")
    case "sábado", "domingo":
        fmt.Println("Fim de semana!")
    default:
        fmt.Println("Dia normal.")
    }
}
```

## Conclusão

As estruturas condicionais são essenciais para controlar o fluxo do programa com base em diferentes condições. O `if`, `else if`, `else` e `switch` fornecem maneiras flexíveis e poderosas de tomar decisões em seu código Go.