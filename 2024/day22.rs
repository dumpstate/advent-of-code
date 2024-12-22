use std::collections::HashMap;
mod aoc;

fn secret(n: i64, pos: usize) -> i64 {
    if pos == 0 { return n; }
    let mut next = ((n * 64) ^ n) % 16777216;
    next = ((next / 32) ^ next) % 16777216;
    next = ((next * 2048) ^ next) % 16777216;
    secret(next, pos - 1)
}

fn total(prices: &Vec<HashMap<String, i64>>, w: &String) -> i64 {
    prices.iter().map(|ps| ps.get(w).unwrap_or(&0)).sum()
}

fn part_1(ns: &Vec<i64>) -> i64 {
    ns.iter().map(|n| secret(*n, 2000)).sum()
}

fn part_2(ns: &Vec<i64>) -> i64 {
    let prices = ns.iter()
        .map(|n| {
            (0..2000)
                .fold(vec![*n], |mut secrets, _| {
                    secrets.push(secret(*secrets.last().unwrap(), 1));
                    secrets
                }).iter()
                .map(|s| s % 10).collect::<Vec<i64>>()
        })
        .map(|secrets| {
            let diff = secrets
                .windows(2)
                .map(|w| w[1] - w[0])
                .collect::<Vec<i64>>();
            secrets[1..].iter()
                .zip(diff.iter())
                .map(|(s, d)| (*s, *d))
                .collect::<Vec<(i64, i64)>>()
                .windows(4)
                .fold(HashMap::new(), |mut window_to_prices, w| {
                    let key = w.iter().map(|(_, d)| d.to_string()).collect::<Vec<String>>().join("");
                    if !window_to_prices.contains_key(&key) {
                        window_to_prices.insert(key, w.iter().map(|(s, _)| *s).last().unwrap());
                    }
                    window_to_prices
                })
        }).collect::<Vec<HashMap<String, i64>>>();

    *prices.iter()
        .fold(HashMap::new(), |cache, ps| {
            ps.iter()
                .fold(cache, |mut cache, (w, _)| {
                    if !cache.contains_key(w) {
                        cache.insert(w.clone(), total(&prices, w));
                    }
                    cache
                })
        })
        .values().max().unwrap()
}

fn main() {
    let input = aoc::input_lines().into_iter()
        .map(|line| line.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();
    println!("Part I: {}", part_1(&input));
    println!("Part II: {}", part_2(&input));
}
