package main

import (
  "fmt" 
  "math"
  "os"
  "strconv"
  "log"
)

func Fibonacci(x int) int {
  if  x <= 0  { return 0 }
  if  x == 1  { return 1 }
  return Fibonacci(x-1) + Fibonacci(x-2)
}

func abs(x int) int {
  if x < 0 { return -x }
  return x
}

func negaFibonacci(x int) int {
  var y = int(0)
  var z = 0
  if x < 0 {
    z = abs(x)
    y = int(math.Pow(-1,float64(z+1)))
  }
  return y * Fibonacci(z)
}

func main() {
  var n = 20
  if len(os.Args) >= 2 {
    n1, err := strconv.Atoi(os.Args[1])
    if err != nil {
      log.Printf("convert string to integer err %s" , err)
      n1 = 20
    }
    n = n1
  }
  b := [2]string{"Hello","cy-wang"}
    fmt.Printf("Hello, GO\n")
  for i:=0 ; i<len(b) ;i++ {
    fmt.Printf("%s ",b[i])
  }
  fmt.Printf("\n")

  for i:=0 ; i <= n ; i++ {
    fmt.Printf("fibonacci number: F(%d)=%d\n", i, Fibonacci(i))
  }

  fmt.Printf("-1^10 = %d\n", int(math.Pow(-1,10)) )
  for i:=-n ; i <= -1 ; i++ {
    fmt.Printf("negaFibonacci nF(%d) = %d\n", i, negaFibonacci(i))
  }
}
