package main

import (
  "fmt"
  "time"
)

var message = make(chan string)

func ping() {
  for {
    msg := <-message
    if msg == "pong" {
      message <- "ping"
      fmt.Printf("PING --> %s\n", msg)
      time.Sleep(time.Duration(200) * time.Millisecond)
      fmt.Printf("Send PING\n")
    }
  }
}

func pong() {
  for {
    msg := <-message
    if msg == "ping" {
      fmt.Printf("PONG <-- %s\n", msg)
      time.Sleep(time.Duration(200) * time.Millisecond)
      fmt.Printf("Send PONG back\n")
      message <- "pong"
    }
  }
}

func main() {
    fmt.Printf("goroutine test\n")
    go ping()
    go pong()
    message <- "ping"
    for {
      time.Sleep(time.Duration(2) * time.Second)
    }
}
