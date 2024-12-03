mod aoc;

fn is_valid(s: &str) -> bool {
    s.len() <= 3 && s.len() > 0 && s.chars().nth(0) != Some('0')
}

fn mul(s: &str) -> Option<i32> {
    let by_comma = s.split(",").collect::<Vec<&str>>();
    if by_comma.len() <= 1 {
        return None;
    }

    let lstr = by_comma[0];
    if !is_valid(lstr) {
        return None;
    }

    let closing = by_comma[1].split(")").collect::<Vec<&str>>();
    if closing.len() <= 1 {
        return None;
    }
    let rstr = closing[0];
    if !is_valid(rstr) {
        return None;
    }

    let l = lstr.parse::<i32>();
    let r = rstr.parse::<i32>();
    if l.is_err() || r.is_err() {
        return None;
    }

    Some(l.unwrap() * r.unwrap())
}

fn part_1(lines: &Vec<String>) -> i32 {
    let mut total = 0;

    for line in lines {
        for segment in line.split("mul(") {
            let res = mul(segment);
            if res.is_none() {
                continue;
            }
            total += res.unwrap();
        }
    }

    total
}

fn part_2(lines: &Vec<String>) -> i32 {
    let mut total = 0;
    let mut rem = lines.join("\n");
    let mut is_enabled = true;

    while rem.len() > 0 {
        if is_enabled {
            let mul_ix = rem.find("mul(");
            if mul_ix.is_none() {
                break;
            }

            let dont_ix = rem.find("don't()");
            if dont_ix.is_some() && dont_ix.unwrap() < mul_ix.unwrap() {
                rem = rem[(dont_ix.unwrap() + 7)..].to_string();
                is_enabled = false;
                continue;
            }

            rem = rem[(mul_ix.unwrap() + 4)..].to_string();
            let next = mul(&rem);
            if next.is_some() {
                total += next.unwrap();
            }
        } else {
            let do_ix = rem.find("do()");
            if do_ix.is_none() {
                break;
            }
            rem = rem[(do_ix.unwrap() + 4)..].to_string();
            is_enabled = true;
        }
    }

    total
}

fn main() {
    let input = aoc::input_lines();

    println!("Part I: {}", part_1(&input));
    println!("Part II: {}", part_2(&input));
}
