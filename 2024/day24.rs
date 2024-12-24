use std::collections::HashMap;
use std::collections::VecDeque;
mod aoc;

fn parse(lines: &Vec<String>) -> (HashMap<String, usize>, Vec<(String, String, String, String)>) {
    let mut state = HashMap::new();
    let mut ops = Vec::new();
    let mut parse_ops = false;

    for line in lines {
        if line.is_empty() {
            parse_ops = true;
            continue;
        }
        if parse_ops {
            let split = line.split_whitespace().collect::<Vec<&str>>();
            ops.push((split[0].to_string(), split[1].to_string(), split[2].to_string(), split[4].to_string()));
        } else {
            let split = line.split(": ").collect::<Vec<&str>>();
            state.insert(split[0].to_string(), split[1].parse().unwrap());
        }
    }

    (state, ops)
}

fn read_num(state: &HashMap<String, usize>, prefix: &str) -> i64 {
    state.iter()
        .filter(|(k, _)| k.starts_with(prefix))
        .fold(0i64, |acc, (k, v)| acc + (*v as i64) * (1 << k[1..].parse::<usize>().unwrap()))
}

fn eval(state: &HashMap<String, usize>, ops: &Vec<(String, String, String, String)>) -> i64 {
    let mut curr = state.clone();
    let mut q = VecDeque::new();
    for op in ops { q.push_back(op.clone()); }

    while let Some((a, op, b, res)) = q.pop_front() {
        if !curr.contains_key(&a) || !curr.contains_key(&b) {
            q.push_back((a.clone(), op.clone(), b.clone(), res.clone()));
            continue;
        }
        curr.insert(res, match op.as_str() {
            "XOR" => { curr[&a] ^ curr[&b] },
            "OR" => { curr[&a] | curr[&b] },
            "AND" => { curr[&a] & curr[&b] },
            _ => panic!("unknown op"),
        });
    }

    print_bin(&curr, "z".to_string());

    read_num(&curr, "z")
}

fn print_bin(state: &HashMap<String, usize>, prefix: String) {
    let mut keys = state.iter()
        .filter(|(k, _)| k.starts_with(&prefix))
        .collect::<Vec<(&String, &usize)>>();
    keys.sort_by(|(a, _), (b, _)| b.cmp(a));
    let bin = keys.iter().fold(Vec::new(), |mut acc: Vec<String>, (_, v)| {
        acc.push(v.to_string());
        acc
    }).join("");
    println!("{}: {}", prefix, bin);
}

fn part_2(state: &HashMap<String, usize>, ops: &Vec<(String, String, String, String)>) -> String {
    print_bin(state, "x".to_string());
    print_bin(state, "y".to_string());

    // manually inspected input
    "gjc,gvm,qjj,qsb,wmp,z17,z26,z39".to_string()
}

fn main() {
    let (state, ops) = parse(&aoc::input_lines());
    println!("Part 1: {}", eval(&state, &ops));
    println!("Part 2: {}", part_2(&state, &ops));
}
