#!/usr/bin/env python3

from itertools import product
from decimal import *

"""
Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
where each “_” is a single digit.
"""

getcontext().prec = 19

exponents = [
    10**1, 10**3, 10**5, 10**7, 10**9, 10**11, 10**13, 10**15, 10**17, 10**19
]

for l in product([0, 1, 2, 3, 4, 5, 6, 7, 8, 9], repeat=9):
    form = Decimal(1020304050607080900)

    for i in range(9):
        form += exponents[i] * l[i]

    if form.sqrt() % 1 == 0:
        print("sqrt({}) = {}".format(form, form.sqrt()))
        break
