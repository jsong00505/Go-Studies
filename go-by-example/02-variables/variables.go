// Created by jsong on 10/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
  // string
  var a string = "initial"
  fmt.Println(a)

  // integers
  var b, c int = 1, 2
  fmt.Println(b, c)

  // boolean
  var d = true
  fmt.Println(d)

  // Variables declared without a corresponding initialization are zero-valued
  var e int
  fmt.Println(e)

  // The := syntax is shorthand for declaring and initializing a variable.
  // e.g. for var f string = "short" in this case
  f := "short"
  fmt.Println(f)
}