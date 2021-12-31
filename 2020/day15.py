import sys
from typing import List

from common.io import read_lines


def nth_number(nums: List[int], n: int) -> int:
    previous, last = dict(), dict()
    next_num = nums[0]
    turn = 0

    while turn < n:
        if turn < len(nums):
            next_num = nums[turn]
        else:
            next_num = (
                turn - previous[next_num] - 1
                if next_num in previous
                else 0
            )

        if next_num in last:
            previous[next_num] = last[next_num]

        last[next_num] = turn

        turn += 1

    return next_num


def main():
    nums = [
        int(num)
        for num in read_lines(sys.argv[1])[0].split(",")
    ]

    print(f"Part I: {nth_number(nums, 2020)}")
    print(f"Part II: {nth_number(nums, 30000000)}")


if __name__ == "__main__":
    main()
