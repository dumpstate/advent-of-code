mod aoc;

fn count_word(board: &Vec<Vec<char>>, word: &str, x: usize, y: usize) -> i32 {
    let mut count = 0;

    let x_size = board[0].len();
    let y_size = board.len();
    let directions = [
        (0, 1),
        (1, 0),
        (1, 1),
        (1, -1),
        (-1, 1),
        (-1, 0),
        (0, -1),
        (-1, -1),
    ];

    for (dx, dy) in directions.iter() {
        let mut found = true;
        for i in 0..word.len() {
            let nx = x as i32 + dx * i as i32;
            let ny = y as i32 + dy * i as i32;

            if nx < 0 || nx >= x_size as i32 || ny < 0 || ny >= y_size as i32 {
                found = false;
                break;
            }

            if board[ny as usize][nx as usize] != word.chars().nth(i).unwrap() {
                found = false;
                break;
            }
        }

        if found {
            count += 1;
        }
    }

    count
}

fn has_xmas(board: &Vec<Vec<char>>, x: usize, y: usize) -> bool {
    let word = ['M', 'A', 'S'];

    for d in 0..3 {
        if board[y + d][x + d] != word[d] &&
            board[y + 2 - d][x + 2 - d] != word[d] {
            return false;
        }

        if board[y + d][x + 2 - d] != word[d] &&
            board[y + 2 - d][x + d] != word[d] {
            return false;
        }
    }

    true
}

fn part_1(board: &Vec<Vec<char>>) -> i32 {
    let word = "XMAS";
    let mut count = 0;

    for y in 0..board.len() {
        for x in 0..board[0].len() {
            count += count_word(board, &word, x, y)
        }
    }

    count
}

fn part_2(board: &Vec<Vec<char>>) -> i32 {
    let mut count = 0;

    for y in 0..(board.len() - 2) {
        for x in 0..(board[0].len() - 2) {
            if has_xmas(board, x, y) {
                count += 1
            }
        }
    }

    count
}

fn main() {
    let board = aoc::input_board();

    println!("Part I: {}", part_1(&board));
    println!("Part II: {}", part_2(&board));
}
