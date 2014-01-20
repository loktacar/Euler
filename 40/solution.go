/*
 * An irrational decimal fraction is created by concatenating the positive
 * integers:
 *
 * 0.12345678910[1]112131415161718192021...
 *
 * It can be seen that the 12th digit of the fractional part is 1.
 *
 * If d_n represents the nth digit of the fractional part, find the value of
 * the following expression.
 *
 * d_1 × d_10 × d_100 × d_1000 × d_10000 × d_100000 × d_1000000
 */


package main

import "fmt"
import "math"
import "strconv"

/*
n=9: 9 (9)

n=10: 1 (10)
n=189: 9 (99)

n=190: 1 (100)
n=2889: 9 (999)

n=2890: 1 (1000)
n=38889: 9 (9999)
*/
// for n < 10:      n +0      /1
// for n < 190:     n +10     /2
// for n < 2890:    n +110    /3
// for n < 38890:   n +1110   /4
// for n < 488890:  n +11110  /5
// for n < 5888890: n +111110 /6
func get_digit(n int) (int, int) {
  padding, divisor := 0, 1
  if n >= 488890 {
    padding, divisor = 111110, 6
  } else if n >= 38890 {
    padding, divisor = 11110, 5
  } else if n >= 2890 {
    padding, divisor = 1110, 4
  } else if n >= 190 {
    padding, divisor = 110, 3
  } else if n >= 10 {
    padding, divisor = 10, 2
  }

  n_p := n + padding

  i_n := int(math.Floor(float64(n_p) / float64(divisor)))

  i_n_str := fmt.Sprintf("%d", i_n)

  d_n_str := string(i_n_str[n_p % divisor])

  d_n, _ := strconv.ParseInt(d_n_str, 0, 64)

  return int(d_n), i_n
}

func print_get_digit_results(n int) {
  d_n, i_n := get_digit(n)
  fmt.Println(fmt.Sprintf("n=%d: %d (%d)", n, d_n, i_n))
}

func main () {
  //print_get_digit_results(9)

  //print_get_digit_results(10)
  //print_get_digit_results(189)

  //print_get_digit_results(190)
  //print_get_digit_results(2889)

  //print_get_digit_results(2890)
  //print_get_digit_results(38889)

  var indicies = []int{1, 10, 100, 1000, 10000, 100000, 1000000}
  var result int
  result = 1

  for i := 0; i < len(indicies); i++ {
    print_get_digit_results(indicies[i])
    d_n, _ := get_digit(indicies[i])
    result *= d_n
  }

  fmt.Println(result)
}
