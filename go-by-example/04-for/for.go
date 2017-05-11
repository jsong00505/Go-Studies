package main

import "fmt"

func main() {

  // can use the for loop such as the while loop of other languages
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // the common for loop
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // for without condition will be continue
  // till meeting the break or return statement.
  for {
    fmt.Println("loop")
    break
  }

  // continue
  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
