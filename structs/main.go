package main

import "fmt"

type contactInfo struct {
  email string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contactInfo
}



func main() {
  /* buddy := person {"code", "buddy"} // not a good practice as it depends on positions
  namit := person {firstName: "Namit", lastName: "Sharma"}

  fmt.Println(buddy)
  fmt.Println(namit)

  var human person
  fmt.Println(human)
  fmt.Printf("%+v", human)

  human.firstName = "ABC"
  human.lastName = "DEF"

  fmt.Println(human)
  fmt.Printf("%+v", human) */

  newPerson := person {
    firstName:"Namit",
    lastName:"Sharma",
    contactInfo: contactInfo {
      email: "thecodebuddy.blogspot.in",
      zipCode: 60001,
    },
  }
  // multiline declarations should have commas at the end

  fmt.Printf("%+v", newPerson)
  fmt.Println("")
  newPerson.print()
  fmt.Println("")
  fmt.Println("updating name::")
  newPersonPointer := &newPerson
  fmt.Println(&newPerson)
  newPersonPointer.updateName("XYZ")
  newPersonPointer.print()

}

func (p person) print() {
  fmt.Printf("%+v", p)
}

func (ptr *person) updateName(newFirstName string) {
  (*ptr).firstName = newFirstName
}
