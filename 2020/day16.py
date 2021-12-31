import sys
from copy import copy
from dataclasses import dataclass
from typing import Dict, List, Tuple

from common.io import read_lines


Rule = List[Tuple[int, int]]


@dataclass(frozen=True)
class TicketNotes:
    rules: Dict[str, Rule]
    my_ticket: List[int]
    nearby_tickets: List[List[int]]


def parse(lines: List[str]) -> TicketNotes:
    rules = dict()
    my_ticket = []
    nearby_tickets = []
    section = 0

    for line in lines:
        next_line = line.strip()

        if next_line == "":
            section += 1
            continue

        if section == 0:
            desc, rule = next_line.split(": ")
            l_rule, r_rule = rule.split(" or ")
            l_rule_d, l_rule_u = l_rule.split("-")
            r_rule_d, r_rule_u = r_rule.split("-")
            rules[desc] = [
                (int(l_rule_d), int(l_rule_u)),
                (int(r_rule_d), int(r_rule_u)),
            ]

        if section == 1:
            if next_line == "your ticket:":
                continue

            my_ticket = [
                int(num)
                for num in next_line.split(",")
            ]

        if section == 2:
            if next_line == "nearby tickets:":
                continue

            nearby_tickets.append([
                int(num)
                for num in next_line.split(",")
            ])

    return TicketNotes(
        rules=rules,
        my_ticket=my_ticket,
        nearby_tickets=nearby_tickets,
    )


def is_valid(ticket_num: int, rule: Rule) -> bool:
    return any(
        ticket_num >= rule_section[0] and ticket_num <= rule_section[1]
        for rule_section in rule
    )


def ticket_err_rate(rules: Dict[str, Rule], ticket: List[int]) -> int:
    return sum(
        num
        for num in ticket
        if not any(is_valid(num, rule) for rule in rules.values())
    )


def total_ticket_err_rate(notes: TicketNotes) -> int:
    return sum(ticket_err_rate(notes.rules, ticket) for ticket in notes.nearby_tickets)


def eliminate_non_matching_rules(rules: Dict[str, Rule], tickets: List[List[int]]) -> List[List[str]]:
    keys = list(rules.keys())
    len_ = len(keys)
    valid_tickets = [
        ticket
        for ticket in tickets
        if ticket_err_rate(rules, ticket) == 0
    ]
    keys_space = [
        copy(keys)
        for _ in range(0, len_)
    ]

    for ticket in valid_tickets:
        for ticket_num, avail_keys in zip(ticket, keys_space):
            to_remove = []

            for key in avail_keys:
                if not is_valid(ticket_num, rules[key]):
                    to_remove.append(key)

            for key in to_remove:
                avail_keys.remove(key)

    changes = 1

    while changes > 0:
        changes = 0
        single_match_keys = [
            key
            for keys in keys_space
            if len(keys) == 1
            for key in keys
        ]

        for keys in keys_space:
            if len(keys) == 1:
                continue

            for single_match_key in single_match_keys:
                if single_match_key in keys:
                    keys.remove(single_match_key)
                    changes += 1

    return keys_space


def total_departure_value(notes: TicketNotes) -> int:
    keys_space = eliminate_non_matching_rules(
        notes.rules,
        notes.nearby_tickets + [notes.my_ticket],
    )
    total = 1

    for ticket_num, avail_keys in zip(notes.my_ticket, keys_space):
        if len(avail_keys) != 1:
            continue

        if not avail_keys[0].startswith("departure"):
            continue

        total *= ticket_num

    return total


def main():
    notes = parse(read_lines(sys.argv[1]))

    print(f"Part I: {total_ticket_err_rate(notes)}")
    print(f"Part II: {total_departure_value(notes)}")


if __name__ == "__main__":
    main()
