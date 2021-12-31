from typing import List


def read_lines(file_path: str) -> List[str]:
    with open(file_path, "r") as f:
        return [line.strip() for line in f.readlines()]


def read_lines_as_ints(file_path: str) -> List[int]:
    return [
        int(line)
        for line in read_lines(file_path)
    ]
