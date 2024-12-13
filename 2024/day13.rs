mod aoc;

#[derive(Debug)]
struct Machine {
    a: (i64, i64),
    b: (i64, i64),
    prize: (i64, i64),
}

fn parse_line(line: &String) -> (i64, i64) {
    let vec = line.split(": ").collect::<Vec<&str>>()[1]
        .split(", ")
        .map(|x| x[2..].parse::<i64>().unwrap())
        .collect::<Vec<i64>>();
    (vec[0], vec[1])
}

fn parse(lines: &Vec<String>) -> Vec<Machine> {
    let mut res = Vec::new();
    let (mut a, mut b, mut prize) = ((0, 0), (0, 0), (0, 0));

    for line in lines.iter() {
        if line.starts_with("Button A") {
            a = parse_line(line);
        } else if line.starts_with("Button B") {
            b = parse_line(line);
        } else if line.starts_with("Prize") {
            prize = parse_line(line);
        } else {
            res.push(Machine {
                a: a,
                b: b,
                prize: prize,
            });
        }
    }

    res.push(Machine {
        a: a,
        b: b,
        prize: prize,
    });

    res
}

fn cost(m: &Machine, offset: i64) -> Option<i64> {
    let (px, py) = (m.prize.0 + offset, m.prize.1 + offset);
    let b1 = px * m.a.1 - py * m.a.0;
    let b2 = m.b.0 * m.a.1 - m.a.0 * m.b.1;
    let b = b1 / b2;
    let a = (px - m.b.0 * b) / m.a.0;
    if (a * m.a.0 + b * m.b.0) != px || (a * m.a.1 + b * m.b.1) != py {
        return None;
    }
    return Some(3 * a + b);
}

fn total_cost(machines: &Vec<Machine>, offset: i64) -> i64 {
    machines.iter().flat_map(|m| cost(m, offset)).sum::<i64>()
}

fn main() {
    let input = parse(&aoc::input_lines());
    println!("Part I: {}", total_cost(&input, 0));
    println!("Part II: {}", total_cost(&input, 10000000000000));
}
