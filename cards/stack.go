package main

import "fmt"
// creating a new type of deck - slice of string

type stack[] string

func (s stack) print() {
  for i, str := range s {
    fmt.Println(i, str)
  }
}
