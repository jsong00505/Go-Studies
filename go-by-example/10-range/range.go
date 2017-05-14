// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {

  // basic example
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  // iterate a map by keys or values
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for k := range kvs {
    fmt.Println("key:", k)
  }

  // iterate over Unicode code points
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
