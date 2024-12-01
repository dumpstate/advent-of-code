mod aoc;

fn parse(input: &Vec<String>) -> (Vec<i32>, Vec<i32>) {
    input
        .iter()
        .map(|line| line
            .split_whitespace()
            .map(|word| word.parse::<i32>().unwrap())
            .collect::<Vec<i32>>())
        .fold((Vec::new(), Vec::new()), |(mut l, mut r), v| {
            l.push(v[0]);
            r.push(v[1]);
            (l, r)
        })
}

fn part_1(mut left: Vec<i32>, mut right: Vec<i32>) -> i32 {
    left.sort();
    right.sort();

    left.iter().zip(right.iter())
        .fold(0, |acc, (l, r)| acc + (r - l).abs())
}

fn part_2(left: Vec<i32>, right: Vec<i32>) -> i32 {
    let cntr = aoc::counter(&right);

    left.iter()
        .map(|i| cntr.get(i).unwrap_or(&0) * i)
        .fold(0, |acc, i| acc + i)
}

fn main() {
    let (left, right) = parse(&aoc::input_lines());

    println!("Part I: {}", part_1(left.clone(), right.clone()));
    println!("Part II: {}", part_2(left.clone(), right.clone()));
}
