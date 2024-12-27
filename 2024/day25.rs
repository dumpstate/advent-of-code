mod aoc;

fn parse(lines: &Vec<String>) -> (Vec<[usize; 5]>, Vec<[usize; 5]>) {
    let (mut keys, mut locks) = (Vec::new(), Vec::new());
    let (mut count, mut curr) = (0, [0, 0, 0, 0, 0]);
    let mut is_key = None;

    for line in lines {
        if line.is_empty() {
            is_key = None;
            count = 0;
            continue;
        }

        if is_key.is_none() {
            is_key = Some(line == ".....");
            continue;
        }

        if count == 5 {
            if is_key == Some(true) { keys.push(curr.clone()); }
            else { locks.push(curr.clone()); }
            curr = [0, 0, 0, 0, 0];
            continue;
        }

        count += 1;
        line.chars().enumerate()
            .filter(|(_, c)| *c == '#')
            .for_each(|(ix, _)| curr[ix] += 1);
    }

    (keys, locks)
}

fn main() {
    let (keys, locks) = parse(&aoc::input_lines());
    println!("Part I: {}", keys.iter()
        .fold(0, |acc, key| {
            locks.iter()
                .fold(acc, |acc, lock| {
                    if key.iter().zip(lock.iter()).all(|(k, l)| k + l <= 5) {
                        acc + 1
                    } else {
                        acc
                    }
                })
        }));
}
