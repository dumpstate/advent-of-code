import sys
from collections import Counter
from typing import Callable, List, Tuple

from common.io import read_lines


PasswordEntry = Tuple[Tuple[str, int, int], str]


def password_with_policy(line: str) -> PasswordEntry:
    policy_str, password = line.split(": ")
    range_str, char = policy_str.split(" ")
    policy_min_str, policy_max_str = range_str.split("-")

    return ((char, int(policy_min_str), int(policy_max_str)), password)


def is_valid(entry: PasswordEntry) -> bool:
    (char, min, max), password = entry
    count = Counter(password)[char]

    return count >= min and count <= max


def is_valid_toboggan(entry: PasswordEntry) -> bool:
    (char, ix_1, ix_2), password = entry

    return (password[ix_1 - 1] == char) ^ (password[ix_2 - 1] == char)


def count_valid_passwords(
    passwords: List[PasswordEntry],
    validator: Callable[[PasswordEntry], bool]
) -> int:
    return len(list(filter(validator, passwords)))


def main():
    passwords = [
        password_with_policy(line)
        for line in read_lines(sys.argv[1])
    ]

    print(f"Part I: {count_valid_passwords(passwords, is_valid)}")
    print(f"Part II: {count_valid_passwords(passwords, is_valid_toboggan)}")


if __name__ == "__main__":
    main()
