# Estruturas (Structs) em Go

Estruturas (structs) em Go são usadas para agrupar variáveis sob um único tipo de dado. Elas são semelhantes às classes em outras linguagens de programação, mas com menos funcionalidades. Uma struct em Go é uma coleção de campos que são propriedades da estrutura.

## Definindo uma Struct

Você define uma struct usando a palavra-chave `type`, seguida pelo nome da struct e a palavra-chave `struct`. Dentro das chaves `{}`, você define os campos da struct, cada um com um nome e um tipo de dado.

```go
package main

import "fmt"

// Definindo uma struct chamada Person
type Person struct {
    Name   string
    Age    int
    Email  string
}

func main() {
    // Criando uma instância da struct Person
    person1 := Person{
        Name:  "Pedro",
        Age:   17,
        Email: "p.rodrigues1229@gmail.com",
    }

    // Acessando campos da struct
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
    fmt.Println("Email:", person1.Email)

    // Alterando um campo da struct
    person1.Age = 18
    fmt.Println("Updated Age:", person1.Age)
}
```
### Inicializando Structs

Existem várias maneiras de inicializar structs em Go:

    Inicialização com Valores: Especificando valores para todos os campos.

```go

person1 := Person{
    Name:  "Pedro",
    Age:   17,
    Email: "p.rodrigues1229@gmail.com",
}

Inicialização com Valores Padrão: Se você não especificar todos os campos, os não especificados serão inicializados com valores padrão (zero values).

go

person2 := Person{
    Name: "Rodrigues",
}
fmt.Println(person2) // Output: {Rodrigues 0 ""}
```

Inicialização com a Notação Posicional: Especificando valores na ordem dos campos definidos na struct (não é recomendado, pois pode levar a erros se a ordem dos campos mudar).

```go

    person3 := Person{"Rodrigues", 17, "p.rodrigues1229@gmail.com"}
```

## Métodos em Structs

Você pode definir métodos para structs para adicionar comportamentos a elas. Em Go, métodos são definidos associando uma função a uma struct.

```go

// Definindo um método para a struct Person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    person1 := Person{
        Name:  "Pedro",
        Age:   17,
        Email: "p.rodrigues1229@gmail.com",
    }
    person1.Greet() // Output: Hello, my name is Pedro and I am 17 years old.
}
```

## Structs Anônimas

Você também pode criar structs anônimas sem definir um tipo específico.

```go

func main() {
    anonymous := struct {
        Name string
        Age  int
    }{
        Name: "Anônimo",
        Age:  30,
    }

    fmt.Println(anonymous) // Output: {Anônimo 30}
}
```

## Embedding Structs

Go suporta um conceito chamado "embedding", que permite incluir uma struct dentro de outra, proporcionando uma forma de herança.

```go

type Address struct {
    City, State string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedding
}

func main() {
    person := Person{
        Name: "Pedro",
        Age:  17,
        Address: Address{
            City:  "São Paulo",
            State: "SP",
        },
    }

    fmt.Println(person) // Output: {Pedro 17 {São Paulo SP}}
}
```

## Conclusão

Structs são uma parte fundamental do Go, permitindo que você crie tipos complexos que combinam diferentes tipos de dados. Elas são usadas para representar entidades reais e adicionar métodos para fornecer comportamentos.