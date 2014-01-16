/*
 * If the numbers 1 to 5 are written out in words: one, two, three, four, five,
 * then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written out
 * in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
 * forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
 * letters. The use of "and" when writing out numbers is in compliance with
 * British usage.
 */

package main

import "fmt"
import "math"

var numbers = map[int]string{
  1: "one",
  2: "two",
  3: "three",
  4: "four",
  5: "five",
  6: "six",
  7: "seven",
  8: "eight",
  9: "nine",
  10: "ten",
  20: "twenty",
  30: "thirty",
  40: "forty",
  50: "fifty",
  60: "sixty",
  70: "seventy",
  80: "eighty",
  90: "ninety",
  100: "onehundred",
  200: "twohundred",
  300: "threehundred",
  400: "fourhundred",
  500: "fivehundred",
  600: "sixhundred",
  700: "sevenhundred",
  800: "eighthundred",
  900: "ninehundred",
  1000: "onethousand",
}

var exceptions = map[int]string{
  11: "eleven",
  12: "twelve",
  13: "thirteen",
  14: "fourteen",
  15: "fifteen",
  16: "sixteen",
  17: "seventeen",
  18: "eighteen",
  19: "nineteen",
}

func num_to_string(number int) string {
  exception_text, ok := exceptions[number]
  // is an exception
  if ok {
    return exception_text
  }

  // one thousand
  if number == 1000 {
    return numbers[number]
  }

  var hundreds_text string

  // hundreds
  hundreds := math.Floor(float64(number) / 100.0)
  if hundreds > 0 {
    hundreds_text = numbers[int(hundreds * 100)]
  }

  number = int(float64(number) - hundreds * 100.0)

  var tens_text string

  // tens
  exception_text, ok = exceptions[number]
  if ok && hundreds_text != "" {
    return hundreds_text + "and" + exception_text
  } else if ok {
    return exception_text
  }

  tens := math.Floor(float64(number) / 10.0)
  if tens > 0 {
    tens_text = numbers[int(tens * 10)]
  }

  number = int(float64(number) - tens * 10)

  singles_text := numbers[int(number)]

  if hundreds_text != "" && (tens_text != "" || singles_text != "") {
    return hundreds_text + "and" + tens_text + singles_text
  } else if hundreds_text != "" {
    return hundreds_text
  }

  return tens_text + singles_text
}

func main() {
  /*
  thing := num_to_string(18)
  fmt.Println(thing)
  fmt.Println(len(thing))
  */

  var total_text string
  num := 1
  for num <= 1000 {
    total_text += num_to_string(num)
    //fmt.Println(num_to_string(num))
    num += 1
  }

  //fmt.Println(total_text)
  fmt.Println(len(total_text))
}
