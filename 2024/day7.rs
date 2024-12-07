mod aoc;

fn parse(input: &Vec<String>) -> Vec<(i64, Vec<i64>)> {
    input.iter()
        .map(|line| {
            let split = line.split(": ").collect::<Vec<&str>>();
            let total = split[0].parse::<i64>().unwrap();
            let nums = split[1].split(" ").map(|x| x.parse().unwrap()).collect();
            (total, nums)
        })
        .collect()
}

fn can_be_equal(target: i64, acc: i64, nums: &[i64], ops: &[fn(i64, i64) -> i64]) -> bool {
    if nums.len() == 0 { return acc == target; }
    if acc > target { return false; }

    if let Some((head, tail)) = nums.split_first() {
        if ops.iter().find(|op| can_be_equal(target, op(acc, *head), tail, ops)).is_some() {
            return true;
        }
    }

    false
}

fn sum(input: &Vec<(i64, Vec<i64>)>, ops: &[fn(i64, i64) -> i64]) -> i64 {
    input.iter()
        .filter(|(target, nums)| can_be_equal(*target, nums[0], &nums[1..], ops))
        .map(|(target, _)| target)
        .sum()
}

fn main() {
    let input = parse(&aoc::input_lines());

    println!("Part 1: {}", sum(&input, &[|a, b| a + b, |a, b| a * b]));
    println!("Part 2: {}", sum(&input, &[
        |a, b| a + b,
        |a, b| a * b,
        |a, b| format!("{}{}", a, b).parse::<i64>().unwrap(),
    ]));
}
