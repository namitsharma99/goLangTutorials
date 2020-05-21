package main

import "fmt"

// the interface fix
type bot interface {
  getGreeting() string
}


type hindiBot struct {}
type englishBot struct {}

func main()  {
  hb := hindiBot{}
  eb := englishBot{}

  printGreeting(hb)
  printGreeting(eb)

}

//func (hb hindiBot) getGreeting() string {
func (hindiBot) getGreeting() string {  // since no hb value is read
  return "namaste!!"
}

// func (eb englishBot) getGreeting() string {
func (eb englishBot) getGreeting() string { // since no eb value is read
  return "hello!!"
}

// ERROR
/* func printGreeting(hb hindiBot) {
  fmt.Println(hb.getGreeting())
}

func printGreeting(eb englishBot) {
  fmt.Println(eb.getGreeting())
} */
// printGreeting redeclared in this block

func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}
