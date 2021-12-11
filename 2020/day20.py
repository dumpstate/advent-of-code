import math
import sys
from dataclasses import dataclass
from typing import List

from common.io import read_lines


PATTERN = """                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
"""


def joined(edge):
    return "".join(edge)


def product(ls):
    res, *tail = ls

    for value in tail:
        res *= value

    return res


@dataclass()
class Tile:
    id: int
    points: List[List[str]]

    def __str__(self):
        res_str = f"Tile {self.id}:\n"

        for y in range(0, self.y_size):
            for x in range(0, self.x_size):
                res_str += self.points[y][x]

            res_str += "\n"

        return res_str

    @property
    def x_size(self):
        return len(self.points[0])

    @property
    def y_size(self):
        return len(self.points)

    def row(self, ix):
        return self.points[ix]

    def col(self, ix):
        column = []

        for y in range(0, self.y_size):
            column.append(self.points[y][ix])

        return column

    @property
    def edge_up(self):
        return joined(self.row(0))

    @property
    def edge_down(self):
        return joined(self.row(-1))

    @property
    def edge_left(self):
        return joined(self.col(0))

    @property
    def edge_right(self):
        return joined(self.col(-1))

    def iter_edges(self):
        yield self.edge_up
        yield self.edge_down
        yield self.edge_left
        yield self.edge_right

    def iter_flipped_edges(self):
        for edge in self.iter_edges():
            yield joined(reversed(edge))

    @property
    def all_edges(self):
        edges = set(self.iter_edges())

        edges.update(self.iter_flipped_edges())

        return edges

    def rotate(self):
        x_s, y_s = self.x_size, self.y_size
        rot_x_s, rot_y_s = y_s, x_s
        res = [["."] * rot_x_s for _ in range(0, rot_y_s)]

        for y in range(0, y_s):
            for x in range(0, x_s):
                res[x][rot_x_s - y - 1] = self.points[y][x]

        self.points = res

    def flip_x(self):
        x_s, y_s = self.x_size, self.y_size
        res = [["."] * x_s for _ in range(0, y_s)]

        for y in range(0, y_s):
            for x in range(0, x_s):
                res[y][x_s - x - 1] = self.points[y][x]

        self.points = res

    def flip_y(self):
        x_s, y_s = self.x_size, self.y_size
        res = [["."] * y_s for _ in range(0, y_s)]

        for y in range(0, y_s):
            for x in range(0, x_s):
                res[y_s - y - 1][x] = self.points[y][x]

        self.points = res


def parse(lines):
    tiles = []
    tile_id = None
    tile = []

    for line in lines:
        if line == "":
            tiles.append(Tile(id=tile_id, points=tile))
            tile_id = None
            tile = []
            continue

        if line.startswith("Tile"):
            _, id_str = line.split()
            tile_id = int(id_str[:-1])
            continue

        tile.append([
            char
            for char in line
        ])

    tiles.append(Tile(id=tile_id, points=tile))

    return tiles


def count_intersecting_edges(tile_id, tile_edges, edges):
    intersecting_edges = set()

    for other_tile_id, other_edges in edges.items():
        if other_tile_id == tile_id:
            continue

        for intersecting_edge in other_edges.intersection(tile_edges):
            intersecting_edges.add(intersecting_edge)

    return len(intersecting_edges)


def find_border_edges(tiles):
    edges = dict()

    for tile in tiles:
        for edge in tile.all_edges:
            curr_count = edges.get(edge, 0)
            edges[edge] = curr_count + 1

    return {
        edge
        for edge, count in edges.items()
        if count == min(edges.values())
    }


def segregate_tiles(tiles):
    edges = {
        tile.id: tile.all_edges
        for tile in tiles
    }
    intersecting_edges = {
        tile_id: count_intersecting_edges(tile_id, tile_edges, edges)
        for tile_id, tile_edges in edges.items()
    }

    res = [
        [
            tile
            for tile in tiles
            if intersecting_edges[tile.id] == count
        ]
        for count in list(sorted(set(intersecting_edges.values())))
    ]

    if len(res) == 1:
        return (res[0], [], [])

    return tuple(res)


def iter_border_pos(board, board_offset):
    board_dim = len(board) - board_offset

    for x_board in range(board_offset, board_dim):
        pos = (x_board, board_offset)
        is_corner = x_board == board_offset or x_board == board_dim - 1

        yield (is_corner, pos)

    for y_board in range(board_offset + 1, board_dim):
        pos = (board_dim - 1, y_board)
        is_corner = y_board == board_offset or y_board == board_dim - 1

        yield (is_corner, pos)

    for x_board in range(board_dim - 2, board_offset - 1, -1):
        pos = (x_board, board_dim - 1)
        is_corner = x_board == board_offset or x_board == board_dim - 1

        yield (is_corner, pos)

    for y_board in range(board_dim - 2, board_offset, -1):
        pos = (board_offset, y_board)
        is_corner = y_board == board_offset or y_board == board_dim - 1

        yield (is_corner, pos)


def check_up_edge(board, border_edges, tile_pos, tile):
    x_board, y_board = tile_pos
    x_n, y_n = x_board, y_board - 1

    if y_n < 0:
        return tile.edge_up in border_edges

    return (
        board[y_n][x_n] is None or
        board[y_n][x_n].edge_down == tile.edge_up
    )


def check_down_edge(board, border_edges, tile_pos, tile):
    x_board, y_board = tile_pos
    x_n, y_n = x_board, y_board + 1

    if y_n >= len(board):
        return tile.edge_down in border_edges

    return (
        board[y_n][x_n] is None or
        board[y_n][x_n].edge_up == tile.edge_down
    )


def check_left_edge(board, border_edges, tile_pos, tile):
    x_board, y_board = tile_pos
    x_n, y_n = x_board - 1, y_board

    if x_n < 0:
        return tile.edge_left in border_edges

    return (
        board[y_n][x_n] is None or
        board[y_n][x_n].edge_right == tile.edge_left
    )


def check_right_edge(board, border_edges, tile_pos, tile):
    x_board, y_board = tile_pos
    x_n, y_n = x_board + 1, y_board

    if x_n >= len(board[0]):
        return tile.edge_right in border_edges

    return (
        board[y_n][x_n] is None or
        board[y_n][x_n].edge_left == tile.edge_right
    )


def check_tile(board, border_edges, tile_pos):
    x_board, y_board = tile_pos
    tile = board[y_board][x_board]
    if tile is None:
        return True

    return (
        check_up_edge(board, border_edges, tile_pos, tile) and
        check_down_edge(board, border_edges, tile_pos, tile) and
        check_left_edge(board, border_edges, tile_pos, tile) and
        check_right_edge(board, border_edges, tile_pos, tile)
    )


def check_board(board, border_edges):
    for y_board in range(0, len(board)):
        for x_board in range(0, len(board[0])):
            if board[y_board][x_board] is None:
                continue

            if not check_tile(board, border_edges, (x_board, y_board)):
                return False

    return True


def try_tile(board, border_edges, board_pos, tile):
    x_board, y_board = board_pos
    current = board[y_board][x_board]
    board[y_board][x_board] = tile

    for _ in range(0, 2):
        for _ in range(0, 2):
            for _ in range(0, 4):
                if check_board(board, border_edges):
                    return True
    
                tile.rotate()

            tile.flip_y()

        tile.flip_x()

    board[y_board][x_board] = current
    return False


def layout_tiles(board, border_edges, board_offset, tiles):
    if len(tiles) == 0:
        return board

    if len(tiles) == 1:
        tile = tiles[0]
        try_tile(board, border_edges, (board_offset, board_offset), tile)
        return board

    corner_tiles, edge_tiles, inner_tiles = segregate_tiles(tiles)

    for is_corner, board_pos in iter_border_pos(board, board_offset):
        x_board, y_board = board_pos
        if board[y_board][x_board] is not None:
            continue

        tiles_to_pick = corner_tiles if is_corner else edge_tiles

        for tile_ix in range(0, len(tiles_to_pick)):
            tile = tiles_to_pick[tile_ix]

            if try_tile(board, border_edges, board_pos, tile):
                del tiles_to_pick[tile_ix]
                break

    return layout_tiles(board, border_edges, board_offset + 1, inner_tiles)


def merge(board):
    x_tile_size, y_tile_size = board[0][0].x_size - 2, board[0][0].y_size - 2
    res_x_s, res_y_s = len(board) * x_tile_size, len(board[0]) * y_tile_size
    points = [["."] * res_x_s for _ in range(0, res_y_s)]

    for y in range(0, len(board) * y_tile_size):
        for x in range(0, len(board[0]) * x_tile_size):
            x_board, y_board = x // x_tile_size, y // y_tile_size
            x_tile, y_tile = (x % x_tile_size) + 1, (y % y_tile_size) + 1

            points[y][x] = board[y_board][x_board].points[y_tile][x_tile]

    return Tile(id=1, points=points)


def draw_pattern(points, pattern, point):
    p_x, p_y = point
    pattern_points = []

    for y in range(0, len(pattern)):
        for x in range(0, len(pattern[0])):
            if pattern[y][x] == "":
                continue

            if pattern[y][x] != points[p_y + y][p_x + x]:
                return False

            pattern_points.append((p_x + x, p_y + y))

    for x, y in pattern_points:
        points[y][x] = "O"

    return True


def find_pattern(tile, pattern):
    p_x_size, p_y_size = len(pattern[0]), len(pattern)

    for y in range(0, tile.y_size - p_y_size):
        for x in range(0, tile.x_size - p_x_size):
            draw_pattern(tile.points, pattern, (x, y))


def count(points, value):
    total = 0

    for y in range(0, len(points)):
        for x in range(0, len(points[0])):
            if points[y][x] == value:
                total += 1

    return total


def main():
    tiles = parse(read_lines(sys.argv[1]))
    layout_dim = int(math.sqrt(len(tiles)))
    corner_tiles, _, _ = segregate_tiles(tiles)

    print(f"Part I: {product(tile.id for tile in corner_tiles)}")

    board = [[None] * layout_dim for _ in range(0, layout_dim)]
    border_edges = find_border_edges(tiles)
    layout_tiles(board, border_edges, 0, tiles)
    merged = merge(board)
    pattern = [
        [char.strip() for char in line]
        for line in PATTERN.split("\n")
        if line != ""
    ]
    find_pattern(merged, pattern)

    print(f"Part II: {count(merged.points, '#')}")


if __name__ == "__main__":
    main()
