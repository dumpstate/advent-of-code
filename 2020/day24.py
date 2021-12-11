import sys
from copy import copy

from common.io import read_lines


WHITE = 0
BLACK = 1


def to_instruction(line):
    i = 0
    instruction = []

    while i < len(line):
        if line[i:i + 2] in ("se", "sw", "ne", "nw"):
            instruction.append(line[i:i + 2])
            i += 2
        else:
            instruction.append(line[i])
            i += 1

    return instruction


def move(instr, pos):
    if not instr:
        return pos

    head, *tail = instr
    x, y = pos

    if head == "e":
        return move(tail, (x + 2, y))
    if head == "w":
        return move(tail, (x - 2, y))
    if head == "nw":
        return move(tail, (x - 1, y + 1))
    if head == "ne":
        return move(tail, (x + 1, y + 1))
    if head == "se":
        return move(tail, (x + 1, y - 1))
    if head == "sw":
        return move(tail, (x - 1, y - 1))
    
    raise ValueError(f"Unknown direction: {head}")


def create_board(instructions):
    board = dict()

    for instr in instructions:
        pos = move(instr, (0, 0))
        board[pos] = (board.get(pos, WHITE) + 1) % 2

    return board


def count(it, pred):
    return sum(1 for val in it if pred(val))


def count_color(board, color):
    return count(board.values(), lambda value: value == color)


def iter_neighbour_pos(pos):
    x, y = pos

    yield (x + 2, y)
    yield (x - 2, y)
    yield (x + 1, y + 1)
    yield (x - 1, y + 1)
    yield (x + 1, y - 1)
    yield (x - 1, y - 1)


def count_neighbours(board, pos, color):
    return count(
        iter_neighbour_pos(pos),
        lambda n_pos: board.get(n_pos, WHITE) == color,
    )


def flip_tile(board, pos):
    color = board.get(pos, WHITE)
    ns = count_neighbours(board, pos, BLACK)

    if color == BLACK:
        if ns == 0 or ns > 2:
            return WHITE
    elif color == WHITE:
        if ns == 2:
            return BLACK
    else:
        raise ValueError(f"Unknown color: {color}")

    return color


def iter_board(board):
    visited = set()

    for pos in board.keys():
        if pos not in visited:
            yield pos
            visited.add(pos)

        for n_pos in iter_neighbour_pos(pos):
            if n_pos not in visited:
                yield n_pos
                visited.add(n_pos)


def flip_board(board):
    board_ = copy(board)

    for tile_pos in iter_board(board):
        tile_color = flip_tile(board, tile_pos)
        if tile_color == WHITE and tile_pos not in board_:
            continue

        board_[tile_pos] = tile_color

    return board_


def animate(board, times):
    board_ = board

    for _ in range(0, times):
        board_ = flip_board(board_)

    return board_


def main():
    instructions = [
        to_instruction(line)
        for line in read_lines(sys.argv[1])
    ]
    board = create_board(instructions)

    print(f"Part I: {count_color(board, BLACK)}")

    animated_board = animate(board, 100)
    print(f"Part II: {count_color(animated_board, BLACK)}")


if __name__ == "__main__":
    main()
