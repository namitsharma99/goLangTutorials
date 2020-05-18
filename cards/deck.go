package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "os"
  "math/rand"
  "time"
)
// creating a new type of deck - slice of string

type deck[]string

func newDeck() deck {
  cards := deck{} // how to cover all combo
  cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string {"Ace", "Two", "Three", "Four", "Five" , "Six", "Seven", "Eight", "Nine" , "Ten", "Jack", "Queen", "King" }

  for _, suit := range cardSuits { // instead of i, j use _ so that no compulsion to use
    for _, value := range cardValues {
      cards = append(cards, value + " of " + suit)
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

func (d deck) toString() string{
  // use strings package
  return strings.Join([]string(d),",")
}

func (d deck) saveToFile(fileName string) error {
  return ioutil.WriteFile(fileName, []byte(d.toString()), 0777) // provide permissions
}


func newDeckFromFile(fileName string) deck {
  bytslc, err := ioutil.ReadFile(fileName) // default error handling
  if err != nil {
    // Option 1 - log error and return a new decks
    // Option 2 - log error and quit
    fmt.Println("Error: ", err)
    os.Exit(1) // imported package os
  }
  s := strings.Split(string(bytslc), ",") // typecasting of byte slice to string, then using split to chunk on the basis of comma
  return deck(s) // using slice of strings to create a deck
}


func (d deck) shuffle() {
  for i := range d {
    newPosition := rand.Intn(len(d)-1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}

func (d deck) shuffle2() {
  source := rand.NewSource(time.Now().UnixNano()) // seed
  r := rand.New(source) // generator

  for i := range d {
    newPosition := r.Intn(len(d)-1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
