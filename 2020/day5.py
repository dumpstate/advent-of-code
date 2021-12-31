import sys

from common.io import read_lines


def get_index(addr, total, up, down):
    def _ix(curr, min, max):
        if not curr or min == max:
            return min

        head, tail = curr[:1], curr[1:]
        span = max - min + 1

        if head == up:
            return _ix(tail, min + span // 2, max)
        elif head == down:
            return _ix(tail, min, max - span // 2)
        else:
            raise ValueError(f"Unexpected character: {head}")

    return _ix(addr, 1, total) - 1


def seat_id(row, col):
    return row * 8 + col


def main():
    lines = read_lines(sys.argv[1])
    seat_ids = [
        seat_id(
            row=get_index(ticket[:7], 128, "B", "F"),
            col=get_index(ticket[7:], 8, "R", "L"),
        )
        for ticket in lines
    ]

    print(f"Part I: {max(seat_ids)}")

    all_seats = {
        seat_id(row, col)
        for row in range(1, 128 + 1)
        for col in range(1, 8 + 1)
    }
    missing = all_seats.difference(seat_ids)

    for seat in missing:
        if (seat - 1) in missing or (seat + 1) in missing:
            continue

        print(f"Part II: {seat}")
        break


if __name__ == "__main__":
    main()
