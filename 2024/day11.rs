use std::collections::HashMap;
mod aoc;

fn is_even_len(x: i64) -> bool {
    x.to_string().len() % 2 == 0
}

fn split(x: i64) -> (i64, i64) {
    let s = x.to_string();
    (s[0..(s.len() / 2)].parse::<i64>().unwrap(), s[(s.len() / 2)..].parse::<i64>().unwrap())
}

fn blink(stones: &HashMap<i64, i64>) -> HashMap<i64, i64> {
    let mut next_stones = HashMap::new();
    for (stone, count) in stones.iter() {
        match *stone {
            0 => { *next_stones.entry(1).or_insert(0) += count; },
            x if is_even_len(x) => {
                let (a, b) = split(x);
                *next_stones.entry(a).or_insert(0) += count;
                *next_stones.entry(b).or_insert(0) += count;
            },
            x => { *next_stones.entry(x * 2024).or_insert(0) += count; }
        }
    }
    next_stones
}

fn count(stones: &HashMap<i64, i64>) -> i64 {
    stones.iter().fold(0, |acc, (_, count)| acc + count)
}

fn blink_times(stones: &HashMap<i64, i64>, times: i32) -> i64 {
    let mut res = stones.clone();
    for _ in 0..times {
        res = blink(&res);
    }
    count(&res)
}

fn main() {
    let stones = aoc::input_lines()[0].split(" ")
        .fold(HashMap::new(), |mut map, x| {
            map.insert(x.parse::<i64>().unwrap(), 1);
            map
        });
    println!("Part I: {}", blink_times(&stones, 25));
    println!("Part II: {}", blink_times(&stones, 75));
}
