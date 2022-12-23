import scala.collection.mutable.HashSet
import scala.io.Source

def mapFromInput(input: String) =
    val map = HashSet[(Int, Int)]()
    Source.fromFile(input)
        .getLines
        .zipWithIndex
        .foreach((line, iy) => line.zipWithIndex
            .foreach((char, ix) => if char == '#' then map.add((ix, iy))))
    map

def countEmpty(map: HashSet[(Int, Int)]) =
    val (xs, ys) = (map.toVector.map(_._1), map.toVector.map(_._2))
    (ys.min to ys.max).map(y => (xs.min to xs.max).count(x => !map.contains((x, y)))).sum

def allEmpty(map: HashSet[(Int, Int)], pts: (Int, Int)*) = pts.forall(!map.contains(_))
def emptyN(map: HashSet[(Int, Int)], x: Int, y: Int) = allEmpty(map, (x - 1, y - 1), (x, y - 1), (x + 1, y - 1))
def emptyS(map: HashSet[(Int, Int)], x: Int, y: Int) = allEmpty(map, (x - 1, y + 1), (x, y + 1), (x + 1, y + 1))
def emptyW(map: HashSet[(Int, Int)], x: Int, y: Int) = allEmpty(map, (x - 1, y - 1), (x - 1, y), (x - 1, y + 1))
def emptyE(map: HashSet[(Int, Int)], x: Int, y: Int) = allEmpty(map, (x + 1, y - 1), (x + 1, y), (x + 1, y + 1))
def noNeighbours(map: HashSet[(Int, Int)], x: Int, y: Int) = List(emptyN, emptyS, emptyW, emptyE).forall(_(map, x, y))
val Check = Vector(
    (emptyN, (x: Int, y: Int) => (x, y - 1)),
    (emptyS, (x: Int, y: Int) => (x, y + 1)),
    (emptyW, (x: Int, y: Int) => (x - 1, y)),
    (emptyE, (x: Int, y: Int) => (x + 1, y)))

def simulate(map: HashSet[(Int, Int)], rounds: Option[Int] = None) =
    var (isStable, round) = (false, 0)

    while rounds.map(round < _).getOrElse(!isStable) do
        val moves = map.toVector
            .flatMap {
                case (x, y) if noNeighbours(map, x, y) => None
                case (x, y) => (0 to 3).iterator
                    .map(it => Check((it + round) % Check.length))
                    .filter((pred, _) => pred(map, x, y))
                    .map((_, prod) => ((x, y), prod(x, y)))
                    .nextOption
            }
        val counts = moves.map(_._2).groupBy(identity).mapValues(_.size)
        moves foreach { (from, to) =>
            if counts(to) < 2 then
                map.remove(from)
                map.add(to)
        }
        if moves.isEmpty then isStable = true
        round += 1

    (map, round)

def partI(input: String) = countEmpty(simulate(mapFromInput(input), Some(10))._1)
def partII(input: String) = simulate(mapFromInput(input))._2

@main def main(input: String) =
    println(s"Part I: ${partI(input)}")
    println(s"Part II: ${partII(input)}")
