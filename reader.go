package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  var path = "test.txt"
  if len(os.Args) >= 2 {
    path = os.Args[1]
  }

  fmt.Printf("open file path %s\n", path)
  file, err := os.Open(path)
  if err != nil {
    fmt.Printf("\n#%s\n", err );
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if  err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

