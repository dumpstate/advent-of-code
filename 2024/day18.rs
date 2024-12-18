use std::collections::BinaryHeap;
mod aoc;

fn shortest_path(pts: &Vec<(usize, usize)>, byte_count: usize) -> Option<i64> {
    let mut board = vec![vec!['.'; 71]; 71];
    for (x, y) in pts[..byte_count].iter() {
        board[*y][*x] = '#';
    }

    let mut dist = vec![vec![std::i64::MAX; 71]; 71];
    let mut heap = BinaryHeap::new();

    dist[0][0] = 0;
    heap.push(aoc::State {
        pos: (0, 0),
        cost: 0,
    });

    while let Some(aoc::State { cost, pos }) = heap.pop() {
        if pos == (70, 70) {
            return Some(cost);
        }

        if dist[pos.1][pos.0] < cost {
            continue;
        }

        for (dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = ((pos.0 as i32 + dx) as usize, (pos.1 as i32 + dy) as usize);
            if !aoc::is_on_board(&board, nx as i32, ny as i32) || board[ny][nx] == '#' {
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

    None
}

fn part_2(pts: &Vec<(usize, usize)>) -> String {
    let (mut l, mut h) = (1024, pts.len());
    while l < h {
        let mid = (l + h) / 2;
        match shortest_path(pts, mid) {
            Some(_) => l = mid + 1,
            None => h = mid,
        }
    }
    format!("{},{}", pts[l - 1].0, pts[l - 1].1)
}

fn main() {
    let pts = aoc::input_lines()
        .into_iter()
        .map(|line| {
            let ns = aoc::split(&line, ",")
                .iter()
                .map(|n| n.parse::<usize>().unwrap())
                .collect::<Vec<_>>();
            (ns[0], ns[1])
        })
        .collect::<Vec<_>>();

    println!("Part I: {}", shortest_path(&pts, 1024).unwrap());
    println!("Part II: {}", part_2(&pts));
}
