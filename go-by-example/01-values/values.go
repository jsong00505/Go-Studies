// Created by jsong on 10/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
  // strings with plus(+)
  fmt.Println("go" + "lang")

  // integers and floats
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("7.0 / 3.0 =", 7.0 / 3.0)

  // booleans with boolean operations
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
}