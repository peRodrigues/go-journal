## Tipos de Dados em Go

Go é uma linguagem de programação estaticamente tipada, o que significa que os tipos das variáveis são conhecidos em tempo de compilação. Aqui estão os principais tipos de dados em Go:

### 1. Tipos Numéricos

#### Inteiros
- **int**: Tamanho depende da arquitetura (32 ou 64 bits).
- **int8, int16, int32, int64**: Inteiros de 8, 16, 32 e 64 bits respectivamente.
- **uint**: Unsigned int, tamanho depende da arquitetura (32 ou 64 bits).
- **uint8, uint16, uint32, uint64**: Unsigned integers de 8, 16, 32 e 64 bits respectivamente.
- **byte**: Alias para uint8.
- **rune**: Alias para int32, usado para representar um caractere Unicode.

#### Ponto Flutuante
- **float32**: Números de ponto flutuante de 32 bits.
- **float64**: Números de ponto flutuante de 64 bits.

#### Complexos
- **complex64**: Números complexos com partes real e imaginária como float32.
- **complex128**: Números complexos com partes real e imaginária como float64.

### 2. Booleanos

- **bool**: Pode ser `true` ou `false`.

### 3. Strings

- **string**: Uma sequência de bytes, geralmente usada para representar texto. As strings em Go são imutáveis.

### 4. Tipos Agregados

#### Arrays
- Um array é uma coleção de elementos do mesmo tipo, com um tamanho fixo.
```go
var arr [5]int
```

#### Slices
- Slices são mais flexíveis que arrays, podendo ter tamanho variável.
```go
var slice []int
```

#### Maps
- Mapas são coleções de pares chave-valor.
```go
var m map[string]int
```

#### Structs
- Structs são coleções de campos, usados para agrupar dados sob um único nome.
```go
type Person struct {
    Name string
    Age  int
}
```

### 5. Tipos de Ponteiros

- Ponteiros armazenam o endereço de memória de outra variável.
```go
var p *int
```

### Exemplos

#### Inteiros e Floats
```go
var a int = 10
var b int32 = 20
var c float32 = 5.5
var d float64 = 9.8
```

#### Booleanos
```go
var isGoFun bool = true
```

#### Strings
```go
var hello string = "Hello, Go!"
```

#### Arrays e Slices
```go
var arr [3]int = [3]int{1, 2, 3}
var slice []int = []int{1, 2, 3, 4, 5}
```

#### Maps
```go
var m map[string]int = map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

#### Structs
```go
type Person struct {
    Name string
    Age  int
}

var p Person = Person{Name: "John", Age: 30}
```

#### Ponteiros
```go
var x int = 10
var p *int = &x
```

### Tamanho
Utilizar um int8, ou um int64 pode não fazer muita diferença, certo? Errado! Em Go, é de suma importância escolher o tipo de dado adequado. Tendo os inteiros como exemplo, o int8 armazena inteiros de -128 a 127, diferentemente do int64, que armazena inteiros entre -9.223.372.036.854.775.808 e 9.223.372.036.854.775.807. Percebe a diferença? Isso mantém a claridade do código, além de reduzir erros e armazenamento. Escolher o tipo apropriado envolve considerar o intervalo necessário, eficiência de memória e processamento, compatibilidade e o significado semântico do tipo. Na prática, int é frequentemente usado por ser o tipo de inteiro padrão, por exemplo, mas usar tipos específicos pode ajudar a otimizar o uso de memória e melhorar a clareza do código.

### Conclusão

Go possui uma variedade de tipos de dados para lidar com diferentes necessidades de programação. Compreender esses tipos é fundamental para escrever programas eficientes e idiomáticos em Go.
