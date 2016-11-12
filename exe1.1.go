package main

import "fmt"
import "os"
import "strings"

func main() {
  fmt.Println(os.Args[0])
  fmt.Println(strings.Join(os.Args[1:], " "))
}
