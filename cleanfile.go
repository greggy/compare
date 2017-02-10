package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
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

    rfile, err := os.Create("c" + f_path1)
    if err != nil {
      log.Fatalf("writeLines: %s", err)
    }
    defer rfile.Close()

    w := bufio.NewWriter(rfile)

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
      line := scanner.Text()
      cdata := strings.TrimSpace(line)
      cdata2 := strings.Trim(cdata, "\n")	    
      if cdata2 != "" && cdata2 != "\n" && cdata2 != "\r\n" && cdata2 != "\t" {
        fmt.Fprintln(w, cdata2)
      }
    }

    fmt.Printf("Cleaned data was written to %s file.\n", "c" + f_path1)
    w.Flush()
  }
}
