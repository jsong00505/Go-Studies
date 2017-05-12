// Created by jsong on 12/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"
import "time"

func main() {

  // basic example
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // use commas for separate multiple expression
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's a weekday")
  }

  // use if/else logic
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  // use type instead of values
  whatAmI := func(i interface{}) {
    switch t:= i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Printf("Don't know type %T\n", t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
