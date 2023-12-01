package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  dat, err := os.ReadFile("/data.txt")
  check(err)
  fmt.Print
}
