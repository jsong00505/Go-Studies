// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// basic example
type person struct {
  name string
  age int
}

func main() {
  fmt.Println(person{"Bob", 20})

  fmt.Println(person{name: "Alice", age: 30})

  fmt.Println(person{name: "Fred"})

  fmt.Println(&person{name: "Ann", age: 40})

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  var sp *person
  sp = &s
  // sp := &s

  fmt.Println("sp:", sp)
  fmt.Println("&sp:", &sp)
  fmt.Println("*sp:", *sp)
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)
}