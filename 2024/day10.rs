use std::collections::VecDeque;
use std::collections::HashSet;
mod aoc;

fn count_paths(board: &Vec<Vec<usize>>, (x, y): (usize, usize)) -> (i32, i32) {
    let (mut targets, mut total_paths) = (HashSet::new(), 0);
    let mut q = VecDeque::new();
    q.push_back((x, y));

    while !q.is_empty() {
        let (x, y) = q.pop_front().unwrap();
        let curr = board[y][x];

        if curr == 9 {
            total_paths += 1;
            targets.insert((x, y));
            continue;
        }

        for (dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = (x as i32 + dx, y as i32 + dy);
            if aoc::is_on_board(board, nx, ny) && board[ny as usize][nx as usize] == curr + 1 {
                q.push_back((nx as usize, ny as usize));
            }
        }
    }

    (targets.len() as i32, total_paths)
}

fn main() {
    let board: Vec<Vec<usize>> = aoc::input_board()
        .iter()
        .map(|line| line.iter().map(|c| c.to_digit(10).unwrap() as usize).collect())
        .collect();
    let (target_count, path_count) = board.iter().enumerate()
        .fold((0, 0), |(tc, pc), (y, row)| {
            row.iter().enumerate().fold((tc, pc), |(tc, pc), (x, &cell)| {
                if cell == 0 {
                    let (targets, paths) = count_paths(&board, (x, y));
                    return (tc + targets, pc + paths);
                }
                (tc, pc)
            })
        });

    println!("Part I: {}", target_count);
    println!("Part II: {}", path_count);
}
