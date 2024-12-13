mod aoc;

#[derive(Debug)]
struct Machine {
    button_a: (i64, i64),
    button_b: (i64, i64),
    prize: (i64, i64),
}

fn parse(lines: &Vec<String>) -> Vec<Machine> {
    let mut res = Vec::new();
    let mut button_a = (0, 0);
    let mut button_b = (0, 0);
    let mut prize = (0, 0);

    for line in lines.iter() {
        match line {
            l if l.starts_with("Button A") => {
                let coords = l.split(": ").collect::<Vec<&str>>()[1]
                    .split(", ")
                    .map(|x| x[2..].parse::<i64>().unwrap())
                    .collect::<Vec<i64>>();
                button_a = (coords[0], coords[1]);
            }
            l if l.starts_with("Button B") => {
                let coords = l.split(": ").collect::<Vec<&str>>()[1]
                    .split(", ")
                    .map(|x| x[2..].parse::<i64>().unwrap())
                    .collect::<Vec<i64>>();
                button_b = (coords[0], coords[1]);
            }
            l if l.starts_with("Prize") => {
                let coords = l.split(": ").collect::<Vec<&str>>()[1]
                    .split(", ")
                    .map(|x| x[2..].parse::<i64>().unwrap())
                    .collect::<Vec<i64>>();
                prize = (coords[0], coords[1]);
            }
            l if l.len() == 0 => {
                res.push(Machine {
                    button_a: button_a,
                    button_b: button_b,
                    prize: prize,
                });
            }
            _ => panic!("Unexpected input"),
        }
    }

    res.push(Machine {
        button_a: button_a,
        button_b: button_b,
        prize: prize,
    });

    res
}

fn cost(m: &Machine, offset: i64) -> Option<i64> {
    let (px, py) = (m.prize.0 + offset, m.prize.1 + offset);
    let b1 = px * m.button_a.1 - py * m.button_a.0;
    let b2 = m.button_b.0 * m.button_a.1 - m.button_a.0 * m.button_b.1;
    let b = b1 / b2;
    let a = (px - m.button_b.0 * b) / m.button_a.0;
    if (a * m.button_a.0 + b * m.button_b.0) != px || (a * m.button_a.1 + b * m.button_b.1) != py {
        return None;
    }
    return Some(3 * a + b);
}

fn main() {
    let input = parse(&aoc::input_lines());
    println!(
        "Part I: {}",
        input.iter().flat_map(|m| cost(m, 0)).sum::<i64>()
    );
    println!(
        "Part II: {}",
        input
            .iter()
            .flat_map(|m| cost(m, 10000000000000))
            .sum::<i64>()
    )
}
