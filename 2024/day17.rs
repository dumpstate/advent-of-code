mod aoc;

fn parse(lines: &Vec<String>) -> ((i64, i64, i64), Vec<i64>) {
    let reg_a = aoc::split(&lines[0], ": ")[1].parse::<i64>().unwrap();
    let reg_b = aoc::split(&lines[1], ": ")[1].parse::<i64>().unwrap();
    let reg_c = aoc::split(&lines[2], ": ")[1].parse::<i64>().unwrap();
    let program = aoc::split(&lines[4], ": ")[1].split(",").map(|x| x.parse::<i64>().unwrap()).collect();
    ((reg_a, reg_b, reg_c), program)
}

fn join(pr: &Vec<i64>) -> String {
    pr.iter().map(|x| x.to_string()).collect::<Vec<String>>().join(",")
}

fn evaluate(reg: (i64, i64, i64), program: &Vec<i64>) -> Vec<i64> {
    let (mut a, mut b, mut c) = reg;
    let mut ip = 0;
    let mut out = Vec::new();

    while ip < program.len() {
        let (instr, lit) = (program[ip], program[ip + 1]);
        let com = match lit {
            4 => a,
            5 => b,
            6 => c,
            x => x,
        };

        match instr {
            0 => { a >>= com; },
            1 => { b ^= lit; },
            2 => { b = com & 7; },
            3 => {
                if a != 0 {
                    ip = lit as usize;
                    continue;
                }
            },
            4 => { b ^= c; },
            5 => { out.push(com & 7); },
            6 => { b = a >> com; },
            7 => { c = a >> com; },
            _ => panic!("invalid instruction"),
        };

        ip += 2;
    }

    out
}

fn solve(program: &Vec<i64>, pos: i32, curr: i64) -> Option<i64> {
    if pos < 0 {
        return Some(curr);
    }

    for off in 0..8 {
        let (mut a, mut b, mut c) = (curr << 3 | off, 0, 0);
        let mut ip = 0;
        let mut out = None;
        while ip < program.len() {
            let (instr, lit) = (program[ip], program[ip + 1]);
            let com = match lit {
                4 => a,
                5 => b,
                6 => c,
                x => x,
            };

            match instr {
                0 => { a >>= com; },
                1 => { b ^= lit; },
                2 => { b = com & 7; },
                3 => {
                    if a != 0 {
                        ip = lit as usize;
                        continue;
                    }
                },
                4 => { b ^= c; },
                5 => {
                    out = Some(com & 7);
                    break;
                },
                6 => { b = a >> com; },
                7 => { c = a >> com; },
                _ => panic!("invalid instruction"),
            };

            ip += 2;
        }

        if out == Some(program[pos as usize]) {
            if let Some(next) = solve(program, pos - 1, curr << 3 | off) {
                return Some(next);
            }
        }
    }
    None
}

fn main() {
    let (reg, program) = parse(&aoc::input_lines());
    println!("Part I: {}", join(&evaluate(reg, &program)));
    println!("Part II: {}", solve(&program, program.len() as i32 - 1, 0).unwrap());
}
