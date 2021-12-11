import sys
from typing import Iterator, List, Set

from common.io import read_lines


class LoopFoundException(Exception):

    acc: int

    def __init__(self, acc: int):
        super().__init__(f"Loop found: {acc}")

        self.acc = acc


class Executor:

    instructions: List[str]
    acc: int
    pointer: int
    loop_guard: Set[int]

    def __init__(self, instructions: List[str]):
        self.instructions = instructions
        self.acc = 0
        self.pointer = 0
        self.loop_guard = set()

    def exec(self):
        while True:
            if self.pointer in self.loop_guard:
                raise LoopFoundException(self.acc)

            if self.pointer == len(self.instructions):
                return self.acc

            self.loop_guard.add(self.pointer)

            op, value = self.instructions[self.pointer].split()

            if op == "nop":
                self.pointer += 1
            elif op == "acc":
                self.acc += int(value)
                self.pointer += 1
            elif op == "jmp":
                self.pointer += int(value)
            else:
                raise ValueError(f"Unknown instruction: {op}")


def iter_instructions(instructions: List[str]) -> Iterator[List[str]]:
    yield instructions

    for i in range(0, len(instructions)):
        op, value = instructions[i].split()

        if op == "nop":
            next_instructions = instructions.copy()
            next_instructions[i] = f"jmp {value}"

            yield next_instructions
        elif op == "jmp":
            next_instructions = instructions.copy()
            next_instructions[i] = f"nop {value}"

            yield next_instructions


def main():
    instructions = read_lines(sys.argv[1])
    executor = Executor(instructions)

    try:
        executor.exec()
    except LoopFoundException as ex:
        print(f"Part I: {ex.acc}")

    for fixed_instructions in iter_instructions(instructions):
        ex = Executor(fixed_instructions)

        try:
            acc = ex.exec()

            print(f"Part II: {acc}")
            break
        except LoopFoundException:
            pass


if __name__ == "__main__":
    main()
