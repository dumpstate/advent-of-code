import sys
from typing import Iterator, List

from common.io import read_lines_as_ints


PREAMBLE_SIZE = 25


def find_invalid_number(numbers: List[int]) -> int:
    for i in range(0, len(numbers) - PREAMBLE_SIZE):
        preamble = numbers[i:i + PREAMBLE_SIZE]
        sums = set(
            preamble[x] + preamble[y]
            for x in range(0, PREAMBLE_SIZE)
            for y in range(x + 1, PREAMBLE_SIZE)
        )
        next_num = numbers[i + PREAMBLE_SIZE]

        if next_num not in sums:
            return next_num

    raise Exception("Invalid number not found")


def find_subarr_of_total(numbers: List[int], needle: int) -> List[int]:
    s_ix, e_ix, total = 0, 0, numbers[0]

    while s_ix < len(numbers) and e_ix < len(numbers):
        if total == needle:
            return numbers[s_ix:e_ix]
        elif total > needle:
            total -= numbers[s_ix]
            s_ix += 1
        else:
            e_ix += 1
            total += numbers[e_ix]

    raise Exception("Subarr not found")


def main():
    numbers = read_lines_as_ints(sys.argv[1])
    invalid_number = find_invalid_number(numbers)

    print(f"Part I: {invalid_number}")

    subarr = find_subarr_of_total(numbers, invalid_number)

    print(f"Part II: {min(subarr) + max(subarr)}")


if __name__ == "__main__":
    main()
