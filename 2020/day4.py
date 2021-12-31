import re
import sys
from functools import partial
from typing import Callable, Dict, List

from common.io import read_lines


Passport = Dict[str, str]


def read_passports(file_path: str) -> List[Passport]:
    passports: List[Passport] = []
    next_passport = dict()

    for line in read_lines(file_path):
        if line == "":
            passports.append(next_passport)
            next_passport = dict()
            continue

        for entry in line.split():
            key, value = entry.split(":")
            next_passport[key] = value

    passports.append(next_passport)

    return passports


REQUIRED_FIELDS = frozenset([
    "byr", "iyr", "eyr",
    "hgt", "hcl", "ecl",
    "pid",
])


def has_required_fields(passport: Passport) -> bool:
    return not REQUIRED_FIELDS.difference(passport.keys())


def is_number(value: str) -> bool:
    try:
        int(value)
        return True
    except ValueError:
        return False


def min_max(min_, max_, value):
    num = int(value)

    return num >= min_ and num <= max_


def is_height(value: str) -> bool:
    if value[-2:] == "cm":
        num = int(value[:-2])

        return num >= 150 and num <= 193

    if value[-2:] == "in":
        num = int(value[:-2])

        return num >= 59 and num <= 76

    return False


VALIDATIONS = {
    "byr": [
        is_number,
        partial(min_max, 1920, 2002),
    ],
    "iyr": [
        is_number,
        partial(min_max, 2010, 2020),
    ],
    "eyr": [
        is_number,
        partial(min_max, 2020, 2030),
    ],
    "ecl": [
        frozenset([
            "amb", "blu", "brn", "gry",
            "grn", "hzl", "oth",
        ]).__contains__,
    ],
    "pid": [
        is_number,
        lambda value: len(value) == 9,
    ],
    "hgt": [
        is_height,
    ],
    "hcl": [
        lambda value: bool(re.match("^#([0-9a-f]{6})$", value))
    ],
}


def is_valid(passport: Passport) -> bool:
    return all(
        key in passport and all(
            fn(passport[key])
            for fn in fns
        )
        for key, fns in VALIDATIONS.items()
    )


def count(passports: List[Passport], validator: Callable[[Passport], bool]) -> int:
    return len(list(filter(validator, passports)))


def main():
    passports = read_passports(sys.argv[1])

    print(f"Part I: {count(passports, has_required_fields)}")
    print(f"Part II: {count(passports, is_valid)}")


if __name__ == "__main__":
    main()
