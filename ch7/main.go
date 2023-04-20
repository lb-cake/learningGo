package main

import "fmt"

type LogicProvider struct {}

func (lp LogicProvider) Process(data string) string {
  // business logic
  return data
}

type Logic interface {
  Process(data string) string
}

type Client struct {
  L Logic
}

func (c Client) Program() {
  // get data from somewhere
  data := "data"
  retVal := c.L.Process(data)
  fmt.Println(retVal)
}

func main () {
  c := Client{
    L: LogicProvider{},
  }
  c.Program()
}

