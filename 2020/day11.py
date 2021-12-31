import sys
from copy import deepcopy
from itertools import repeat
from typing import Callable, Iterator, List, Tuple

from common.io import read_lines


Board = List[List[str]]


def dim(board: Board) -> Tuple[int, int]:
    return len(board[0]), len(board)


def map_board(fn: Callable[[Tuple[int, int], Board], str]) -> Callable[[Board], Tuple[Board, int]]:
    def map_(board: Board) -> Tuple[Board, int]:
        result_board = deepcopy(board)
        x_size, y_size = dim(board)
        moves = 0

        for y in range(0, y_size):
            for x in range(0, x_size):
                value = fn((x, y), board)

                if value != result_board[y][x]:
                    moves += 1

                result_board[y][x] = value

        return result_board, moves

    return map_


def iter_neighbours(seat: Tuple[int, int], board: Board) -> Iterator[str]:
    x, y = seat
    x_size, y_size = dim(board)

    for i in range(max(x - 1, 0), min(x + 2, x_size)):
        for j in range(max(y - 1, 0), min(y + 2, y_size)):
            if i == x and j == y:
                continue

            yield board[j][i]


def count_neighbours(seat: Tuple[int, int], board: Board) -> int:
    return sum(
        1
        for neighbour in iter_neighbours(seat, board)
        if neighbour == "#"
    )


def iter_in_direction(seat: Tuple[int, int], direction: str, board: Board) -> Iterator[Tuple[int, int]]:
    x, y = seat
    x_size, y_size = dim(board)

    if "U" in direction:
        y_iter = range(y - 1, -1, -1)
    elif "D" in direction:
        y_iter = range(y + 1, y_size)
    else:
        y_iter = repeat(y)

    if "L" in direction:
        x_iter = range(x - 1, -1, -1)
    elif "R" in direction:
        x_iter = range(x + 1, x_size)
    else:
        x_iter = repeat(x)

    return zip(x_iter, y_iter)


def is_occupied_in_direction(seat: Tuple[int, int], direction: str, board: Board) -> bool:
    for x, y in iter_in_direction(seat, direction, board):
        neighbour = board[y][x]

        if neighbour == "#":
            return True
        elif neighbour == "L":
            return False

    return False


def count_visible(seat: Tuple[int, int], board: Board) -> int:
    return sum(
        1
        for direction in ["U", "D", "L", "R", "UL", "UR", "DL", "DR"]
        if is_occupied_in_direction(seat, direction, board)
    )


def take_seat(seat: Tuple[int, int], board: Board) -> str:
    x, y = seat
    seat_state = board[y][x]
    neighbours = count_neighbours(seat, board)

    if seat_state == "L" and neighbours == 0:
        return "#"
    elif seat_state == "#" and neighbours >= 4:
        return "L"
    else:
        return seat_state


def take_seat_2(seat: Tuple[int, int], board: Board) -> str:
    x, y = seat
    seat_state = board[y][x]
    visible = count_visible(seat, board)

    if seat_state == "L" and visible == 0:
        return "#"
    elif seat_state == "#" and visible >= 5:
        return "L"
    else:
        return seat_state


def count_occupied(board: Board) -> int:
    return sum(
        1
        for row in board
        for seat in row
        if seat == "#"
    )


def apply_until_stable(board: Board, fn: Callable[[Board], Tuple[Board, int]]) -> Board:
    next_board, moves = fn(board)

    while moves > 0:
        next_board, moves = fn(next_board)

    return next_board

def main():
    board = [
        [char for char in line]
        for line in read_lines(sys.argv[1])
    ]

    stable_board = apply_until_stable(board, map_board(take_seat))
    print(f"Part I: {count_occupied(stable_board)}")

    stable_board_2 = apply_until_stable(board, map_board(take_seat_2))
    print(f"Part II: {count_occupied(stable_board_2)}")

if __name__ == "__main__":
    main()
