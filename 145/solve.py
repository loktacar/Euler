#!/usr/bin/env python3

"""
Some positive integers n have the property that the sum [ n + reverse(n) ] consists entirely of odd (decimal) digits. For instance, 36 + 63 = 99 and 409 + 904 = 1313. We will call such numbers reversible; so 36, 63, 409, and 904 are reversible. Leading zeroes are not allowed in either n or reverse(n).

There are 120 reversible numbers below one-thousand.

How many reversible numbers are there below one-billion (10^9)?
"""

def reverse(num):
    rev = 0

    c = 0
    while num:
        digit = num % 10
        num //= 10
        rev = (rev*10) + digit
        c += 1

    return rev

def is_all_odd(num):
    while num:
        digit = num % 10
        num //= 10

        if digit % 2 == 0:
            return False
    return True


def how_many():
    reversible = 0
    for i in range(1, pow(10, 6)):
        if i % 10 == 0:
            continue

        i_rev = reverse(i)
        if is_all_odd(i + i_rev):
            reversible += 1

    return reversible

print("number of reversible: {}".format(how_many()))
