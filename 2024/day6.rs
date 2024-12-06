use std::collections::HashSet;
mod aoc;

fn find(board: &Vec<Vec<char>>, c: char) -> (i32, i32) {
    for y in 0..board.len() {
        for x in 0..board[0].len() {
            if board[y][x] == c {
                return (x as i32, y as i32);
            }
        }
    }

    panic!("No start found");
}

fn is_on_board(board: &Vec<Vec<char>>, x: i32, y: i32) -> bool {
    x >= 0 && y >= 0 && y < board.len() as i32 && x < board[0].len() as i32
}

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
    let (mut pos_x, mut pos_y) = find(&board, '^');
    let (mut dir_x, mut dir_y) = (0, -1);
    let mut visited = HashSet::new();
    let mut path = HashSet::new();

    while is_on_board(&board, pos_x, pos_y) {
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
