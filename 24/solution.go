/*
 * A permutation is an ordered arrangement of objects. For example, 3124 is one
 * possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
 * are listed numerically or alphabetically, we call it lexicographic order.
 * The lexicographic permutations of 0, 1 and 2 are:

 * 012   021   102   120   201   210

 * What is the millionth lexicographic permutation of the digits
 * 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
 */

package main

import "fmt"

func main () {
  //list := []int{1, 2, 3, 4}
  list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

  c := 1
  for {
    //fmt.Println(fmt.Sprintf("%d: %v", c, list))

    //////
    // Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
    k := -1
    i := 0
    for i+1 < len(list) {
      if list[i] < list[i+1] {
        k = i
      }
      i += 1
    }
    if k == -1 {
      break // last permutation
    }

    //////
    // Find the largest index l such that a[k] < a[l].
    l := 0
    i = 0
    for i < len(list) {
      if list[k] < list[i] {
        l = i
      }
      i += 1
    }

    //////
    // Swap the value of a[k] with that of a[l].
    old_k := list[k]
    list[k] = list[l]
    list[l] = old_k

    //////
    // Reverse the sequence from a[k + 1] up to and including the final element a[n].

    // Copy slices
    to_reverse := append([]int(nil), list[k+1:]...)
    list = append([]int(nil), list[:k+1]...)

    // Reverse
    i = len(to_reverse) - 1
    for i >= 0 {
      list = append(list, to_reverse[i])
      i -= 1
    }

    c += 1

    // Stop on one MILLION permutations
    if c == 1000000 {
      break
    }
  }

  fmt.Println(list)
}
