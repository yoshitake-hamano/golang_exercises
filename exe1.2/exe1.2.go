package main

import "fmt"
import "os"
import "strings"

func main() {
  fmt.Println(os.Args[0])
  for i, arg := range os.Args[1:] {
    fmt.Printf("%d %s\n", i, arg)
  }
  fmt.Println(strings.Join(os.Args[1:], " "))
}
