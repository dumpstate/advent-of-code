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

pub fn input_board() -> Vec<Vec<char>> {
    input_lines()
        .iter()
        .map(|line| line.chars().collect())
        .collect()
}

pub fn counter(vec: &Vec<i32>) -> HashMap<i32, i32> {
    vec.iter()
        .fold(HashMap::new(), |mut map, i| {
            let count = map.entry(*i).or_insert(0);
            *count += 1;
            map
        })
}

pub fn find(board: &Vec<Vec<char>>, c: char) -> (i32, i32) {
    for y in 0..board.len() {
        for x in 0..board[0].len() {
            if board[y][x] == c {
                return (x as i32, y as i32);
            }
        }
    }
    panic!("character not found");
}

pub fn is_on_board<T>(board: &Vec<Vec<T>>, x: i32, y: i32) -> bool {
    x >= 0 && y >= 0 && y < board.len() as i32 && x < board[0].len() as i32
}

pub fn split(s: &str, delim: &str) -> Vec<String> {
    s.split(delim).map(String::from).collect()
}
