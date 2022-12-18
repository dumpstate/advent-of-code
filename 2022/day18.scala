import scala.collection.mutable
import scala.io.Source

def iterNeighbours(cube: (Int, Int, Int), diag: Boolean = false) = (cube, diag) match
    case ((x, y, z), false) => IndexedSeq(
        (x + 1, y, z),
        (x - 1, y, z),
        (x, y + 1, z),
        (x, y - 1, z),
        (x, y, z + 1),
        (x, y, z - 1))
    case ((x, y, z), true) => for (
        a <- (x - 1) to (x + 1);
        b <- (y - 1) to (y + 1);
        c <- (z - 1) to (z + 1);
        if (a, b, c) != (x, y, z)
    ) yield (a, b, c)

def surface(cubes: Set[(Int, Int, Int)], airPockets: Set[(Int, Int, Int)] = Set()) =
    cubes.iterator.map(iterNeighbours(_)
        .filter(!cubes.contains(_))
        .filter(!airPockets.contains(_)).length).sum

def partI(cubes: Set[(Int, Int, Int)]) = surface(cubes)

def partII(cubes: Set[(Int, Int, Int)]) =
    val surfaceCubes = cubes.iterator.flatMap(iterNeighbours(_, true))
        .filter(!cubes.contains(_)).toSet
    val stack = mutable.Stack.from(surfaceCubes)
    val visited = mutable.HashSet[(Int, Int, Int)]()
    val groups = mutable.ArrayBuffer[mutable.HashSet[(Int, Int, Int)]]()
    while stack.nonEmpty do
        val cube = stack.pop()
        if !visited.contains(cube) then
            visited.add(cube)
            groups.find(gr => iterNeighbours(cube).find(n => gr.contains(n)).isDefined) match
                case Some(group) => group.add(cube)
                case None => groups.addOne(mutable.HashSet(cube))
            iterNeighbours(cube).filter(n => surfaceCubes.contains(n))
                .foreach(n => stack.push(n))

    surface(cubes, groups.toList.sortWith(_.size > _.size).tail.reduce(_ union _).toSet)

@main def main(input: String) =
    val cubes = Source.fromFile(input)
        .getLines.map(line =>
            val split = line.split(",").map(_.toInt)
            (split(0), split(1), split(2))).toSet

    println(s"Part I: ${partI(cubes)}")
    println(s"Part II: ${partII(cubes)}")
