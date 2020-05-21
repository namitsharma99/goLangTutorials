package main

import "fmt"
import "net/http"
import "os"
import "io"

type logWriter struct {}

func main() {
  resp, err := http.Get("http://www.google.com")
  if nil != err {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
  // fmt.Println("response: ",resp)

  // bs := make([]byte, 99999)
  // resp.Body.Read(bs)
  // fmt.Println("byte slice: ", bs)
  // fmt.Println("byte slice as string: ", string(bs))

  // io.Copy(os.Stdout, resp.Body)

  lw := logWriter{}
  io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
  // return 1, nil
  fmt.Println(string(bs))
  fmt.Println("length ---- ", len(bs))
  return len(bs), nil
}
