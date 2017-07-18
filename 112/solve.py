#!/usr/bin/env python3

"""
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy. In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.

Find the least number for which the proportion of bouncy numbers is exactly 99%.
"""

def is_bouncy(num):
    last_digit = num % 10

    is_decreasing = True
    is_increasing = True

    while num:
        #print(num)
        digit = num % 10
        #print("{} - {}".format(digit, last_digit))

        # not decreasing
        if digit < last_digit:
            is_decreasing = False

        # not increasing
        if digit > last_digit:
            is_increasing = False

        # it is not decreasing nor decreasing
        if not is_decreasing and not is_increasing:
            return True

        num //= 10
        last_digit = digit

    #if is_decreasing:
    #    print("decreasing")
    #if is_increasing:
    #    print("increasing")

    return False

#numbers = [
#    134468,
#    66420,
#    155349
#]
#
#for num in numbers:
#    print("{} is bouncy: {}".format(num, is_bouncy(num)))

bouncy_count = 0

num = 1

while bouncy_count / num < 0.99:
    num += 1
    if is_bouncy(num):
        bouncy_count += 1

print("bouncy_count: {}".format(bouncy_count))
print("num: {}".format(num))
print("ratio: {}".format(bouncy_count / num))
