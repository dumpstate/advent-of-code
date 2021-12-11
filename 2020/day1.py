import sys
from typing import List

from common.io import read_lines_as_ints


def matching_pair_product(values: List[int], value: int) -> int:
    for i in range(0, len(values)):
        for j in range(i + 1, len(values)):
            if values[i] + values[j] == value:
                return values[i] * values[j]

    raise Exception(f"No match for {value}")


def matching_triple_product(values: List[int], value: int) -> int:
    for i in range(0, len(values)):
        for j in range(i + 1, len(values)):
            for k in range(j + 1, len(values)):
                if values[i] + values[j] + values[k] == value:
                    return values[i] * values[j] * values[k]

    raise Exception(f"No match for {value}")


def main():
    expenses = read_lines_as_ints(sys.argv[1])

    print(f"Part I: {matching_pair_product(expenses, 2020)}")
    print(f"Part II: {matching_triple_product(expenses, 2020)}")


if __name__ == "__main__":
    main()
