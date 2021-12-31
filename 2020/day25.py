import sys
from itertools import count

from common.io import read_lines_as_ints


SUBJECT = 7
DIV = 20201227


def increment(div, subject, key):
    return (key * subject) % div


def encrypted(div, loop_size, subject):
    key = 1

    for _ in range(0, loop_size):
        key = increment(div, subject, key)

    return key


def brut_force_loop_size(div, subject, pub_key):
    key = 1

    for loop_size in count(1):
        key = increment(div, subject, key)

        if key == pub_key:
            return loop_size


def main():
    card_pub, door_pub = read_lines_as_ints(sys.argv[1])

    card_loop_size = brut_force_loop_size(DIV, SUBJECT, card_pub)
    door_loop_size = brut_force_loop_size(DIV, SUBJECT, door_pub)
    card_enc_key = encrypted(DIV, card_loop_size, door_pub)
    door_enc_key = encrypted(DIV, door_loop_size, card_pub)

    if card_enc_key != door_enc_key:
        raise Exception("Handshake failed")

    print(f"Part I: {card_enc_key}")


if __name__ == "__main__":
    main()
