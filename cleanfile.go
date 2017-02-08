package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"	
  "io/ioutil"
  "regexp"
)

func main() {

  args := os.Args[1:]
  if len(args) != 1 {
    fmt.Println("You should specify one argument.")
  } else {

    f_path1 := args[0]

    file, err := ioutil.ReadFile(f_path1)
    if err != nil {
      log.Fatalf("readLines in file 1: %s", err)
    }

    data := string(file)
    cdata := strings.Replace(data, "\r", "\n", -1)
    cdata2 := strings.Replace(cdata, "\n\n", "\n", -1)
    re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
    cdata3 := re.ReplaceAllString(cdata2, "")
    cdata4 := strings.Replace(cdata3, "\t", "", -1)

    rfile, err := os.Create("c" + f_path1)
    w := bufio.NewWriter(rfile)
    _, err = w.WriteString(cdata4)

    fmt.Printf("Cleaned data was written to %s file.\n", "c" + f_path1)
    w.Flush()
  }
}
