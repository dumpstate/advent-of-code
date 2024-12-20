use std::collections::BinaryHeap;
use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn dijkstra(board: &Vec<Vec<char>>, start: (usize, usize)) -> Vec<Vec<i64>> {
    let mut dist = vec![vec![std::i64::MAX; board[0].len()]; board.len()];
    let mut heap = BinaryHeap::new();

    dist[start.1][start.0] = 0;
    heap.push(aoc::State {
        pos: start,
        cost: 0,
    });

    while let Some(aoc::State { cost, pos }) = heap.pop() {
        if dist[pos.1][pos.0] < cost {
            continue;
        }

        for (dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = ((pos.0 as i32 + dx) as usize, (pos.1 as i32 + dy) as usize);
            if !aoc::is_on_board(board, nx as i32, ny as i32) || board[ny][nx] == '#' {
                continue;
            }
            if dist[ny][nx] > cost + 1 {
                heap.push(aoc::State {
                    pos: (nx, ny),
                    cost: cost + 1,
                });
                dist[ny][nx] = cost + 1;
            }
        }
    }

    dist
}

fn part_1(board: &mut Vec<Vec<char>>) -> usize {
    let (sx, sy) = aoc::find(board, 'S');
    let start = (sx as usize, sy as usize);
    let distances = dijkstra(board, start);

    let mut cheats: HashMap<i64, HashSet<((usize, usize), (usize, usize))>> = HashMap::new();
    let mut q = vec![start];

    while let Some((x, y)) = q.pop() {
        if board[y][x] == 'E' {
            break;
        }

        for (dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = ((x as i32 + dx) as usize, (y as i32 + dy) as usize);
            if !aoc::is_on_board(board, nx as i32, ny as i32) {
                continue;
            }
            if board[ny][nx] == '#' {
                for (d2x, d2y) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
                    let (nx2, ny2) = ((nx as i32 + d2x) as usize, (ny as i32 + d2y) as usize);
                    if !aoc::is_on_board(board, nx2 as i32, ny2 as i32) ||
                        (nx2, ny2) == (x, y) ||
                        board[ny2][nx2] == '#' {
                        continue;
                    }

                    let saved_cost = distances[ny2][nx2] - distances[y][x] - 2;
                    if saved_cost > 0 {
                        cheats.entry(saved_cost).or_insert(HashSet::new()).insert(((nx, ny), (nx2, ny2)));
                    }
                }
                continue;
            }
            if distances[ny][nx] == distances[y][x] + 1 {
                q.push((nx, ny));
            }
        }
    }

    let mut count = 0;
    for (cost, tunnels) in cheats.iter() {
        if *cost >= 100 {
            count += tunnels.len();
        }
    }
    count
}

fn main() {
    let mut board = aoc::input_board();

    println!("Part I: {}", part_1(&mut board));
}
