use std::collections::HashMap;
use std::env;
use std::fs::read_to_string;
use std::path::PathBuf;

pub struct Args {
    pub path: PathBuf
}

pub fn parse_args() -> Args {
    let input = env::args().nth(1).expect("no path given");
    Args { path: PathBuf::from(input) }
}

pub fn read_lines(fname: &PathBuf) -> Vec<String> {
    read_to_string(fname.display().to_string())
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

pub fn input_lines() -> Vec<String> {
    let args = parse_args();
    read_lines(&args.path)
}

pub fn counter(vec: &Vec<i32>) -> HashMap<i32, i32> {
    vec.iter()
        .fold(HashMap::new(), |mut map, i| {
            let count = map.entry(*i).or_insert(0);
            *count += 1;
            map
        })
}
