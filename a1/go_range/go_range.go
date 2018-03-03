package main

import (
  "fmt"
)

func main() {
  go_range(3, 10)
}

func go_range(start int, end int) {
  for i := start; i < end; i++ {
    fmt.Println(i)
  }
}
