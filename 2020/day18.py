import sys

from common.io import read_lines


def is_integer(string):
    try:
        int(string)
        return True
    except:
        return False


def parse_atom(line, pos):
    if pos < 0:
        raise ValueError("Out of bounds")

    char = line[pos]

    if is_integer(char):
        all_chars = ""

        while pos >= 0 and is_integer(line[pos]):
            all_chars += line[pos]
            pos -= 1

        return int(all_chars[::-1]), pos

    if char in ("+", "*"):
        return char, pos - 1

    raise ValueError(f"Oops; char={char}; pos={pos}")


def parse_expression_a(line, start_pos = None):
    if start_pos is None:
        start_pos = len(line) - 1

    if line[start_pos] == ")":
        left, next_pos = parse_expression_a(line, start_pos - 1)
    else:
        left, next_pos = parse_atom(line, start_pos)

    op, next_pos = parse_atom(line, next_pos)

    if line[next_pos] == ")":
        right, final_pos = parse_expression_a(line, next_pos - 1)
    else:
        right, final_pos = parse_atom(line, next_pos)

    if final_pos >= 0 and line[final_pos] == "(":
        right, final_pos = right, final_pos - 1
    elif final_pos >= 0:
        right, final_pos = parse_expression_a(line, next_pos)

    return (op, left, right), final_pos


def parse_expression_b(line, start_pos = None):
    print(f"line='{line}', {start_pos}")
    if start_pos is None:
        start_pos = len(line) - 1

    if line[start_pos] == ")":
        left, next_pos = parse_expression_b(line, start_pos - 1)
    else:
        left, next_pos = parse_atom(line, start_pos)

    print(f"\tleft='{left}'")

    op, next_pos = parse_atom(line, next_pos)
    print(f"\top={op}")

    if line[next_pos] == ")":
        right, final_pos = parse_expression_b(line, next_pos - 1)
    else:
        right, final_pos = parse_atom(line, next_pos)

    print(f"\tright={right}")

    if final_pos >= 0 and line[final_pos] == "(":
        right, final_pos = right, final_pos - 1
    elif final_pos >= 0 and op == "*":
        right, final_pos = parse_expression_b(line, next_pos)
    elif final_pos >= 0 and op == "+":
        print(f"\tprocessing + | {final_pos}")
        next_op, next_pos = parse_atom(line, final_pos)
        print(f"\tnext_op={next_op} | {next_pos}")
        if line[next_pos] == ")":
            next_right, final_pos = parse_expression_b(line, next_pos - 1)
        else:
            next_right, final_pos = parse_atom(line, next_pos)

        if final_pos >= 0 and line[final_pos] == "(":
            next_right, final_pos = next_right, final_pos - 1
        elif final_pos >= 0:
            next_right, final_pos = parse_expression_b(line, next_pos)

        return (next_op, (op, left, right), next_right), final_pos

    return (op, left, right), final_pos


def read(file_path):
    with open(file_path, "r") as f:
        return [
            "".join(line.strip().split())
            for line in f.readlines()
        ]


def evaluate(ex):
    if type(ex) is tuple:
        op, left, right = ex

        if op == "+":
            return evaluate(left) + evaluate(right)

        if op == "*":
            return evaluate(left) * evaluate(right)

    return ex


def sum_expr(expressions):
    total = 0

    for ex, _ in expressions:
        total += evaluate(ex)

    return total


def parse_expression(expr):
    def parse(acc, pos):
        if pos >= len(expr):
            return acc, pos

        if is_integer(expr[pos]):
            all_chars = ""

            while pos < len(expr) and is_integer(expr[pos]):
                all_chars += expr[pos]
                pos += 1

            return parse(acc + [int(all_chars)], pos)

        if expr[pos] in ("+", "*"):
            return parse(acc + [expr[pos]], pos + 1)

        if expr[pos] == "(":
            nested, next_pos = parse([], pos + 1)
            return parse(acc + [nested], next_pos)

        if expr[pos] == ")":
            return acc, pos + 1

    return parse([], 0)[0]


def eval_expr_a(expr):
    total = expr[0] if type(expr[0]) is not list else eval_expr_a(expr[0])

    for i in range(1, len(expr) - 1, 2):
        op = expr[i]
        next_ = expr[i + 1] if type(expr[i + 1]) is not list else eval_expr_a(expr[i + 1])

        if op == "+":
            total += next_
        elif op == "*":
            total *= next_
        else:
            raise ValueError(f"Unknown operator: {op}")

    return total


def eval_expr_b(expr):
    total = expr[0] if type(expr[0]) is not list else eval_expr_b(expr[0])
    tmp = []

    for i in range(1, len(expr) - 1, 2):
        op = expr[i]
        next_ = expr[i + 1] if type(expr[i + 1]) is not list else eval_expr_b(expr[i + 1])

        if op == "+":
            total += next_
        elif op == "*":
            tmp += [total, op]
            total = next_
        else:
            raise ValueError(f"Unknown operator: {op}")

        if i >= len(expr) - 2:
            tmp += [total]

    total = tmp[0] if type(tmp[0]) is not list else eval_expr_b(tmp[0])

    for i in range(1, len(tmp) - 1, 2):
        op = tmp[i]
        next_ = tmp[i + 1] if type(tmp[i + 1]) is not list else eval_expr_b(tmp[i + 1])

        if op == "*":
            total *= next_
        else:
            raise ValueError(f"Unknown operator: {op}")

    return total


def main():
    expressions = [
        parse_expression("".join(line.strip().split()))
        for line in read_lines(sys.argv[1])
    ]

    print(f"Part I: {sum(eval_expr_a(expr) for expr in expressions)}")
    print(f"Part II: {sum(eval_expr_b(expr) for expr in expressions)}")


if __name__ == "__main__":
    main()
