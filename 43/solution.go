
/*
 * The number, 1406357289, is a 0 to 9 pandigital number because it is made up
 * of each of the digits 0 to 9 in some order, but it also has a rather
 * interesting sub-string divisibility property.
 *
 * Let d_1 be the 1st digit, d_2 be the 2nd digit, and so on. In this way, we
 * note the following:
 *
 * d_2d_3d_4=406 is divisible by 2
 * d_3d_4d_5=063 is divisible by 3
 * d_4d_5d_6=635 is divisible by 5
 * d_5d_6d_7=357 is divisible by 7
 * d_6d_7d_8=572 is divisible by 11
 * d_7d_8d_9=728 is divisible by 13
 * d_8d_9d_10=289 is divisible by 17
 *
 * Find the sum of all 0 to 9 pandigital numbers with this property.
 */

package main

import "fmt"

func main() {
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
