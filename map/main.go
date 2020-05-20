package main

import "fmt"

func main() {
  // colors := map[string] string {
  //   "red" : "#ff0000",
  //   "green" : "#433243",
  // }

  // var colors map[string] string
  colors := make(map[string]string)

  colors["white"] = "#ffffff"
  colors["grey"] = "#222222"
  fmt.Println(colors)

  // delete(colors, "white")
  // fmt.Println(colors)

  printMap(colors)
}

func printMap (c map[string]string ) {
  for color, hex := range c {
    fmt.Println("Hex code for ", color, "is ", hex)
  }
}
