/*
 * If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
 *
 * Not all numbers produce palindromes so quickly. For example,
 *
 * 349 + 943 = 1292,
 * 1292 + 2921 = 4213
 * 4213 + 3124 = 7337
 *
 * That is, 349 took three iterations to arrive at a palindrome.
 *
 * Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.
 *
 * How many Lychrel numbers are there below ten-thousand?
*/

package main

import "fmt"
import "strconv"
import "math"

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func reverse_and_add(num string) (string, string) {
  overflow := int64(0)
  new_num := ""
  for i, j := 0, len(num)-1; j >= 0 && i < len(num); i, j = i+1, j-1 {
    first_num, _ := strconv.ParseInt(fmt.Sprintf("%c", num[i]), 0, 0)
    second_num, _ := strconv.ParseInt(fmt.Sprintf("%c", num[j]), 0, 0)
    sum := first_num + second_num + overflow

    overflow = int64(math.Floor(float64(sum) / 10.0))

    new_num = fmt.Sprintf("%d%s", sum % 10, new_num)
  }

  if overflow > 0 {
    new_num = fmt.Sprintf("%d%s", overflow, new_num)
  }

  return new_num, reverse(new_num)
}

func is_lychrel(num string) bool {
  curr_num := num
  curr_rev := reverse(num)
  //fmt.Println(fmt.Sprintf("%s - %s", curr_num, curr_rev))
  for i := 0; i < 50; i++ {
    curr_num, curr_rev = reverse_and_add(curr_num)
    //fmt.Println(fmt.Sprintf("%s - %s", curr_num, curr_rev))

    if curr_num == curr_rev {
      return false
    }
  }
  return true
}

func main() {
  /**/
  count := 0
  for i := 10; i < 10000; i += 1 {
    if is_lychrel(fmt.Sprintf("%d", i)) {
      //fmt.Println(i)
      count += 1
    }
  }
  fmt.Println(count)
  /**/
  /*
  // Non Lychrel numbers
  fmt.Println(is_lychrel("7326"))
  fmt.Println(is_lychrel("47"))
  fmt.Println(is_lychrel("349"))
  fmt.Println(is_lychrel("89"))
  fmt.Println()
  // Lychrel numbers
  fmt.Println(is_lychrel("196"))
  fmt.Println(is_lychrel("295"))
  fmt.Println(is_lychrel("493"))
  fmt.Println(is_lychrel("592"))
  fmt.Println(is_lychrel("689"))
  fmt.Println(is_lychrel("691"))
  fmt.Println(is_lychrel("788"))
  fmt.Println(is_lychrel("790"))
  fmt.Println(is_lychrel("879"))
  fmt.Println(is_lychrel("887"))
  fmt.Println(is_lychrel("4994"))
  /**/
}
