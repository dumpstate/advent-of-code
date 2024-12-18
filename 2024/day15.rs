mod aoc;

fn parse(lines: &Vec<String>) -> (Vec<Vec<char>>, Vec<char>) {
    let (mut board, mut moves) = (Vec::new(), Vec::new());
    let mut parse_board = true;

    for line in lines {
        if line.is_empty() {
            parse_board = false;
            continue;
        }

        if parse_board {
            board.push(line.chars().collect());
        } else {
            moves.extend(line.chars());
        }
    }

    (board, moves)
}

fn score(board: &Vec<Vec<char>>) -> i64 {
    let mut total = 0;
    for y in 0..board.len() {
        for x in 0..board[y].len() {
            if board[y][x] == 'O' || board[y][x] == '[' {
                total += x as i64 + (y as i64) * 100;
            }
        }
    }
    total
}

fn find_robot(board: &Vec<Vec<char>>) -> (usize, usize) {
    for y in 0..board.len() {
        for x in 0..board[y].len() {
            if board[y][x] == '@' {
                return (x, y);
            }
        }
    }
    panic!("Robot not found!");
}

fn try_push(board: &mut Vec<Vec<char>>, bx: (usize, usize), dir: (i32, i32)) -> bool {
    let (mut nx, mut ny) = bx;
    while board[ny][nx] == 'O' {
        nx = (nx as i32 + dir.0) as usize;
        ny = (ny as i32 + dir.1) as usize;
    }

    if board[ny][nx] == '.' {
        board[ny][nx] = 'O';
        board[bx.1][bx.0] = '.';
        return true;
    }

    (nx, ny) = bx;
    if board[ny][nx] == '[' || board[ny][nx] == ']' {
        let mut to_move = vec![];
        let mut do_move = vec![];

        if board[ny][nx] == '[' {
            to_move.push((nx, ny));
        }
        if board[ny][nx] == ']' {
            to_move.push((nx - 1, ny));
        }

        while let Some((x, y)) = to_move.pop() {
            let surface = if dir.0 == 0 {
                vec![(x, y), (x + 1, y)]
            } else {
                if dir.0 == -1 { vec![(x, y)] }
                else { vec![(x + 1, y)] }
            };

            for (sx, sy) in surface {
                let (dx, dy) = ((sx as i32 + dir.0) as usize, (sy as i32 + dir.1) as usize);
                match board[dy][dx] {
                    '#' => return false,
                    '[' => {
                        if !to_move.contains(&(dx, dy)) {
                            to_move.push((dx, dy));
                        }
                    },
                    ']' => {
                        if !to_move.contains(&(dx - 1, dy)) {
                            to_move.push((dx - 1, dy));
                        }
                    },
                    _ => {},
                };
            }

            do_move.push((x, y));
        }

        for (x, y) in do_move.iter().rev() {
            let (dx, dy) = ((*x as i32 + dir.0) as usize, (*y as i32 + dir.1) as usize);
            board[dy][dx] = '[';
            board[dy][dx + 1] = ']';
            match dir {
                (-1, 0) => { board[*y as usize][*x as usize + 1] = '.'; },
                (1, 0) => { board[*y as usize][*x as usize] = '.'; },
                (0, _) => {
                    board[*y as usize][*x as usize] = '.';
                    board[*y as usize][*x as usize + 1] = '.';
                },
                _ => panic!("invalid direction"),
            }
        }

        return true;
    }

    false
}

fn move_robot(board: &mut Vec<Vec<char>>, m: char) {
    let (rx, ry) = find_robot(board);
    let (dx, dy): (i32, i32) = match m {
        '>' => (1, 0),
        '<' => (-1, 0),
        '^' => (0, -1),
        'v' => (0, 1),
        _ => panic!("Invalid move"),
    };
    let (nx, ny) = ((rx as i32 + dx) as usize, (ry as i32 + dy) as usize);
    match board[ny][nx] {
        '#' => (),
        '.' => {
            board[ry][rx] = '.';
            board[ny][nx] = '@';
        }
        'O' | '[' | ']' => {
            if try_push(board, (nx, ny), (dx, dy)) {
                board[ry][rx] = '.';
                board[ny][nx] = '@';
            }
        }
        _ => panic!("Invalid character"),
    }
}

fn inflate(board: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut inflated = Vec::new();
    for line in board {
        let mut new_line = Vec::new();
        for c in line {
            match *c {
                '#' => {
                    new_line.push('#');
                    new_line.push('#');
                }
                'O' => {
                    new_line.push('[');
                    new_line.push(']');
                }
                '.' => {
                    new_line.push('.');
                    new_line.push('.');
                }
                '@' => {
                    new_line.push('@');
                    new_line.push('.');
                }
                _ => panic!("Invalid character"),
            }
        }
        inflated.push(new_line);
    }
    inflated
}

fn apply_moves(board: &Vec<Vec<char>>, moves: &Vec<char>) -> i64 {
    let mut board = board.clone();
    for m in moves {
        move_robot(&mut board, *m);
    }
    score(&board)
}

fn main() {
    let (board, moves) = parse(&aoc::input_lines());
    println!("Part I: {}", apply_moves(&board, &moves));
    println!("Part II: {}", apply_moves(&inflate(&board), &moves));
}
