package structs

import "fmt"

// Definindo uma struct chamada Person
type Person struct {
    Name   string
    Age    int
    Email  string
}

// Definindo um método para a struct Person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Definindo uma struct chamada Address
type Address struct {
    City, State string
}

// Definindo uma struct chamada Employee que incorpora Person e Address
type Employee struct {
    Person  // Embedding Person
    Address // Embedding Address
    JobTitle string
}

func Structs() {
	fmt.Println("Structs:")
    // Criando uma instância da struct Employee
    employee := Employee{
        Person: Person{
            Name:  "Pedro",
            Age:   17,
            Email: "p.rodrigues1229@gmail.com",
        },
        Address: Address{
            City:  "São Paulo",
            State: "SP",
        },
        JobTitle: "Estagiário",
    }

    // Acessando campos da struct Employee
    fmt.Println("Name:", employee.Name)
    fmt.Println("Age:", employee.Age)
    fmt.Println("Email:", employee.Email)
    fmt.Println("City:", employee.City)
    fmt.Println("State:", employee.State)
    fmt.Println("Job Title:", employee.JobTitle)
	fmt.Println("Address:", employee.Address)

    // Chamando o método Greet da struct Person
    employee.Greet() // Output: Hello, my name is Pedro and I am 17 years old.
}