import sys
from typing import List, Tuple

from common.io import read_lines


Coord = Tuple[int, int]
Command = Tuple[str, int]


DIRECTIONS = ["N", "E", "S", "W"]


def move(pos: Coord, direction: str, value: int):
    x, y = pos

    if direction == "N":
        return (x, y + value)
    elif direction == "S":
        return (x, y - value)
    elif direction == "E":
        return (x + value, y)
    elif direction == "W":
        return (x - value, y)
    else:
        raise ValueError(f"Unsupported direction: {direction}")


def rotation_value(command: str, angle: int) -> int:
    return angle // 90 * (1 if command == "R" else -1)


def rotate(direction: str, command: str, angle: int) -> str:
    ix = DIRECTIONS.index(direction) + rotation_value(command, angle)

    return DIRECTIONS[ix % len(DIRECTIONS)]


def rotate_around(pos: Coord, command: str, angle: int) -> Coord:
    rotation_dir = rotation_value(command, angle) % len(DIRECTIONS)
    x, y = pos

    if rotation_dir == 0:
        return pos
    elif rotation_dir == 1:
        return (y, -x)
    elif rotation_dir == 2:
        return (-x, -y)
    elif rotation_dir == 3:
        return (-y, x)
    else:
        raise ValueError(f"Unsupported rotation direction: {rotation_dir}")


def move_ship_a(commands: List[Command], init: Coord, init_direction: str) -> Coord:
    pos = init
    direction = init_direction

    for command, value in commands:
        if command == "F":
            pos = move(pos, direction, value)
        elif command in DIRECTIONS:
            pos = move(pos, command, value)
        elif command in ["L", "R"]:
            direction = rotate(direction, command, value)
        else:
            raise ValueError(f"Unsupported command: ({command}, {value})")

    return pos


def move_ship_b(commands: List[Command], init_ship: Coord, init_waypoint: Coord) -> Coord:
    ship_pos = init_ship
    waypoint_rel_pos = init_waypoint

    for command, value in commands:
        if command == "F":
            x, y = ship_pos
            x_w, y_w = waypoint_rel_pos

            ship_pos = (value * x_w + x, value * y_w + y)
        elif command in DIRECTIONS:
            waypoint_rel_pos = move(waypoint_rel_pos, command, value)
        elif command in ["L", "R"]:
            waypoint_rel_pos = rotate_around(waypoint_rel_pos, command, value)
        else:
            raise ValueError(f"Unsupported command: ({command}, {value})")

    return ship_pos


def manhattan(first: Coord, second: Coord) -> int:
    return sum(abs(b - a) for a, b in zip(first, second))


def main():
    commands = [
        (line[0], int(line[1:]))
        for line in read_lines(sys.argv[1])
    ]

    ship_position_a = move_ship_a(commands, (0, 0), "E")
    print(f"Part I: {manhattan((0, 0), ship_position_a)}")

    ship_position_b = move_ship_b(commands, (0, 0), (10, 1))
    print(f"Part II: {manhattan((0, 0), ship_position_b)}")


if __name__ == "__main__":
    main()
