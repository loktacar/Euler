#!/usr/bin/env python3

from itertools import product
import string


# Settings
CIPHER_FILE_NAME = 'cipher.txt'
SEARCH_STRING = 'the'

# Helper function
def xor_list_with_list(a, b):
    out = []
    for i, x in enumerate(a):
        out.append(x ^ b[i % len(b)])
    return out

# Read and parse file
cipher_numbers = []
with open(CIPHER_FILE_NAME, 'r') as f:
    cipher_numbers = [int(num) for num in f.read().split(',')]

# Try all combinations of keys (given that it is three lower-case letters) and
# find the output string with the most occurances of the search string
max_occur = 0
max_occur_value = None
max_occur_key = None
for key in product(string.ascii_lowercase, repeat=3):
    key_ints = [ord(c) for c in key]
    list = [chr(i) for i in xor_list_with_list(cipher_numbers, key_ints)]
    string = ''.join(list)

    if string.count(SEARCH_STRING) > max_occur:
        max_occur = string.count(SEARCH_STRING)
        max_occur_key = ''.join(key)
        max_occur_value = string

# Print the key and the actual value
#print('{}: {}'.format(max_occur_key, max_occur_value))

# Print the sum of the ascii values of the original string
print('Answer: {}'.format(sum([ord(c) for c in max_occur_value])))
