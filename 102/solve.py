#!/usr/bin/env python3

"""
Three distinct points are plotted at random on a Cartesian plane, for which -1000 ≤ x, y ≤ 1000, such that a triangle is formed.

Consider the following two triangles:

    A(-340,495), B(-153,-910), C(835,-947)

    X(-175,41), Y(-421,-714), Z(574,-645)

It can be verified that triangle ABC contains the origin, whereas triangle XYZ does not.

Using triangles.txt (right click and 'Save Link/Target As...'), a 27K text file containing the co-ordinates of one thousand "random" triangles, find the number of triangles for which the interior contains the origin.

NOTE: The first two examples in the file represent the triangles in the example given above.
"""

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
