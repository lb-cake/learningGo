package main

import "fmt"

func main() {
  ch := make(chan bool)
  go func() {
    fmt.Println("hello from goroutines")
    ch <- true
  }()
  fmt.Println("end of main")
  <- ch
}
