use std::collections::HashMap;
mod aoc;

fn count(cache: &mut HashMap<String, i64>, stock: &Vec<String>, design: &String) -> i64 {
    if cache.contains_key(design) {
        return cache[design];
    }

    if design.is_empty() {
        return 1;
    }

    let total = stock.iter()
        .filter(|pat| design.starts_with(*pat))
        .map(|pat| count(cache, stock, &design[pat.len()..].to_string()))
        .sum::<i64>();
    cache.insert(design.clone(), total);
    total
}

fn main() {
    let input = aoc::input_lines();
    let (stock, designs) = (aoc::split(&input[0], ", "), &input[2..]);
    let mut cache = HashMap::new();
    let counts = designs.iter()
        .map(|d| count(&mut cache, &stock, d))
        .collect::<Vec<i64>>();
    println!("Part I: {}", counts.clone().into_iter().filter(|c| *c > 0).count());
    println!("Part II: {}", counts.clone().into_iter().sum::<i64>());
}
