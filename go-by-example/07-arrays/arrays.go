// Created by jsong on 13/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {

  var a [5]int
  fmt.Println("emp:", a)

  // set a value in array
  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

  fmt.Println("len:", len(a))

  // declare and init values in array
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl", b)

  // 2D array
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("twoD:", twoD)
}