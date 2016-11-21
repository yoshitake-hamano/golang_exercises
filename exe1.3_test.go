package main

import (
  "testing"
)

func BenchmarkStringsJoin(b *testing.B) {
  args := []string{"name",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
  }

  for i := 0; i < b.N; i++ {
    stringsJoin(args)
  }
}

func BenchmarkStringPlus(b *testing.B) {
  args := []string{"name",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
    "piyo", "foo", "var", "four", "japan",
  }

  for i := 0; i < b.N; i++ {
    stringPlus(args)
  }
}
