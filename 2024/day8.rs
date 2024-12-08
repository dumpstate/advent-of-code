use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn count_antinodes(board: &Vec<Vec<char>>, count_all: bool) -> usize {
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
                    if !count_all {
                        break;
                    }
                    a1x -= dx;
                    a1y -= dy;
                }

                let (mut a2x, mut a2y) = (x as i32 + dx, y as i32 + dy);
                while aoc::is_on_board(board, a2x, a2y) {
                    antinodes.insert((a2x, a2y));
                    if !count_all {
                        break;
                    }
                    a2x += dx;
                    a2y += dy;
                }
            }
            if count_all {
                antinodes.insert((x as i32, y as i32));
            }
            found.insert((x, y));
        }
    }

    antinodes.len()
}

fn main() {
    let board = aoc::input_board();

    println!("Part I: {}", count_antinodes(&board, false));
    println!("Part II: {}", count_antinodes(&board, true));
}
