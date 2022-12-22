import scala.collection.mutable.ArrayDeque
import scala.io.Source

def mix(ns: Vector[Long], loops: Int = 1) =
    val data = ArrayDeque[(Long, Int)]().addAll(ns.zipWithIndex)

    for
        _ <- 1 to loops
        n <- ns.zipWithIndex
    do
        val ix = data.indexOf(n)
        val (off, _) = data.remove(ix)
        data.insert(Math.floorMod(ix + off, data.size.toLong).toInt, n)

    data.toVector

def coords(ns: Vector[(Long, Int)]) =
    List(1000, 2000, 3000)
        .map(n => (n + ns.indexOf(ns.find(_._1 == 0).get)) % ns.length)
        .map(ns(_)._1).sum

def partI(ns: Vector[Long]) = coords(mix(ns))
def partII(ns: Vector[Long]) = coords(mix(ns.map(_ * 811589153), 10))

@main def main(input: String) =
    val ns = Source.fromFile(input).getLines.map(_.toLong).toVector

    println(s"Part I: ${partI(ns)}")
    println(s"Part II: ${partII(ns)}")
