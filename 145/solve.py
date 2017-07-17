#!/usr/bin/env python3

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
