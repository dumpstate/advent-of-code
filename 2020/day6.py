import sys
from functools import reduce
from typing import Iterator, List, Tuple

from common.io import read_lines


def iter_groups(lines: List[str]) -> Iterator[List[str]]:
    group = []

    for line in lines:
        if line == "":
            yield group
            group = []
            continue

        group.append(line)

    yield group


def sum_answers(lines: List[str]) -> Tuple[int, int]:
    anyone, everyone = 0, 0

    for group in iter_groups(lines):
        anyone += len(set([*"".join(group)]))
        everyone += len(reduce(
            lambda a, b: a.intersection(b),
            map(set, group),
        ))

    return anyone, everyone


def main():
    lines = read_lines(sys.argv[1])
    anyone, everyone = sum_answers(lines)

    print(f"Part I: {anyone}")
    print(f"Part II: {everyone}")


if __name__ == "__main__":
    main()
