mod aoc;

fn expand(disk_map: &Vec<usize>) -> Vec<Option<usize>> {
    disk_map
        .iter()
        .fold(
            (Vec::<Option<usize>>::new(), false, 0),
            |(res, next_space, id), c| match next_space {
                true => ([res, vec![None; *c]].concat(), false, id),
                false => ([res, vec![Some(id); *c]].concat(), true, id + 1),
            },
        )
        .0
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

fn next_space(disk_map: &Vec<Option<usize>>, ix: usize) -> Option<(usize, usize)> {
    let mut i = ix;
    while i < disk_map.len() && disk_map[i].is_none() {
        i += 1;
    }
    while i < disk_map.len() && disk_map[i].is_some() {
        i += 1;
    }
    match i {
        j if j == ix => None,
        _ => Some((i, lsize(disk_map, i))),
    }
}

fn find_id(disk_map: &Vec<Option<usize>>, id: usize) -> usize {
    disk_map
        .iter()
        .enumerate()
        .find(|(_, x)| match x {
            Some(y) => *y == id,
            None => false,
        })
        .unwrap()
        .0
}

fn compress_whole(disk_map: &mut Vec<Option<usize>>) {
    for id in (0..=disk_map.iter().max().unwrap().unwrap()).rev() {
        let ix = find_id(disk_map, id);
        let (size, mut l) = (lsize(disk_map, ix), None);
        loop {
            match next_space(disk_map, l.unwrap_or(0)) {
                Some((next_l, next_size)) if next_size >= size => {
                    l = Some(next_l);
                    break;
                }
                Some((next_l, _)) => l = Some(next_l),
                None => break,
            }
        }

        if l.is_some() && l.unwrap() < ix {
            for i in 0..size {
                disk_map.swap(ix + i, l.unwrap() + i);
            }
        }
    }
}

fn checksum(disk_map: &Vec<Option<usize>>) -> i64 {
    disk_map
        .iter()
        .enumerate()
        .map(|(ix, id)| match id {
            Some(x) => *x as i64 * ix as i64,
            None => 0,
        })
        .sum()
}

fn compress_and_checksum(disk_map: &Vec<usize>, cmpr: fn(&mut Vec<Option<usize>>)) -> i64 {
    let mut expanded = expand(disk_map);
    cmpr(&mut expanded);
    checksum(&expanded)
}

fn main() {
    let input = &aoc::input_lines()[0]
        .chars()
        .map(|x| x.to_digit(10).unwrap() as usize)
        .collect::<Vec<usize>>();

    println!("Part I: {}", compress_and_checksum(&input, compress));
    println!("Part II: {}", compress_and_checksum(&input, compress_whole));
}
