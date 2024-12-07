mod aoc;

type Op = fn(i64, i64) -> i64;

fn parse(input: &Vec<String>) -> Vec<(i64, Vec<i64>)> {
    let mut res = Vec::new();

    for line in input.iter() {
        let split = line.split(": ").collect::<Vec<&str>>();
        let total = split[0].parse::<i64>().unwrap();
        let nums = split[1].split(" ").map(|x| x.parse().unwrap()).collect();
        res.push((total, nums));
    }

    res
}

fn can_be_equal(target: i64, acc: i64, nums: &Vec<i64>, ops: &Vec<Op>) -> bool {
    if nums.len() == 0 {
        return acc == target;
    }

    if acc > target {
        return false;
    }

    let head = nums[0];
    let tail = &nums[1..].to_vec();

    for op in ops.iter() {
        if can_be_equal(target, op(acc, head), tail, ops) {
            return true;
        }
    }

    false
}

fn sum_valid(input: &Vec<(i64, Vec<i64>)>, ops: &Vec<Op>) -> i64 {
    let mut total: i64 = 0;

    for (target, nums) in input.iter() {
        if can_be_equal(*target, nums[0], &nums[1..].to_vec(), ops) {
            total += target;
        }
    }

    total
}

fn part_1(input: &Vec<(i64, Vec<i64>)>) -> i64 {
    sum_valid(input, &vec![
        |a, b| a + b,
        |a, b| a * b,
    ])
}

fn part_2(input: &Vec<(i64, Vec<i64>)>) -> i64 {
    sum_valid(input, &vec![
        |a, b| a + b,
        |a, b| a * b,
        |a, b| format!("{}{}", a, b).parse::<i64>().unwrap(),
    ])
}

fn main() {
    let input = parse(&aoc::input_lines());

    println!("Part 1: {}", part_1(&input));
    println!("Part 2: {}", part_2(&input));
}
