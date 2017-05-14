// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// take an arbitrary number of ints as arguments
func sum(nums ... int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  // basic examples
  sum(1, 2)
  sum(1, 2, 3)

  // using (slice...)
  nums := []int{1, 2, 3, 4}
  sum(nums ...)
}
