// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func intSeq() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func main() {
  // basic example
  nextInt := intSeq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // confirm the function was unique by making the new example
  newInts := intSeq()
  fmt.Println(newInts())
}
