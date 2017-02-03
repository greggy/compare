package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {

  args := os.Args[1:]
  if len(args) != 2 {
    fmt.Println("You should specify two arguments.")
  } else {

    f_path1 := args[0]
    f_path2 := args[1]

    file, err := os.Open(f_path2)
    if err != nil {
      log.Fatalf("readLines in file 2: %s", err)
    }
    defer file.Close()

    set := make(map[string]bool)
    scanner1 := bufio.NewScanner(file)
    for scanner1.Scan() {
      set[scanner1.Text()] = true
    }

    rfile, err := os.Create("result.txt")
    if err != nil {
      log.Fatalf("writeLines: %s", err)
    }
    defer rfile.Close()

    w := bufio.NewWriter(rfile)

    ofile, err := os.Open(f_path1)
    if err != nil {
      log.Fatalf("readLines in file 1: %s", err)
    }
    defer ofile.Close()

    i := 0
    scanner := bufio.NewScanner(ofile)
    for scanner.Scan() {
      line := scanner.Text()		  
      if !set[line] {
	i += 1
	fmt.Fprintln(w, line)
      }
    }
	  
    fmt.Println("We've finished writing", i, "lines into result.txt")
    w.Flush()
  }
}
