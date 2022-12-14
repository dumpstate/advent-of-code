import scala.collection.mutable
import scala.io.Source

def pathsFromInput(input: String) = Source.fromFile(input).getLines
    .map(_.split(" -> ").map(_.split(",").map(_.toInt))
        .map(p => (p(0), p(1))).toList).toList

def iter(from: (Int, Int), to: (Int, Int)) = (from, to) match
    case ((a, b), (c, d)) if a == c && d >= b => for (y <- b to d) yield (a, y)
    case ((a, b), (c, d)) if a == c && d < b => for (y <- d to b) yield (a, y)
    case ((a, b), (c, d)) if b == d && c >= a => for (x <- a to c) yield (x, b)
    case ((a, b), (c, d)) if b == d && c < a => for (x <- c to a) yield (x, b)

class Cave(val map: mutable.Map[(Int, Int), Char]):
    val ymax = map.keySet.map(_._2).max

    def isEmpty(pt: (Int, Int)) = map(pt) == '.'

    def sandCount() = map.count((_, v) => v == 'o')

    def drop(source: (Int, Int) = (500, 0)): Option[Char] = source match
        case (x, y) if !isEmpty((x, y)) || y > ymax => None
        case (x, y) if isEmpty((x, y + 1)) => drop((x, y + 1))
        case (x, y) if isEmpty((x - 1, y + 1)) => drop((x - 1, y + 1))
        case (x, y) if isEmpty((x + 1, y + 1)) => drop((x + 1, y + 1))
        case (x, y) if isEmpty((x, y)) => map.put((x, y), 'o')

    def fill() =
        var count = sandCount()
        drop()
        while count != sandCount() do
            count = sandCount()
            drop()
        count

object Cave:
    def apply(paths: List[List[(Int, Int)]], floorOffset: Option[Int] = None, inf: Int = 1000): Cave =
        val map = mutable.HashMap[(Int, Int), Char]()

        for
            path <- paths
            pair <- path.sliding(2)
            point <- iter(pair(0), pair(1))
        do
            map.put(point, '#')

        floorOffset.foreach(offset =>
            val xs = map.keySet.map(_._1)
            val y = map.keySet.map(_._2).max + 2
            for x <- (xs.min - inf) to (xs.max + inf) do
                map.put((x, y), '#'))

        new Cave(map.withDefaultValue('.'))

def partI(input: String) = Cave(pathsFromInput(input)).fill()
def partII(input: String) = Cave(pathsFromInput(input), Some(2)).fill()

@main def main(input: String) =
    println(s"Part I: ${partI(input)}")
    println(s"Part II: ${partII(input)}")
