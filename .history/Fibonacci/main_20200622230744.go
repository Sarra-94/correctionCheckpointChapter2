package main

import "fmt"

func Fibonacci(n uint) uint {
  if n <= 1 {
     return n
  }
  return Fibonacci(n-1) + Fibonacci(n-2)
}

func main()  {
  fmt.Println("f(25) =",Fibonacci(25))
}
