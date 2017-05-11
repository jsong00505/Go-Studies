// Created by jsong on 10/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"
import "math"

// const declares a constant value.
const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 500000000

  // Constant expressions perform arithmetic with arbitrary precision.
  const d = 3e20 / n
  fmt.Println(d)

  // A numeric constant has no type until it’s given one,
  // such as by an explicit cast.
  fmt.Println(int64(d))

  // A number can be given a type by using it in a context that requires one,
  // such as a variable assignment or function call.
  fmt.Println(math.Sin(n))
}
