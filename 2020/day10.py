import sys
from collections import Counter
from typing import Iterator, List

from common.io import read_lines_as_ints


COMBINATIONS = {
    2: 2,
    3: 4,
    4: 7,
}


def iter_diffs(ratings: List[int]) -> Iterator[int]:
    return (
        ratings[i] - ratings[j]
        for i, j in zip(
            range(1, len(ratings)),
            range(0, len(ratings) - 1)
        )
    )


def iter_lengths(ratings: List[int]) -> Iterator[int]:
    len_ = 0

    for diff in iter_diffs(ratings):
        if diff == 1:
            len_ += 1
        else:
            if len_ > 1:
                yield len_

            len_ = 0


def main():
    ratings = read_lines_as_ints(sys.argv[1])
    ext_ratings = sorted([0] + ratings + [max(ratings) + 3])
    freq = Counter(iter_diffs(ext_ratings))

    print(f"Part I: {freq[3] * freq[1]}")

    combinations = 1

    for len_ in iter_lengths(ext_ratings):
        combinations *= COMBINATIONS[len_]

    print(f"Part II: {combinations}")


if __name__ == "__main__":
    main()
