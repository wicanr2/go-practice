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

const PORT = 2048

func handleConnection(client net.Conn){
  b := bufio.NewReader(client)
  for {
    line, err := b.ReadBytes('\n')
    if err != nil {
      log.Printf("%s", err)
      client.Close()
      break;
    }
    client.Write(line)
    fmt.Printf("recv: %s\n", line)
  }
}

func main() {
  ln, err := net.Listen("tcp",":" + strconv.Itoa(2048))
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("listen on port :%d\n", PORT)
  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("recv connection local_addr %s, remote_addr %s\n", conn.LocalAddr(), conn.RemoteAddr())
    go handleConnection(conn)
  }
}
