use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn count_fence(board: &Vec<Vec<char>>, pt: (usize, usize)) -> (i64, HashSet<char>) {
    let (x, y) = pt;
    let c = board[y][x];
    let (mut count, mut sides) = (0, HashSet::new());

    for (dx, dy, d) in &[(-1, 0, '<'), (1, 0, '>'), (0, -1, '^'), (0, 1, 'v')] {
        let (nx, ny) = (x as i32 + dx, y as i32 + dy);
        if !aoc::is_on_board(board, nx, ny) {
            count += 1;
            sides.insert(*d);
        } else {
            let nc = board[ny as usize][nx as usize];
            if nc != c {
                count += 1;
                sides.insert(*d);
            }
        }
    }

    (count, sides)
}

fn count_sides(sides: &HashMap<char, HashSet<(usize, usize)>>) -> i64 {
    let mut count = 0;
    for (d, pts) in sides {
        let dirs = match d {
            '<' | '>' => &[(0, -1), (0, 1)],
            '^' | 'v' => &[(1, 0), (-1, 0)],
            _ => panic!("Invalid direction"),
        };
        let mut q = pts.clone().into_iter().map(|(x, y)| (x as i32, y as i32)).collect::<HashSet<_>>();
        loop {
            let (x, y) = match q.iter().next() {
                Some(pt) => *pt,
                None => break,
            };
            count += 1;
            q.remove(&(x, y));
            for (dx, dy) in dirs {
                let (mut nx, mut ny) = (x as i32 + dx, y as i32 + dy);
                while q.contains(&(nx, ny)) {
                    q.remove(&(nx, ny));
                    (nx, ny) = (nx as i32 + dx, ny as i32 + dy);
                }
            }
        }
    }
    count
}

fn flood(board: &Vec<Vec<char>>, start: (usize, usize), visited: &mut HashSet<(usize, usize)>) -> (char, i64, i64, i64) {
    let (mut area, mut fence) = (0, 0);
    let (c, mut q, mut all_sides) = (board[start.1][start.0], Vec::new(), HashMap::new());
    q.push(start);

    while !q.is_empty() {
        let (x, y) = q.pop().unwrap();
        if board[y][x] != c || visited.contains(&(x, y)) {
            continue;
        }

        visited.insert((x, y));
        area += 1;
        let (f, sides) = count_fence(board, (x, y));
        fence += f;
        for d in sides {
            let pts = all_sides.entry(d).or_insert(HashSet::new());
            pts.insert((x, y));
        }
        for (dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = (x as i32 + dx, y as i32 + dy);
            if aoc::is_on_board(board, nx, ny) {
                q.push((nx as usize, ny as usize));
            }
        }
    }

    (c, area, fence, count_sides(&all_sides))
}

fn main() {
    let board = aoc::input_board();
    let (mut visited, mut regions) = (HashSet::new(), Vec::new());

    for y in 0..board.len() {
        for x in 0..board[0].len() {
            if visited.contains(&(x, y)) {
                continue;
            }
            regions.push(flood(&board, (x, y), &mut visited));
        }
    }

    println!("Part I: {}", regions.iter().map(|(_, a, f, _)| { *a * *f }).sum::<i64>());
    println!("Part II: {}", regions.iter().map(|(_, a, _, s)| { *a * *s }).sum::<i64>());
}
