#!/usr/bin/env python3

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
