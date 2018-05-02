package main
import (
    "fmt"
)

type person struct {
    name    string
    phone   string
}

func (p *person) show() {
    fmt.Printf("%s=%s", p.name, p.phone)
}

func readPerson() person {
    var name    string
    var phone   string

    fmt.Scan(&name)
    fmt.Scan(&phone)

    return person{name: name, phone: phone}
}

func printPerson(input *string, people []person) {
    for i:= 0; i < len(people); i++ {
        if people[i].name == *input {
            people[i].show()
            return
        }
    }

    fmt.Printf("Not found")
}

func main() {
    var times   int
    var people  []person
    var input   string

    fmt.Scan(&times)

    for i := 0; i < times; i++ {
        p := readPerson()
        people  = append(people, p)
    }

    previous := "foo"
    for {
        fmt.Scan(&input)
        if(previous == input) {
            break
        }

        printPerson(&input, people)
        fmt.Println()

        previous = input
    }
}
