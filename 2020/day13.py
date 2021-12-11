import sys
from functools import reduce
from typing import List, Tuple, Union

from common.io import read_lines


def parse(path: str) -> Tuple[int, List[Union[int, str]]]:
    lines = read_lines(path)

    return (
        int(lines[0]),
        [
            int(bus) if bus != "x" else bus
            for bus in lines[1].split(",")
        ],
    )


def find_a(next_ts: int, buses: List[Union[int, str]]) -> int:
    rem: List[Tuple[int, int]] = [
        (int(bus), int(bus) - next_ts % int(bus))
        for bus in buses
        if type(bus) == int
    ]

    return reduce(lambda x, y: x * y, min(rem, key=lambda pair: pair[1]))


def inverse(num: int, mod: int) -> int:
    i = 1

    while True:
        if int((i * num) % mod) == 1:
            return i

        i += 1


def chinese(nums: List[Tuple[int, int]]) -> int:
    N = reduce(lambda x, y: x * y, (num for num, _ in nums))
    total = 0

    for num, offset in nums:
        n = int(N / num)
        total += n * inverse(n, num) * offset

    return int(total % N)


def collect_offsets(buses: List[Union[int, str]]) -> List[Tuple[int, int]]:
    offsets = []
    offset = -1

    for bus in buses:
        offset += 1

        if type(bus) == str:
            continue

        offsets.append((int(bus), -offset))

    return offsets


def find_b(buses: List[Union[int, str]]) -> int:
    return chinese(collect_offsets(buses))


def main():
    next_ts, buses = parse(sys.argv[1])

    print(f"Part I: {find_a(next_ts, buses)}")
    print(f"Part II: {find_b(buses)}")


if __name__ == "__main__":
    main()
