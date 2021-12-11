import sys
from copy import deepcopy
from typing import Dict, List

from common.io import read_lines


def parse(lines: List[str]) -> Dict[int, Dict[int, str]]:
    cubes = { 0: dict() }
    line_ix = 0

    for line in lines:
        line_dct = dict()
        pos_ix = 0

        for char in line.strip():
            line_dct[pos_ix] = char
            pos_ix += 1

        cubes[0][line_ix] = line_dct
        line_ix += 1

    return cubes


def hyper_w_size(hypercube):
    return len(hypercube)


def hyper_z_size(hypercube):
    return len(hypercube[0])


def hyper_y_size(hypercube):
    return len(hypercube[0][0])


def hyper_x_size(hypercube):
    return len(hypercube[0][0][0])


def z_size(cubes):
    return len(cubes)


def y_size(cubes):
    return len(cubes[0])


def x_size(cubes):
    return len(cubes[0][0])


def hyper_w_range(hcube):
    return sorted(hcube.keys())


def hyper_z_range(hcube):
    return sorted(hcube[0].keys())


def hyper_y_range(hcube):
    return sorted(hcube[0][0].keys())


def hyper_x_range(hcube):
    return sorted(hcube[0][0][0].keys())


def z_range(cubes):
    return sorted(cubes.keys())


def y_range(cubes):
    return sorted(cubes[0].keys())


def x_range(cubes):
    return sorted(cubes[0][0].keys())


def ext(rng):
    base_range = list(rng)

    return range(min(base_range) - 1, max(base_range) + 2)


def apply_n_cycles(cubes, fn, n):
    result_cubes = cubes

    for i in range(0, n):
        result_cubes = fn(result_cubes)

    return result_cubes


def map_cubes(fn):
    def map_(cubes):
        result_cubes = deepcopy(cubes)

        for z in ext(z_range(cubes)):
            for y in ext(y_range(cubes)):
                for x in ext(x_range(cubes)):
                    if z not in result_cubes:
                        result_cubes[z] = dict()

                    if y not in result_cubes[z]:
                        result_cubes[z][y] = dict()

                    result_cubes[z][y][x] = fn((x, y, z), cubes)

        return result_cubes

    return map_


def map_hypercube(fn):
    def map_(hypercube):
        result_hypercube = deepcopy(hypercube)

        for w in ext(hyper_w_range(hypercube)):
            for z in ext(hyper_z_range(hypercube)):
                for y in ext(hyper_y_range(hypercube)):
                    for x in ext(hyper_x_range(hypercube)):
                        if w not in result_hypercube:
                            result_hypercube[w] = dict()

                        if z not in result_hypercube[w]:
                            result_hypercube[w][z] = dict()

                        if y not in result_hypercube[w][z]:
                            result_hypercube[w][z][y] = dict()

                        result_hypercube[w][z][y][x] = fn((x, y, z, w), hypercube)

        return result_hypercube

    return map_


def iter_neighbours(pos, cubes):
    x, y, z = pos
    z_values = list(z_range(cubes))
    y_values = list(y_range(cubes))
    x_values = list(x_range(cubes))

    for k in range(max(z - 1, min(z_values)), min(z + 2, max(z_values) + 1)):
        for j in range(max(y - 1, min(y_values)), min(y + 2, max(y_values) + 1)):
            for i in range(max(x - 1, min(x_values)), min(x + 2, max(x_values) + 1)):
                if k == z and j == y and i == x:
                    continue

                yield cubes[k][j][i]


def iter_hneighbours(pos, hcube):
    x, y, z, w = pos
    w_values = list(hyper_w_range(hcube))
    z_values = list(hyper_z_range(hcube))
    y_values = list(hyper_y_range(hcube))
    x_values = list(hyper_x_range(hcube))

    for l in range(max(w - 1, min(w_values)), min(w + 2, max(w_values) + 1)):
        for k in range(max(z - 1, min(z_values)), min(z + 2, max(z_values) + 1)):
            for j in range(max(y - 1, min(y_values)), min(y + 2, max(y_values) + 1)):
                for i in range(max(x - 1, min(x_values)), min(x + 2, max(x_values) + 1)):
                    if l == w and k == z and j == y and i == x:
                        continue

                    yield hcube[l][k][j][i]


def iter_cubes(cubes):
    for z in z_range(cubes):
        for y in y_range(cubes):
            for x in x_range(cubes):
                yield cubes[z][y][x]


def iter_hcube(hcube):
    for w in hyper_w_range(hcube):
        for z in hyper_z_range(hcube):
            for y in hyper_y_range(hcube):
                for x in hyper_x_range(hcube):
                    yield hcube[w][z][y][x]


def count_active(iterator):
    return sum(1 for cube in iterator if cube == "#")


def change_state_a(pos, cubes):
    x, y, z = pos
    cube_state = cubes.get(z, {}).get(y, {}).get(x, ".")
    neighbours = count_active(iter_neighbours(pos, cubes))

    if cube_state == "#":
        if neighbours not in (2, 3):
            return "."
    else:
        if neighbours == 3:
            return "#"

    return cube_state


def change_state_b(pos, hcube):
    x, y, z, w = pos
    cube_state = hcube.get(w, {}).get(z, {}).get(y, {}).get(x, ".")
    neighbours = count_active(iter_hneighbours(pos, hcube))

    if cube_state == "#":
        if neighbours not in (2, 3):
            return "."
    else:
        if neighbours == 3:
            return "#"

    return cube_state


def main():
    cubes = parse(read_lines(sys.argv[1]))

    cubes_a = apply_n_cycles(cubes, map_cubes(change_state_a), 6)
    print(f"Part I: {count_active(iter_cubes(cubes_a))}")

    hypercube_b = apply_n_cycles({0: cubes}, map_hypercube(change_state_b), 6)
    print(f"Part II: {count_active(iter_hcube(hypercube_b))}")


if __name__ == "__main__":
    main()
