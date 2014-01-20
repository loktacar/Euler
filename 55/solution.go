/*

*/

package main

import "fmt"
import "strconv"

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func reverse_and_add(num int) (int, int) {
  new_num := num + reverse_num(num)
  return new_num, reverse_num(new_num)
}

func reverse_num(num int) int {
  num_str, _ := strconv.ParseInt(reverse(fmt.Sprintf("%d", num)), 0, 64)
  return int(num_str)
}

func is_lychrel(num int) bool {
  curr_num := num
  curr_rev := reverse_num(num)
  for i := 0; i < 50; i++ {
    curr_num, curr_rev = reverse_and_add(num)

    if curr_num == curr_rev {
      return true
    }
  }
  return false
}

func main() {
  fmt.Println(is_lychrel(47))
  fmt.Println(is_lychrel(349))
}
