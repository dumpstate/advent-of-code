use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn step(board: &Vec<Vec<char>>, pos: (i32, i32), dir: usize) -> Option<(i32, i32)> {
    let (dx, dy) = match dir {
        0 => (1, 0),
        1 => (0, 1),
        2 => (-1, 0),
        3 => (0, -1),
        _ => panic!("invalid direction"),
    };
    let (nx, ny) = (pos.0 + dx, pos.1 + dy);
    if board[ny as usize][nx as usize] == '#' {
        return None;
    }
    Some((nx, ny))
}

fn best_score(board: &Vec<Vec<char>>, visited: &HashMap<((i32, i32), usize), i64>) -> i64 {
    let end = aoc::find(&board, 'E');
    visited
        .iter()
        .filter_map(
            |((pos, _), score)| {
                if *pos == end {
                    Some(*score)
                } else {
                    None
                }
            },
        )
        .min()
        .unwrap()
}

fn count_tiles(paths: &Vec<(Vec<(i32, i32)>, i64)>, min_score: i64) -> usize {
    paths
        .iter()
        .filter(|(_, score)| *score == min_score)
        .flat_map(|(path, _)| path.iter())
        .collect::<HashSet<_>>()
        .len()
}

fn find_paths(board: &Vec<Vec<char>>) -> (i64, usize) {
    let mut q = Vec::new();
    let mut visited = HashMap::new();
    let mut paths = Vec::new();
    let start = aoc::find(board, 'S');

    q.push((start, 0, 0, vec![start]));

    while let Some((pos, dir, score, path)) = q.pop() {
        if board[pos.1 as usize][pos.0 as usize] == 'E' {
            paths.push((path.clone(), score));
        }

        match visited.get(&(pos, dir)) {
            Some(prev_score) if *prev_score < score => continue,
            Some(_) | None => visited.insert((pos, dir), score),
        };
        if let Some(next_pos) = step(board, pos, dir) {
            q.push((
                next_pos,
                dir,
                score + 1,
                [path.clone(), vec![next_pos]].concat(),
            ));
        }
        let dl = (dir as i32 - 1).rem_euclid(4) as usize;
        if let Some(_) = step(board, pos, dl) {
            q.push((pos, dl, score + 1000, path.clone()));
        }
        let dr = (dir + 1).rem_euclid(4) as usize;
        if let Some(_) = step(board, pos, dr) {
            q.push((pos, dr, score + 1000, path.clone()));
        }
    }

    let min_score = best_score(board, &visited);
    (min_score, count_tiles(&paths, min_score))
}

fn main() {
    let board = aoc::input_board();
    let (part_1, part_2) = find_paths(&board);
    println!("Part I: {}", part_1);
    println!("Part II: {}", part_2);
}
