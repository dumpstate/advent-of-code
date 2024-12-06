use std::collections::HashSet;
mod aoc;

fn next_dir(dir_x: i32, dir_y: i32) -> (i32, i32) {
    match (dir_x, dir_y) {
        (0, -1) => (1, 0),
        (1, 0) => (0, 1),
        (0, 1) => (-1, 0),
        (-1, 0) => (0, -1),
        _ => panic!("Invalid direction")
    }
}

fn traverse(board: &Vec<Vec<char>>) -> (HashSet<(i32, i32)>, bool) {
    let (mut pos_x, mut pos_y) = aoc::find(&board, '^');
    let (mut dir_x, mut dir_y) = (0, -1);
    let mut visited = HashSet::new();
    let mut path = HashSet::new();

    while aoc::is_on_board(&board, pos_x, pos_y) {
        let px = ((pos_x, pos_y), (dir_x, dir_y));
        if visited.contains(&px) {
            return (path, true);
        }
        visited.insert(px);

        match board[pos_y as usize][pos_x as usize] {
            '.' | '^' => {
                path.insert((pos_x, pos_y));
                (pos_x, pos_y) = (pos_x + dir_x, pos_y + dir_y);
            },
            '#' => {
                (pos_x, pos_y) = (pos_x - dir_x, pos_y - dir_y);
                (dir_x, dir_y) = next_dir(dir_x, dir_y);
            },
            _ => panic!("Invalid character")
        }
    }

    (path, false)
}

fn part_1(board: &Vec<Vec<char>>) -> i32 {
    let (path, _) = traverse(board);
    path.len() as i32
}

fn part_2(board: &Vec<Vec<char>>) -> i32 {
    let mut b = board.clone();
    let mut count = 0;
    let (path, _) = traverse(&b);

    for (x, y) in path.iter() {
        if b[*y as usize][*x as usize] == '.' {
            b[*y as usize][*x as usize] = '#';
            let (_, found_cycle) = traverse(&b);
            if found_cycle {
                count += 1;
            }
            b[*y as usize][*x as usize] = '.';
        }
    }

    count
}

fn main() {
    let board = aoc::input_board();

    println!("Part I: {}", part_1(&board));
    println!("Part II: {}", part_2(&board));
}
