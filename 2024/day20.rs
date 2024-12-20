use std::collections::VecDeque;
mod aoc;

fn count_paths(board: &Vec<Vec<char>>, tunnel_len: i64) -> usize {
    let start = aoc::find(&board, 'S');
    let (sx, sy) = (start.0 as usize, start.1 as usize);
    let mut q = VecDeque::new();
    let mut dist = vec![vec![std::i64::MAX; board[0].len()]; board.len()];
    let mut best_path = None;

    dist[sy][sx] = 0;
    q.push_back(((sx, sy), 0, vec![(sx, sy)]));

    while let Some(((x, y), c, path)) = q.pop_front() {
        if board[y][x] == 'E' {
            best_path = Some(path);
            break;
        }

        for (dx, dy) in [(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = ((x as i32 + dx) as usize, (y as i32 + dy) as usize);
            if !aoc::is_on_board(&board, nx as i32, ny as i32) || board[ny][nx] == '#' {
                continue;
            }
            if dist[ny][nx] > c + 1 {
                dist[ny][nx] = c + 1;
                q.push_back(((nx, ny), c + 1, path.iter().cloned().chain(std::iter::once((nx, ny))).collect()));
            }
        }
    }

    let path = best_path.clone().unwrap();
    path.iter().flat_map(|a| path.iter().map(move |b| (a, b)))
        .filter(|(a, b)| a != b)
        .filter(|((ax, ay), (bx, by))| {
            let tunnel_dist = (*ax as i64 - *bx as i64).abs() + (*ay as i64 - *by as i64).abs();
            let saved_cost = dist[*by][*bx] - dist[*ay][*ax] - tunnel_dist;
            tunnel_dist <= tunnel_len && saved_cost >= 100
        })
        .count()
}


fn main() {
    let board = aoc::input_board();
    println!("Part I: {}", count_paths(&board, 2));
    println!("Part II: {}", count_paths(&board, 20));
}
