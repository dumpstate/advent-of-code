import sys
from math import ceil
from typing import List, Tuple

from common.io import read_lines


def count_trees(grid: List[List[str]], slope: Tuple[int, int]) -> int:
    slope_x, slope_y = slope
    grid_size_x, grid_size_y = len(grid[0]), len(grid)
    trees = 0

    total_ix = ceil(grid_size_y / slope_y)
    x_ix = range(0, total_ix * slope_x, slope_x)
    y_ix = range(0, total_ix * slope_y, slope_y)

    for x, y in zip(x_ix, y_ix):
        if grid[y][x % grid_size_x] == "#":
            trees += 1

    return trees


def count_trees_product(grid: List[List[str]], slopes: List[Tuple[int, int]]) -> int:
    product = 1

    for slope in slopes:
        product *= count_trees(grid, slope)

    return product


def main():
    grid = [
        [char for char in line]
        for line in read_lines(sys.argv[1])
    ]
    slopes = [
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    ]

    print(f"Part I: {count_trees(grid, (3, 1))}")
    print(f"Part II: {count_trees_product(grid, slopes)}")


if __name__ == "__main__":
    main()
