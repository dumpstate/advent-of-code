import hashlib
import pickle
import sys
from copy import copy

from common.io import read_lines


def parse(lines):
    player_a = []
    player_b = []
    segment = 0

    for line in lines:
        if line == "":
            segment += 1
            continue

        if "Player" in line:
            continue

        deck = player_a if segment == 0 else player_b
        deck.append(int(line))

    return player_a, player_b


def digest(obj):
    return hashlib.md5(pickle.dumps(obj)).hexdigest()


def combat(a, b):
    a_, b_ = copy(a), copy(b)

    while a_ and b_:
        a_head, b_head = a_.pop(0), b_.pop(0)

        if a_head > b_head:
            a_.extend([a_head, b_head])
        elif a_head < b_head:
            b_.extend([b_head, a_head])
        else:
            raise ValueError(f"A duplicate {a_head}!")

    return a_ if a_ else b_


def recursive_combat(a, b):
    a_, b_ = copy(a), copy(b)
    prev_rounds = set()

    while a_ and b_:
        round_digest = digest((a_, b_))
        if round_digest in prev_rounds:
            return a_, b_

        prev_rounds.add(round_digest)

        a_head, b_head = a_.pop(0), b_.pop(0)

        if a_head <= len(a_) and b_head <= len(b_):
            a_sub, _ = recursive_combat(a_[:a_head], b_[:b_head])

            if a_sub:
                a_.extend([a_head, b_head])
            else:
                b_.extend([b_head, a_head])
        else:
            if a_head > b_head:
                a_.extend([a_head, b_head])
            else:
                b_.extend([b_head, a_head])

    return a_, b_


def score(deck):
    return sum(
        card * ix
        for card, ix in zip(reversed(deck), range(1, len(deck) + 1))
    )


def main():
    player_a, player_b = parse(read_lines(sys.argv[1]))

    winner_deck = combat(player_a, player_b)
    winner_score = score(winner_deck)

    print(f"Part I: {winner_score}")

    rec_a_res, rec_b_res = recursive_combat(player_a, player_b)
    rec_winner_score = score(rec_a_res or rec_b_res)

    print(f"Part II: {rec_winner_score}")


if __name__ == "__main__":
    main()
