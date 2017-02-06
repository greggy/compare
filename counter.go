package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {

  args := os.Args[1:]
  if len(args) != 1 {
    fmt.Println("You should specify one argument.")
  } else {

    f_path1 := args[0]

    file, err := os.Open(f_path1)
    if err != nil {
      log.Fatalf("readLines in file 1: %s", err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      count += 1
    }
    fmt.Printf("File %s has %d lines.\n", f_path1, count)
  }
}
