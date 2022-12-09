import scala.io.Source

type Knot = ((Int, Int), List[(Int, Int)])

def commandsFromInput(input: String) =
    Source.fromFile(input)
        .getLines()
        .map(_.split(" ") match
            case Array(dir, count) => (dir, count.toInt))
        .toList

def expand(cmds: List[(String, Int)]) =
    cmds.foldLeft(List[String]())((acc, next) => next match
        case (dir, count) => List.fill(count)(dir).concat(acc)
    ).reverse

def diag(p1: (Int, Int), p2: (Int, Int)) = p1._1 != p2._1 && p1._2 != p2._2

def wait(p1: (Int, Int), p2: (Int, Int)) =
    Math.abs(p1._1 - p2._1) + Math.abs(p1._2 - p2._2) <= (if diag(p1, p2) then 2 else 1)

def next(p: (Int, Int), dir: String) = (p, dir) match
    case ((x, y), "R") => (x + 1, y)
    case ((x, y), "L") => (x - 1, y)
    case ((x, y), "U") => (x, y - 1)
    case ((x, y), "D") => (x, y + 1)

def follow(h: (Int, Int), t: (Int, Int)) =
    if wait(h, t) then t
    else (h, t) match
        case ((hx, hy), (tx, ty)) if hy == ty && hx > tx => (tx + 1, ty)
        case ((hx, hy), (tx, ty)) if hy == ty && hx < tx => (tx - 1, ty)
        case ((hx, hy), (tx, ty)) if hx == tx && hy > ty => (tx, ty + 1)
        case ((hx, hy), (tx, ty)) if hx == tx && hy < ty => (tx, ty - 1)
        case ((hx, hy), (tx, ty)) if hx > tx && hy > ty => (tx + 1, ty + 1)
        case ((hx, hy), (tx, ty)) if hx > tx && hy < ty => (tx + 1, ty - 1)
        case ((hx, hy), (tx, ty)) if hx < tx && hy > ty => (tx - 1, ty + 1)
        case ((hx, hy), (tx, ty)) if hx < tx && hy < ty => (tx - 1, ty - 1)
        case _ => throw RuntimeException("oops")

def newRope(length: Int): List[Knot] = List.fill(length)(((0, 0), List[(Int, Int)]()))

def propagate(lead: (Int, Int), rope: List[Knot]): List[Knot] = rope match
    case (h, path) :: tail => (follow(lead, h), h :: path) :: propagate(follow(lead, h), tail)
    case Nil => Nil

def apply(rope: List[Knot], dir: String): List[Knot] = rope match
    case (h, path) :: tail => (next(h, dir), h :: path) :: propagate(next(h, dir), tail)
    case Nil => Nil

def lastKnotCoverage(rope: List[Knot]) = (rope.last._2.toSet + rope.last._1).size

def simulate(cmds: List[(String, Int)], length: Int) =
    expand(cmds).foldLeft(newRope(length))(apply(_, _))

def partI(cmds: List[(String, Int)]) = lastKnotCoverage(simulate(cmds, 2))
def partII(cmds: List[(String, Int)]) = lastKnotCoverage(simulate(cmds, 10))

@main def main(input: String) =
    val cmds = commandsFromInput(input)

    println(s"Part I: ${partI(cmds)}")
    println(s"Part II: ${partII(cmds)}")
