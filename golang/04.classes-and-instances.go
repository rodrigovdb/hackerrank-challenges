package main
import "fmt"

type person struct {
  age int
}

func (p person) NewPerson(initialAge int) person {
  if (initialAge < 0) {
    fmt.Println("Age is not valid, setting age to 0.");
    initialAge = 0;
  }

  p.age = initialAge;

  return p
}

func (p person) amIOld() {
  //Do some computation in here and print out the correct statement to the console

  var label string

    if p.age < 13 {
        label = "young";
    } else if p.age < 18 {
        label = "a teenager";
    } else {
        label = "old";
    }

  fmt.Printf("You are %s.\n", label);
}

func (p person) yearPasses() person {
  //Increment the age of the person in here

  p.age += 1;

  return p
}

func main() {
  var T,age int

  fmt.Scan(&T)

  for i := 0; i < T; i++ {
    fmt.Scan(&age)
    p := person{age: age}
    p = p.NewPerson(age)
    p.amIOld()
    for j := 0; j < 3; j++ {
      p = p.yearPasses()
    }
    p.amIOld()
    fmt.Println()
  }
}
