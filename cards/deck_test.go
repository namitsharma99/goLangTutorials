package main

import "testing"
import "os"

/* func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 52 {
    t.Errorf("Expected total 52, but got %v", len(d))
  }

  if (d[0] != "Ace of Spades") {
    t.Errorf("Expected 1st card as Ace of Spades, but got %v", d[0])
  }

  if d[len(d) - 1] != "King of Clubs" {
    t.Errorf("Expected last card as King of Clubs, but got %v", d[51])
  }

} */

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
  os.Remove("helloCards")

  deck := newDeck()
  deck.saveToFile("helloCards")

  loadedDeck := newDeckFromFile("helloCards")

  if len(loadedDeck) != 52 {
    t.Errorf("Expected total 52, but got %v", len(loadedDeck))
  }

  os.Remove("helloCards")
}
