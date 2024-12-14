use std::collections::HashMap;
mod aoc;

type Robot = ((i64, i64), (i64, i64));

const SIZE_X: i64 = 101;
const SIZE_Y: i64 = 103;

fn parse(lines: &Vec<String>) -> Vec<Robot> {
    lines
        .iter()
        .map(|line| {
            let split = line.split(" ").collect::<Vec<&str>>();
            let p = split[0][2..].split(",").collect::<Vec<&str>>();
            let v = split[1][2..].split(",").collect::<Vec<&str>>();
            (
                (p[0].parse::<i64>().unwrap(), p[1].parse::<i64>().unwrap()),
                (v[0].parse::<i64>().unwrap(), v[1].parse::<i64>().unwrap()),
            )
        })
        .collect()
}

fn pos(r: Robot, sec: i64) -> (i64, i64) {
    let ((px, py), (vx, vy)) = r;
    let nx = (px + vx * sec).rem_euclid(SIZE_X);
    let ny = (py + vy * sec).rem_euclid(SIZE_Y);
    (nx, ny)
}

fn quadrant(pos: (i64, i64)) -> Option<usize> {
    match pos {
        (x, y) if x < SIZE_X / 2 && y < SIZE_Y / 2 => Some(0),
        (x, y) if x > SIZE_X / 2 && y < SIZE_Y / 2 => Some(1),
        (x, y) if x < SIZE_X / 2 && y > SIZE_Y / 2 => Some(2),
        (x, y) if x > SIZE_X / 2 && y > SIZE_Y / 2 => Some(3),
        _ => None,
    }
}

fn max_tile(input: &Vec<Robot>, time: i64) -> i64 {
    input
        .iter()
        .map(|r| pos(*r, time))
        .fold(HashMap::new(), |mut acc, p| {
            *acc.entry(p).or_insert(0) += 1;
            acc
        })
        .values()
        .max()
        .unwrap()
        .clone()
}

fn part_1(input: &Vec<Robot>) -> i64 {
    input
        .iter()
        .map(|r| pos(*r, 100))
        .flat_map(|p| quadrant(p))
        .fold(HashMap::new(), |mut acc, q| {
            *acc.entry(q).or_insert(0) += 1;
            acc
        })
        .into_iter()
        .fold(1, |acc, (_, count)| acc * count)
}

fn part_2(input: &Vec<Robot>) -> i64 {
    for i in 0..1000000 {
        if max_tile(input, i) == 1 {
            return i;
        }
    }
    panic!("No solution found");
}

fn main() {
    let input = parse(&aoc::input_lines());
    println!("Part I: {}", part_1(&input));
    println!("Part II: {}", part_2(&input));
}
