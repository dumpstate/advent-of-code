use std::cmp::Ordering;
use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn parse(input: &Vec<String>) -> (HashMap<i32, HashSet<i32>>, Vec<Vec<i32>>) {
    let mut ord_rules: HashMap<i32, HashSet<i32>> = HashMap::new();
    let mut updates = Vec::new();
    let mut first_section = true;

    for line in input.iter() {
        if line == "" {
            first_section = false;
            continue;
        }

        if first_section {
            let split = line.split("|").collect::<Vec<&str>>();
            let before = split[0].trim().parse::<i32>().unwrap();
            let after = split[1].trim().parse::<i32>().unwrap();
            ord_rules.entry(before).or_insert(HashSet::new()).insert(after);
        } else {
            let split = line.split(",");
            updates.push(split.map(|x| x.parse().unwrap()).collect());
        }
    }

    (ord_rules, updates)
}

fn is_valid(rules: &HashMap<i32, HashSet<i32>>, update: &Vec<i32>) -> bool {
    for i in 0..(update.len() - 1) {
        let before = update[i];
        let after = update[i + 1];

        if !rules.contains_key(&before) {
            return false;
        }

        if !rules[&before].contains(&after) {
            return false;
        }
    }

    true
}

fn sort(rules: &HashMap<i32, HashSet<i32>>, update: &Vec<i32>) -> Vec<i32> {
    let mut sorted = update.clone();
    sorted.sort_by(|a, b| {
        if rules.contains_key(a) {
            if rules[&a].contains(b) {
                return Ordering::Less;
            }
        }

        if rules.contains_key(b) {
            if rules[&b].contains(a) {
                return Ordering::Greater;
            }
        }

        Ordering::Equal
    });
    sorted
}

fn part_1(rules: &HashMap<i32, HashSet<i32>>, updates: &Vec<Vec<i32>>) -> i32 {
    let mut res = 0;

    for update in updates.iter() {
        if is_valid(rules, update) {
            res += update[update.len() / 2];
        }
    }

    res
}

fn part_2(rules: &HashMap<i32, HashSet<i32>>, updates: &Vec<Vec<i32>>) -> i32 {
    let mut res = 0;

    for update in updates.iter() {
        if is_valid(rules, update) {
            continue
        }

        let sorted = sort(rules, update);
        res += sorted[sorted.len() / 2];
    }

    res
}

fn main() {
    let (rules, updates) = parse(&aoc::input_lines());

    println!("Part I: {}", part_1(&rules, &updates));
    println!("Part II: {}", part_2(&rules, &updates));
}
