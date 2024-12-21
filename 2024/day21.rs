use std::collections::HashMap;
mod aoc;

fn index(pad: &Vec<[char; 3]>, c: char) -> (usize, usize) {
    for (y, row) in pad.iter().enumerate() {
        for (x, &cell) in row.iter().enumerate() {
            if cell == c {
                return (x, y);
            }
        }
    }
    panic!("not found")
}

fn has_key(pad: &Vec<[char; 3]>, (x, y): (usize, usize)) -> bool {
    if x >= pad[0].len() || y >= pad.len() {
        return false;
    }
    pad[y][x] != '?'
}

fn step((x, y): (usize, usize), (tx, ty): (usize, usize), pad: &Vec<[char; 3]>) -> (Vec<char>, (usize, usize)) {
    let (dx, dy) = (tx as i32 - x as i32, ty as i32 - y as i32);
    let mut h = Vec::new();
    for _ in 0..dx { h.push('>'); }
    for _ in 0..-dx { h.push('<'); }
    let mut v = Vec::new();
    for _ in 0..dy { v.push('v'); }
    for _ in 0..-dy { v.push('^'); }

    if dx > 0 && has_key(pad, (x, ty)) {
        return (aoc::concat(v, h, vec!['A']), (tx, ty));
    }
    if has_key(pad, (tx, y)) {
        return (aoc::concat(h, v, vec!['A']), (tx, ty));
    }
    (aoc::concat(v, h, vec!['A']), (tx, ty))
}

fn enter_code(code: &String, start: (usize, usize), pad: &Vec<[char; 3]>) -> Vec<Vec<char>> {
    let mut pos = start;
    let mut moves = Vec::new();
    for c in code.chars() {
        let target = index(pad, c);
        let (next_moves, next_arm) = step(pos, target, pad);
        moves.push(next_moves);
        pos = next_arm;
    }
    moves
}

fn chain(codes: &Vec<String>, dpads: usize) -> i64 {
    let npad: Vec<[char; 3]> = vec![
        ['7', '8', '9'],
        ['4', '5', '6'],
        ['1', '2', '3'],
        ['?', '0', 'A'],
    ];

    let dpad: Vec<[char; 3]> = vec![
        ['?', '^', 'A'],
        ['<', 'v', '>'],
    ];

    let mut res = 0;
    for code in codes {
        let curr = aoc::join(&enter_code(&code, (2, 3), &npad));
        let mut sub_codes: HashMap<String, i64> = HashMap::new();
        sub_codes.insert(curr, 1);
        for _ in 0..dpads {
            let mut counts: HashMap<String, i64> = HashMap::new();
            for (sub_code, q) in &sub_codes.clone() {
                let mut next_codes: HashMap<String, i64> = HashMap::new();
                for sub in enter_code(&sub_code, (2, 0), &dpad) {
                    *next_codes.entry(sub.iter().collect()).or_insert(0) += 1;
                }
                for (sub, _) in next_codes.clone() {
                    *next_codes.entry(sub).or_insert(0) *= q;
                }
                for (sub, c) in &next_codes {
                    *counts.entry(sub.clone()).or_insert(0) += c;
                }
            }
            sub_codes = counts.clone();
        }
        let mut len = 0;
        for (sub, q) in sub_codes {
            len += sub.len() as i64 * q;
        }
        res += len * code[..code.len() - 1].parse::<i64>().unwrap();
    }
    res
}

fn main() {
    let codes = aoc::input_lines();
    println!("Part I: {}", chain(&codes, 2));
    println!("Part II: {}", chain(&codes, 25));
}
