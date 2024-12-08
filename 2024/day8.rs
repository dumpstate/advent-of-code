use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn part_1(board: &Vec<Vec<char>>) -> usize {
    let mut antinodes = HashSet::new();
    let mut antennas: HashMap<char, HashSet<(usize, usize)>> = HashMap::new();

    for y in 0..board.len() {
        for x in 0..board.len() {
            if board[y][x] == '.' {
                continue;
            }

            let found = antennas.entry(board[y][x]).or_insert(HashSet::new());
            for (ax, ay) in found.iter() {
                let (dx, dy) = (x as i32 - *ax as i32, y as i32 - *ay as i32);

                let (a1x, a1y) = (*ax as i32 - dx, *ay as i32 - dy);
                if aoc::is_on_board(board, a1x, a1y) {
                    antinodes.insert((a1x, a1y));
                }

                let (a2x, a2y) = (x as i32 + dx, y as i32 + dy);
                if aoc::is_on_board(board, a2x, a2y) {
                    antinodes.insert((a2x, a2y));
                }
            }
            found.insert((x, y));
        }
    }

    antinodes.len()
}

fn part_2(board: &Vec<Vec<char>>) -> usize {
    let mut antinodes = HashSet::new();
    let mut antennas: HashMap<char, HashSet<(usize, usize)>> = HashMap::new();

    for y in 0..board.len() {
        for x in 0..board.len() {
            if board[y][x] == '.' {
                continue;
            }

            let found = antennas.entry(board[y][x]).or_insert(HashSet::new());
            for (ax, ay) in found.iter() {
                let (dx, dy) = (x as i32 - *ax as i32, y as i32 - *ay as i32);

                let (mut a1x, mut a1y) = (*ax as i32 - dx, *ay as i32 - dy);
                while aoc::is_on_board(board, a1x, a1y) {
                    antinodes.insert((a1x, a1y));
                    a1x -= dx;
                    a1y -= dy;
                }

                let (mut a2x, mut a2y) = (x as i32 + dx, y as i32 + dy);
                while aoc::is_on_board(board, a2x, a2y) {
                    antinodes.insert((a2x, a2y));
                    a2x += dx;
                    a2y += dy;
                }
            }
            antinodes.insert((x as i32, y as i32));
            found.insert((x, y));
        }
    }

    antinodes.len()
}

fn main() {
    let board = aoc::input_board();

    println!("Part I: {}", part_1(&board));
    println!("Part II: {}", part_2(&board));
}
