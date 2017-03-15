package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
  "log"
  "bufio"
  "net"
)

var _ = os.Args
var _ = math.Abs
var _ = strconv.Itoa
var _ = bufio.NewReader
var _ = fmt.Printf

const PORT = 2048

func main() {
  conn, err := net.Dial("tcp",":2048")
  if err != nil {
    log.Fatal(err)
  }
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  fmt.Println(text)

  conn.Write([]byte(text))
  fmt.Printf("send data %s to server", text)
  defer conn.Close()
  fmt.Printf("close connection")
}
