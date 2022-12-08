import scala.io.Source

type Grid = Vector[Vector[Int]]

def gridFromInput(input: String) =
    Source.fromFile(input)
        .getLines()
        .map(_.split("").map(_.toInt).toVector)
        .toVector

def iter(grid: Grid) =
    for(
        x <- 0 until grid(0).length;
        y <- 0 until grid.length
    ) yield (x, y)

def iterUp(g: Grid, p: (Int, Int)) =
    for (y <- p._2 - 1 to 0 by -1 if y >= 0) yield g(y)(p._1)

def iterLeft(g: Grid, p: (Int, Int)) =
    for (x <- p._1 - 1 to 0 by -1 if x >= 0) yield g(p._2)(x)

def iterRight(g: Grid, p: (Int, Int)) =
    for (x <- p._1 + 1 until g(0).length if x < g(0).length) yield g(p._2)(x)

def iterDown(g: Grid, p: (Int, Int)) =
    for (y <- p._2 + 1 until g.length if y < g.length) yield g(y)(p._1)

def isVisible(g: Grid)(p: (Int, Int)) =
    List(iterUp, iterLeft, iterRight, iterDown)
        .exists(_(g, p).find(_ >= g(p._2)(p._1)).isEmpty)

def score(g: Grid, p: (Int, Int))(it: (Grid, (Int, Int)) => IndexedSeq[Int]) =
    val items = it(g, p)
    val lower = items.takeWhile(_ < g(p._2)(p._1)).size
    if (lower < items.size) lower + 1 else lower

def partI(grid: Grid) = iter(grid).filter(isVisible(grid)).size

def partII(grid: Grid) = iter(grid)
    .map(p => List(iterUp, iterLeft, iterRight, iterDown)
        .map(score(grid, p))
        .product)
    .max

@main def main(input: String) =
    val grid = gridFromInput(input)

    println(s"Part I: ${partI(grid)}")
    println(s"Part II: ${partII(grid)}")
