import sys
from dataclasses import dataclass
from itertools import combinations, repeat
from typing import Dict, Iterator, List, Tuple, cast

from common.io import read_lines


@dataclass()
class Op:
    mask: str
    mem: List[Tuple[int, int]]


def parse(lines: List[str]) -> List[Op]:
    ops: List[Op] = []
    op = None

    for line in lines:
        if line.startswith("mask"):
            if op:
                ops.append(op)

            op = Op(mask=line[7:], mem=[])

        if line.startswith("mem"):
            l, r = line[4:].split("] = ")
            cast(Op, op).mem.append((int(l), int(r)))
    
    ops.append(cast(Op, op))

    return ops


def pad_right(arr: List[int], len_: int, value: int) -> List[int]:
    if len(arr) > len_:
        raise ValueError("Too long, cannot pad")

    return arr + list(repeat(value, len_ - len(arr)))


def binary_arr(num: int) -> List[int]:
    n = num
    binary = []

    while n > 0:
        binary.append(n % 2)
        n = n // 2

    return binary


def join(arr: List[int]) -> str:
    return "".join([str(value) for value in arr])


def with_mask(bnry: str, mask: str) -> str:
    masked = []

    for value, mask_value in zip(bnry, mask):
        if mask_value == "X":
            masked.append(value)
        else:
            masked.append(mask_value)

    return join(masked)


def with_floating_mask(bnry: str, mask: str) -> str:
    masked = []

    for value, mask_value in zip(bnry, mask):
        if mask_value == "0":
            masked.append(value)
        else:
            masked.append(mask_value)

    return join(masked)


def binary(num: int, mask: str) -> str:
    padded = pad_right(binary_arr(num), 36, 0)
    padded.reverse()

    return with_mask(join(padded), mask)


def binary_with_floats(num: int, mask: str) -> str:
    padded = pad_right(binary_arr(num), 36, 0)
    padded.reverse()

    return with_floating_mask(join(padded), mask)


def binary_to_decimal(num: str) -> int:
    res = 0

    for i in range(len(num) - 1, -1, -1):
        res += int(num[i]) * pow(2, len(num) - i - 1)

    return res


def decimal_with_mask(num: int, mask: str) -> int:
    return binary_to_decimal(binary(num, mask))


def count_floats(masked: str) -> int:
    floats = 0

    for bit in masked:
        if bit == "X":
            floats += 1

    return floats


def iter_combinations(len_: int) -> Iterator[Tuple[int, ...]]:
    yield from set(combinations([0, 1] * len_, len_))


def apply_combination(masked: str, combination: Tuple[int, ...]) -> int:
    res = ""
    ix = 0

    for bit in masked:
        if bit == "X":
            res += str(combination[ix])
            ix += 1
        else:
            res += bit

    return binary_to_decimal(res)


def iter_masked(masked: str) -> Iterator[int]:
    floats = count_floats(masked)

    for combination in iter_combinations(floats):
        yield apply_combination(masked, combination)


def total_register(register: Dict[int, int]) -> int:
    total = 0

    for value in register.values():
        total += value

    return total


def exec_a(ops: List[Op]) -> int:
    register = dict()

    for op in ops:
        for mem_op in op.mem:
            ix, value = mem_op
            register[ix] = decimal_with_mask(value, op.mask)

    return total_register(register)


def exec_b(ops: List[Op]) -> int:
    register = dict()

    for op in ops:
        for mem_op in op.mem:
            ix, value = mem_op
            masked_index = binary_with_floats(ix, op.mask)

            for i in iter_masked(masked_index):
                register[i] = value

    return total_register(register)


def main():
    ops = parse(read_lines(sys.argv[1]))

    print(f"Part I: {exec_a(ops)}")
    print(f"Part II: {exec_b(ops)}")


if __name__ == "__main__":
    main()
