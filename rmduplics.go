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

    set := make(map[string]bool)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      line := scanner.Text()
      if !set[line] {
	set[line] = true
      }
    }

    rfile, err := os.Create("rmduplics.txt")
    if err != nil {
      log.Fatalf("writeLines: %s", err)
    }
    defer rfile.Close()

    w := bufio.NewWriter(rfile)

    i := 0
    for k := range set {
      i += 1
      fmt.Fprintln(w, k)
    }

    fmt.Println("We've finished writing", i, "lines into rmduplics.txt")
    w.Flush()
  }
}
