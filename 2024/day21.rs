mod aoc;

fn concat(mut v: Vec<char>, mut h: Vec<char>, mut a: Vec<char>) -> Vec<char> {
    let mut res = Vec::new();
    res.append(&mut v);
    res.append(&mut h);
    res.append(&mut a);
    res
}

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
        return (concat(v, h, vec!['A']), (tx, ty));
    }
    if has_key(pad, (tx, y)) {
        return (concat(h, v, vec!['A']), (tx, ty));
    }
    (concat(v, h, vec!['A']), (tx, ty))
}

fn enter_code(code: &String, start: (usize, usize), pad: &Vec<[char; 3]>) -> String {
    let mut pos = start;
    let mut moves = Vec::new();
    for c in code.chars() {
        let target = index(pad, c);
        let (next_moves, next_arm) = step(pos, target, pad);
        for m in next_moves {
            moves.push(m);
        }
        pos = next_arm;
    }
    moves.iter().collect()
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
        let mut curr = enter_code(&code, (2, 3), &npad);
        let num_part = code[..code.len() - 1].parse::<i64>().unwrap();
        for _ in 0..dpads {
            curr = enter_code(&curr, (2, 0), &dpad);
        }
        res += curr.len() as i64 * num_part;
    }
    res
}

fn main() {
    let codes = aoc::input_lines();
    println!("Part I: {}", chain(&codes, 2));
    // println!("Part II: {}", chain(&codes, 25));
}
