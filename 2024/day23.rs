use std::collections::VecDeque;
use std::collections::HashMap;
use std::collections::HashSet;
mod aoc;

fn part_1(graph: &HashMap<String, HashSet<String>>) -> usize {
    let mut paths = HashSet::new();

    for (from, _) in graph.iter() {
        let mut q = VecDeque::new();
        q.push_back((from.clone(), vec![]));

        while let Some((node, path)) = q.pop_front() {
            if path.len() == 2 {
                let mut next_path = aoc::append(&path, node.clone());
                next_path.sort();
                paths.insert(next_path.join("-"));
                continue;
            }

            let next_path = aoc::append(&path, node.clone());
            for n in graph.get(&node).unwrap() {
                if next_path.contains(n) {
                    continue;
                }
                if next_path.iter().all(|p| graph.get(p).unwrap().contains(n)) {
                    q.push_back((n.clone(), next_path.clone()));
                }
            }
        }
    }

    paths.iter()
        .filter(|p| p.starts_with("t") || p.contains("-t"))
        .count()
}

fn bron_kerbosch(
    clique: &HashSet<String>,
    potential: &mut Vec<String>,
    not_included: &mut HashSet<String>,
    graph: &HashMap<String, HashSet<String>>,
    cliques: &mut Vec<HashSet<String>>,
) {
    if potential.is_empty() && not_included.is_empty() {
        cliques.push(clique.clone());
        return;
    }

    while let Some(p) = potential.pop() {
        bron_kerbosch(
            &clique.union(&vec![p.clone()].into_iter().collect()).cloned().collect(),
            &mut potential.clone().into_iter().filter(|n| graph.get(&p).unwrap().contains(n)).collect(),
            &mut not_included.clone().into_iter().filter(|n| graph.get(&p).unwrap().contains(n)).collect(),
            graph,
            cliques,
        );
        not_included.insert(p);
    }
}

fn part_2(graph: &HashMap<String, HashSet<String>>) -> String {
    let mut cliques = vec![];
    bron_kerbosch(
        &HashSet::new(),
        &mut graph.keys().cloned().collect(),
        &mut HashSet::new(),
        graph,
        &mut cliques,
    );
    cliques.iter()
        .map(|c| {
            let mut cvec = c.clone().into_iter().collect::<Vec<String>>();
            cvec.sort();
            cvec
        })
        .max_by_key(|c| c.len())
        .unwrap()
        .join(",")
}

fn main() {
    let graph = aoc::input_lines().into_iter()
        .map(|line| {
            let s = line.split("-").collect::<Vec<&str>>();
            (s[0].to_string(), s[1].to_string())
        })
        .fold(HashMap::new(), |mut map, (a, b)| {
            map.entry(a.clone()).or_insert(HashSet::new()).insert(b.clone());
            map.entry(b.clone()).or_insert(HashSet::new()).insert(a.clone());
            map
        });
    println!("Part I: {}", part_1(&graph));
    println!("Part II: {}", part_2(&graph));
}
