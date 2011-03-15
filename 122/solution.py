#!/usr/bin/env python
# -*- coding: utf-8 -*-

import math

primes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199]

def main():
    """
    Find the minimum number of multiplications to compute n^k for 1 <= k <= 200.
    """
    sum = 0

    for k in range(1, 201):
    #for k in range(1, 16):
        y = m(k)
        sum += y
        #print k, y
        #print '___'

    print 'sum: ', sum

def m(k):
    if k == 1:
        return 0
    elif k == 2:
        return 1
    elif k in primes:
        return m(k-1) + 1

    sqrt = math.sqrt(k)

    for p in primes:
       if p > sqrt:
           break
       if k % p == 0:
           #print m(p), q(k/p)
           return m(p) + q(k/p)

    return 0

def q(n):
    if n == 1:
        return 0
    elif n == 2:
        return 1

    return q(n/2) + 1 + n%2

if __name__ == '__main__':
    main()
