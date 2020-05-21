package main

import "fmt"
import "net/http"
import "time"

func main() {
  links := []string {
    "http://www.google.com",
    "http://www.golang.org",
    "http://www.stackoverflow.com",
    "http://www.amazon.com",
  }

  c := make(chan string) // need to pass this channel in function call

  for _, link := range links {
    go checkLink(link, c)
  }
  // fmt.Println(<- c)
  // fmt.Println(<- c)

  // for i :=0; i< len(links); i++ {
  //   fmt.Println(<- c)
  // }

  // for {
  //   go checkLink(<- c, c)
  // }

  // Alternative syntax
  for l := range c {
    // time.Sleep(5*time.Second) // 5 seconds, bad way - blocking and queuing
    // go checkLink(l, c)
    go func(link string) {  // function literals, like unnamed and anonymous functions
        time.Sleep(5*time.Second)
        go checkLink(link, c)
      }  (l)
  }

}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if nil != err {
    fmt.Println(link, " is down!!")
    // c <- "down !!"
    c <- link
    return
  }
  fmt.Println(link, " is up!!")
  // c <- "up!!"
  c <- link
}
