// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {

  // basic example
  s := make([]string, 3)
  fmt.Println("emp", s)

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"

  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  fmt.Println("len:", len(s))

  // append elements in the list
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  // copy the list to another list
  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  // slicing examples
  l := s[2:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  // declare and initialize a list in a single
  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  // multiple arrays
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}