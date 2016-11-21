package main

import (
  "os"
  "fmt"
  "strings"
)

func stringsJoin(args []string) {
  strings.Join(args[1:], " ")
}

func stringPlus(args []string) {
  var s, sep string
  for i := 1; i < len(args); i++ {
    s += sep + args[i]
    sep = " "
  }
}

func main() {
  stringsJoin(os.Args)
  stringPlus(os.Args)
  fmt.Println("foo")
}
