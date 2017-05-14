// Created by jsong on 14/05/2017. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {

  // basic example
  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1:", v1)

  fmt.Println("len:", len(m))

  // delete an element in the map
  delete(m, "k2")
  fmt.Println("map:", m)

  // if a key is present in the map, it returns true
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  // declare and initialize a map in a single
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map", n)
}
