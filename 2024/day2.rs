mod aoc;

fn parse(input: &Vec<String>) -> Vec<Vec<i32>> {
    input
        .iter()
        .map(|report| report
            .split_whitespace()
            .map(|level| level.parse::<i32>().unwrap())
            .collect::<Vec<i32>>())
        .collect()
}

fn is_safe(report: &Vec<i32>) -> bool {
    let mut last: Option<i32> = None;
    let mut sgn: Option<i32> = None;

    for level in report.iter() {
        if last == None {
            last = Some(*level);
        } else {
            let diff = *level - last.unwrap();
            if diff.abs() > 3 {
                return false;
            }

            if sgn == None {
                sgn = Some(diff.signum());
            } else if sgn.unwrap() != diff.signum() {
                return false;
            }

            last = Some(*level);
        }
    }

    true
}

fn versions(report: &Vec<i32>) -> Vec<Vec<i32>> {
    let mut versions = Vec::new();

    for i in 0..report.len() {
        let mut version = report.clone();
        version.remove(i);
        versions.push(version);
    }

    versions
}

fn part_1(reports: &Vec<Vec<i32>>) -> i32 {
    reports
        .iter()
        .filter(|report| is_safe(report))
        .count() as i32
}

fn part_2(reports: &Vec<Vec<i32>>) -> i32 {
    reports
        .iter()
        .filter(|report| {
            if is_safe(report) {
                return true
            }

            for version in versions(report) {
                if is_safe(&version) {
                    return true
                }
            }

            false
        })
        .count() as i32
}

fn main() {
    let reports = parse(&aoc::input_lines());

    println!("Part I: {}", part_1(&reports));
    println!("Part II: {}", part_2(&reports));
}
