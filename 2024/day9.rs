mod aoc;

fn expand(disk_map: &Vec<usize>) -> Vec<Option<usize>> {
    let mut expanded = Vec::new();
    let mut next_space = false;
    let mut next_id = 0;

    for c in disk_map.iter() {
        if next_space {
            for _ in 0..*c {
                expanded.push(None);
            }
            next_id += 1;
        } else {
            for _ in 0..*c {
                expanded.push(Some(next_id));
            }
        }
        next_space = !next_space;
    }

    expanded
}

fn compress(disk_map: &mut Vec<Option<usize>>) {
    let (mut l_ix, mut r_ix) = (0, disk_map.len() - 1);

    while l_ix < r_ix {
        while disk_map[l_ix].is_some() {
            l_ix += 1;
        }

        while disk_map[r_ix].is_none() {
            r_ix -= 1;
        }

        while l_ix < r_ix && disk_map[l_ix].is_none() && disk_map[r_ix].is_some() {
            disk_map.swap(l_ix, r_ix);
            l_ix += 1;
            r_ix -= 1;
        }
    }
}

fn lsize(disk_map: &Vec<Option<usize>>, ix: usize) -> usize {
    let (mut size, mut i) = (0, ix);

    while i < disk_map.len() && disk_map[i] == disk_map[ix] {
        size += 1;
        i += 1;
    }

    size
}

fn next_space(disk_map: &Vec<Option<usize>>, ix: usize) -> (usize, usize) {
    let mut i = ix;

    while i < disk_map.len() && disk_map[i].is_none() {
        i += 1;
    }

    while i < disk_map.len() && disk_map[i].is_some() {
        i += 1;
    }

    (i, lsize(disk_map, i))
}

fn find_id(disk_map: &Vec<Option<usize>>, id: usize) -> usize {
    disk_map.iter().enumerate().find(|(_, x)| {
        match x {
            Some(y) => *y == id,
            None => false
        }
    }).unwrap().0
}

fn compress_whole(disk_map: &mut Vec<Option<usize>>) {
    let mut id = disk_map.iter().max().unwrap().unwrap();

    loop {
        let ix = find_id(disk_map, id);
        let size = lsize(disk_map, ix);

        let (mut l_ix, mut found) = (0, false);
        loop {
            let (next_l_ix, next_size) = next_space(disk_map, l_ix);
            if next_size >= size {
                found = true;
                l_ix = next_l_ix;
                break;
            }
            if next_l_ix == l_ix {
                break;
            }
            l_ix = next_l_ix;
        }

        if found && l_ix < ix {
            for i in 0..size {
                disk_map.swap(ix + i, l_ix + i);
            }
        }

        if id == 0 {
            break;
        }
        id -= 1;
    }
}

fn checksum(disk_map: &Vec<Option<usize>>) -> i64 {
    disk_map.iter().enumerate().map(|(ix, id)| {
        match id {
            Some(x) => *x as i64 * ix as i64,
            None => 0
        }
    }).sum()
}

fn part_1(input: &Vec<usize>) -> i64 {
    let mut expanded = expand(input);
    compress(&mut expanded);
    checksum(&expanded)
}

fn part_2(input: &Vec<usize>) -> i64 {
    let mut expanded = expand(input);
    compress_whole(&mut expanded);
    checksum(&expanded)
}

fn main() {
    let input = &aoc::input_lines()[0].chars().map(|x| x.to_digit(10).unwrap() as usize).collect::<Vec<usize>>();

    println!("Part I: {}", part_1(&input));
    println!("Part II: {}", part_2(&input));
}
