## Arrays e Slices em Go

### Arrays

1. **Definição e Inicialização**:
   - Um array é uma coleção de elementos de um tipo específico, com um tamanho fixo que é definido no momento da declaração.
   - A sintaxe para declarar um array é:
     ```go
     var nomeDoArray [tamanho]tipo
     ```
   - Exemplo:
     ```go
     var números [5]int
     ```
   - Você pode inicializar um array ao declarálos:
     ```go
     var números = [5]int{1, 2, 3, 4, 5}
     ```

2. **Acesso e Modificação**:
   - Os elementos do array são acessados por índice, que começa em 0.
   - Exemplo:
     ```go
     números[0] = 10 // Modifica o primeiro elemento
     fmt.Println(números[0]) // Imprime 10
     ```

3. **Tamanho Fixo**:
   - O tamanho do array não pode ser alterado após a sua criação. Se você precisar de um tamanho variável, deve usar um slice.

### Slices

1. **Definição e Inicialização**:
   - Um slice é uma abstração sobre arrays que fornece uma maneira mais flexível de trabalhar com sequências de elementos.
   - Slices são dinâmicos e podem crescer ou encolher conforme necessário.
   - A sintaxe para criar um slice é:
     ```go
     var nomeDoSlice []tipo
     ```
   - Você pode criar um slice a partir de um array:
     ```go
     var números = [5]int{1, 2, 3, 4, 5}
     var slice = números[1:4] // Inclui os elementos dos índices 1 a 3
     ```

2. **Acesso e Modificação**:
   - Assim como arrays, você pode acessar e modificar elementos de um slice usando índices.
   - Exemplo:
     ```go
     slice[0] = 20 // Modifica o primeiro elemento do slice
     fmt.Println(slice[0]) // Imprime 20
     ```

3. **Capacidade e Comprimento**:
   - Um slice possui um comprimento (`len`) e uma capacidade (`cap`), que podem ser obtidos usando as funções `len()` e `cap()`.
   - Exemplo:
     ```go
     fmt.Println(len(slice)) // Comprimento do slice
     fmt.Println(cap(slice)) // Capacidade do slice
     ```

4. **Funções Úteis**:
   - A função `append()` é frequentemente usada para adicionar elementos a um slice:
     ```go
     slice = append(slice, 6) // Adiciona o elemento 6 ao final do slice
     ```

5. **Sub-slices**:
   - Você pode criar sub-slices a partir de um slice:
     ```go
     subSlice := slice[1:3] // Cria um novo slice que inclui elementos dos índices 1 a 2 do slice original
     ```

### Resumo

- **Arrays**: Tamanho fixo, mais rápido, mas menos flexível.
- **Slices**: Tamanho dinâmico, mais flexível, construído sobre arrays.
