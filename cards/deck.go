package main

import "fmt"
// creating a new type of deck - slice of string

type deck[]string

func newDeck() deck {
  cards := deck{"Ace of Spades"} // how to cover all combo
  cardSuits := []string {"Spades", "Diamonds", "Hearts", "Club"}
  cardValues := []string {"Ace", "Two", "Three", "Four", "Five" , "Six", "Seven", "Eight", "Nine" , "Ten", "Jack", "Queen", "King" }

  for _, suit := range cardSuits { // instead of i, j use _ so that no compulsion to use
    for _, value := range cardValues {
      cards = append(cards, suit + " of " + value)
    }
  }

  return cards
}

func deal (d deck, handsize int) (deck, deck) { // multiple returns of type deck each
  return d[:handsize], d[handsize:]
}

func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
