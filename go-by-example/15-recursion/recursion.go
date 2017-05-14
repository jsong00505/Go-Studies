// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
  // basic recursion example
  fmt.Println(fact(7))
}
