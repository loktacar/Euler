#!/usr/bin/env python3

from itertools import product
import string

"""
Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes. The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key. If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message. The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower case characters. Using cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
"""

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
