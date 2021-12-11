import sys
from copy import deepcopy
from itertools import permutations
from random import shuffle

from common.io import read_lines


def parse(lines):
    rules = dict()
    messages = []
    section = 0

    for line in lines:
        next_line = line.strip()

        if next_line == "":
            section += 1
            continue

        if section == 0:
            id_, def_ = next_line.split(": ")
            if "\"" in def_:
                rules[int(id_)] = def_.replace("\"", "")
            else:
                rules[int(id_)] = [
                    [
                        int(ref_id)
                        for ref_id in segment.split()
                    ]
                    for segment in def_.split(" | ")
                ]
        elif section == 1:
            messages.append(next_line)
        else:
            raise ValueError(f"Unexpected section {section}")

    return rules, messages


def matches(msg, rules, rule_id):
    def matches_at_pos(pos, rule):
        if pos == len(msg):
            return False, pos

        seg_pos = 0

        if type(rule) is str:
            if msg[pos] == rule:
                return True, pos + 1
            else:
                return False, pos + 1
        elif type(rule) is list:
            if (
                rule == [[42], [42, 8]] or
                rule == [[42, 31], [42, 11, 31]]
            ):
                segs_rnd = list(permutations(deepcopy(rule)))
                shuffle(segs_rnd)
            else:
                segs_rnd = [rule]
            
            for segs in segs_rnd:
                seg_pos = pos

                for rule_seg in segs:
                    next_pos = seg_pos
                    success = True

                    for entry in rule_seg:
                        res, next_pos = matches_at_pos(next_pos, rules[entry])

                        if not res:
                            success = False
                            break

                    if success:
                        return True, next_pos

            return False, seg_pos + 1
        else:
            raise ValueError(f"Unsupported rule: {rule}")

    success, count = matches_at_pos(0, rules[rule_id])

    return success and len(msg) <= count


def count_valid(messages, rules, rule_id, repeat):
    total = 0

    for msg in messages:
        for _ in range(0, repeat):
            res = matches(msg, rules, rule_id)

            if res:
                total += 1
                break

    return total


def main():
    rules, messages = parse(read_lines(sys.argv[1]))

    print(f"Part I: {count_valid(messages, rules, 0, 1)}")

    rules[8] = [[42], [42, 8]]
    rules[11] = [[42, 31], [42, 11, 31]]

    print(f"Part II: {count_valid(messages, rules, 0, 1000)}")


if __name__ == "__main__":
    main()
