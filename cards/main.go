package main

import "fmt"

var num int

func main() {
  // var vari0 string = "Hello World"
  // vari1 := "Hello Again World"
  // fmt.Println(vari0)
  // fmt.Println(vari1)
  // fmt.Println(num)
  //
  // vari2 := newFunc()
  // fmt.Println(vari2)

  // varis := []string{"Hello0 ", newFunc()}
  // fmt.Println(varis)
  // varis = append(varis, "Hello2")
  // fmt.Println(varis)
  //
  // for i, vari := range varis {
  //   fmt.Println(i, vari)
  // }

  //cards := []string{"Ace of Diamonds", newCard()}
  cards := deck{"Ace of Diamonds", newCard()}
  cards = append(cards, "Six of Spades")

  for i, card := range cards {
    fmt.Println(i, card) // because cards is of type deck, hence print() is available
  }

  cards.print()

  fmt.Println("########## entire deck #########")

  cards = newDeck()
  for i, card := range cards {
    fmt.Println(i, card) // because cards is of type deck, hence print() is available
  }

  fmt.Println("########## sliced decks #########")
  hand, remainingCards := deal(cards, 5) // hand will have first sliced deck and remaining will go in the later

  fmt.Println("hand :: ")
  hand.print()
  fmt.Println("remainingCards :: ")
  remainingCards.print()



  /////

  stacks := stack{"stk1", "stk2", "stk3"}
  stacks = append(stacks, "stk4")

  for i, stack := range stacks {
    fmt.Println(i, stack)
  }

  //// Part 2 ////

  // byte slices

  temp := "Hello World!"
  fmt.Println(temp)
  fmt.Println([]byte(temp))

  cards = newDeck()
  fmt.Println(cards.toString())

  // saving //

  cards = newDeck()
  cards.saveToFile("helloCards")

  //  retrieving //
  cards = newDeckFromFile("helloCards")
  cards.print()

  // shuffle 1 //
  cards = newDeck()
  cards.shuffle()
  cards.print()

  // shuffle 2 //
  cards = newDeck()
  cards.shuffle2()
  cards.print()




}

func newCard() string {
  return "Five of Diamonds"
}

// func newFunc() string {
//   return "Hello1"
// }
