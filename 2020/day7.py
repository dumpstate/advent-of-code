import sys
from typing import Dict, List, Tuple

from common.io import read_lines


def parse(lines: List[str]) -> Dict[str, Dict[str, int]]:
    bags = dict()

    for line in lines:
        color, desc = line.split(" bags contain ")
        if desc == "no other bags.":
            bags[color] = dict()
        else:
            for contained_bag in desc.split(", "):
                count, *color_words = (
                    contained_bag
                    .replace(" bags.", "")
                    .replace(" bag.", "")
                    .replace(" bags", "")
                    .replace(" bag", "")
                ).split()

                if color not in bags:
                    bags[color] = dict()

                bags[color][" ".join(color_words)] = int(count)

    return bags


def expand_all(bags: Dict[str, Dict[str, int]]) -> Dict[str, Dict[str, int]]:
    acc = dict()
    queue: List[Tuple[str, str, int]] = [
        (color, color, 1)
        for color in bags.keys()
    ]

    while queue:
        top_lvl_color, next_color, count = queue.pop()

        for contained_color, contained_count in bags[next_color].items():
            if top_lvl_color not in acc:
                acc[top_lvl_color] = dict()
            
            if contained_color not in acc[top_lvl_color]:
                acc[top_lvl_color][contained_color] = 0

            acc[top_lvl_color][contained_color] += count * contained_count

            if bags[contained_color]:
                queue.append((top_lvl_color, contained_color, contained_count * count))

    return acc


def main():
    bags = parse(read_lines(sys.argv[1]))
    expanded = expand_all(bags)

    total_shiny_gold = sum(
        1
        for contained_colors in expanded.values()
        if "shiny gold" in contained_colors
    )

    print(f"Part I: {total_shiny_gold}")

    total_contained_bags = sum(expanded["shiny gold"].values())

    print(f"Part II: {total_contained_bags}")


if __name__ == "__main__":
    main()
