package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "io"	
)

func main() {

  args := os.Args[1:]
  if len(args) != 2 {
    fmt.Println("You should specify two arguments.")
  } else {

    f_path1 := args[0]
    f_path2 := args[1]

    file, err := os.Open(f_path1)
    if err != nil {
      log.Fatalf("readLines in file 1: %s", err)
    }
    defer file.Close()

    rfile, err := os.Create("cresult.txt")
    if err != nil {
      log.Fatalf("writeLines: %s", err)
    }
    defer rfile.Close()

    _, err = io.Copy(rfile, file)
    if err != nil {
      log.Fatalf("copyFile: %s", err)	
    }

    err = rfile.Sync()
    if err != nil {
        log.Fatal(err)
    }

    w := bufio.NewWriter(rfile)

    ofile, err := os.Open(f_path2)
    if err != nil {
      log.Fatalf("readLines in file 2: %s", err)
    }
    defer ofile.Close()

    scanner := bufio.NewScanner(ofile)
    for scanner.Scan() {
      fmt.Fprintln(w, scanner.Text())
    }
	  
    fmt.Printf("We've finished concatinate two files %s and %s into cresult.txt\n", f_path1, f_path2)
    w.Flush()
  }
}
