package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func elemInArray(a string, list []string) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}

func main() {
  lines := []string{}
  args := os.Args[1:]
  if len(args) != 2 {
    fmt.Println("You should specify two arguments.")
  } else {

    f_path1 := args[0]
    f_path2 := args[1]

    lines1, err := readLines(f_path1)
    if err != nil {
      log.Fatalf("readLines in file 1: %s", err)
    }

    lines2, err := readLines(f_path2)
    if err != nil {
      log.Fatalf("readLines in file 2: %s", err)
    }

    for _, line := range lines1 {
      if !elemInArray(line, lines2) {
	lines = append(lines, line)
      }	    
    }

    fmt.Println("We're going to write", len(lines), "lines into result.txt")
    if err := writeLines(lines, "result.txt"); err != nil {
      log.Fatalf("writeLines: %s", err)
    }
  }
}
