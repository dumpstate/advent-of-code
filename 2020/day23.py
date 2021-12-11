import sys

from common.io import read_lines


def find_dest_value(cups, current_value):
    dest_value = current_value - 1

    while True:
        if dest_value in cups:
            break

        if dest_value < 1:
            dest_value = max(cups.keys())
            continue

        dest_value -= 1

    return dest_value


def pick_up_cups(cups, current_pos):
    pick_up_pos = [
        ix % len(cups)
        for ix in range(current_pos + 1, current_pos + 4)
    ]
    pick_up_cups = [cups[pos] for pos in pick_up_pos]

    for pos in sorted(pick_up_pos, reverse=True):
        del cups[pos]

    return pick_up_cups


def pick_up_cups_(cups, current_value):
    p1 = cups[current_value]
    p2 = cups[p1]
    p3 = cups[p2]

    cups[current_value] = cups[p3]
    del cups[p1]
    del cups[p2]
    del cups[p3]

    return [p1, p2, p3]


def flatten(cups, curr):
    flat_cups = []
    it = curr

    for _ in range(0, len(cups)):
        flat_cups.append(it)
        it = cups[it]

    return flat_cups


def play_game(cups, moves):
    cups__ = dict(zip(cups, cups[1:] + [cups[0]]))
    curr = None

    for _ in range(1, moves + 1):
        if curr is None:
            curr = cups[0]
        else:
            curr = cups__[curr]

        pick_up = pick_up_cups_(cups__, curr)
        dest_value = find_dest_value(cups__, curr)

        prev_dest_value = cups__[dest_value]
        cups__[dest_value] = pick_up[0]
        cups__[pick_up[0]] = pick_up[1]
        cups__[pick_up[1]] = pick_up[2]
        cups__[pick_up[2]] = prev_dest_value

    return flatten(cups__, curr)


def score(cups):
    ix = cups.index(1)

    return int("".join(
        str(i)
        for i in cups[ix + 1:] + cups[:ix]
    ))


def find_stars(cups, value):
    ix = cups.index(value)

    return cups[ix + 1] * cups[ix + 2]


def main():
    cups = [
        int(char)
        for line in read_lines(sys.argv[1])
        for char in line
    ]

    cups_a = play_game(cups, 100)
    print(f"Part I: {score(cups_a)}")

    cups_b = play_game(cups + list(range(max(cups) + 1, 10 ** 6 + 1)), 10 ** 7)
    print(f"Part II: {find_stars(cups_b, 1)}")


if __name__ == "__main__":
    main()
