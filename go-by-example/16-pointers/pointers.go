// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  // basic example of pointers
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}
