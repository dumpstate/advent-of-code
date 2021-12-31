import sys

from common.io import read_lines


def parse(lines):
    products = []

    for line in lines:
        if " (contains " not in line:
            products.append((set(line.split()), set()))
        else:
            ingr_str, contains_str = line.split(" (contains ")
            products.append((set(ingr_str.split()), set(contains_str[:-1].split(", "))))

    return products


def find_allergen_free_products(products):
    allergens = {
        allergen
        for _, prod_allergens in products
        for allergen in prod_allergens
    }
    potentially_taken = set()
    ingr_by_allergen = dict()

    for a in allergens:
        potentially_contain = set()

        for p_i, p_a in products:
            if a not in p_a:
                continue

            if not potentially_contain:
                potentially_contain = p_i
            else:
                potentially_contain = potentially_contain.intersection(p_i)

        for i in potentially_contain:
            potentially_taken.add(i)

        ingr_by_allergen[a] = potentially_contain

    return (
        {
            ingr
            for ingrs, _ in products
            for ingr in ingrs
            if ingr not in potentially_taken
        },
        ingr_by_allergen
    )


def count_ingrs(products, ingrs):
    return sum(
        1
        for ingredients, _ in products
        for ingr in ingrs
        if ingr in ingredients

    )


def resolve_allergens(ingr_by_allergen):
    taken = set()
    res = dict()

    while len(res) < len(ingr_by_allergen):
        for a, ingrs in ingr_by_allergen.items():
            potential_ingrs = ingrs.difference(taken)

            if len(potential_ingrs) == 1:
                ingr = list(potential_ingrs)[0]
                taken.add(ingr)
                res[a] = ingr

    return ",".join([res[a] for a in sorted(res.keys())])


def main():
    products = parse(read_lines(sys.argv[1]))

    allergen_free, ingr_by_allergen = find_allergen_free_products(products)
    print(f"Part I: {count_ingrs(products, allergen_free)}")

    allergens = resolve_allergens(ingr_by_allergen)
    print(f"Part II: {allergens}")


if __name__ == "__main__":
    main()
