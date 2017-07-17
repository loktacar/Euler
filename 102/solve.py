#!/usr/bin/env python3

TRIANGLES_FILE_NAME = 'triangles.txt'

def det(ux, uy, vx, vy):
    return ux * vy - uy * vx

contains_origin_count = 0

with open(TRIANGLES_FILE_NAME, 'r') as f:
    for line in f:
        nums = [float(n) for n in line.split(',')]
        v_0 = [nums[0], nums[1]]
        v_1 = [nums[2] - nums[0], nums[3] - nums[1]]
        v_2 = [nums[4] - nums[0], nums[5] - nums[1]]

        a = -det(v_0[0], v_0[1], v_2[0], v_2[1]) / det(v_1[0], v_1[1], v_2[0], v_2[1])
        b = det(v_0[0], v_0[1], v_1[0], v_1[1]) / det(v_1[0], v_1[1], v_2[0], v_2[1])

        if a > 0 and b > 0 and a + b <= 1:
            contains_origin_count += 1

print("{} triangles contain the origin".format(contains_origin_count))
